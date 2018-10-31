// Copyright (C) 2016-Present Pivotal Software, Inc. All rights reserved.
// This program and the accompanying materials are made available under the terms of the under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

package task

import (
	"fmt"
	"log"
	"reflect"

	"errors"

	"github.com/pivotal-cf/on-demand-service-broker/boshdirector"
	"github.com/pivotal-cf/on-demand-service-broker/broker"
	"github.com/pivotal-cf/on-demand-services-sdk/bosh"
	"github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
	"gopkg.in/yaml.v2"
)

//go:generate counterfeiter -o fakes/fake_bosh_client.go . BoshClient
type BoshClient interface {
	Deploy(manifest []byte, contextID string, logger *log.Logger, reporter *boshdirector.AsyncTaskReporter) (int, error)
	GetTasks(deploymentName string, logger *log.Logger) (boshdirector.BoshTasks, error)
	GetDeployment(name string, logger *log.Logger) ([]byte, bool, error)
}

//go:generate counterfeiter -o fakes/fake_manifest_generator.go . ManifestGenerator
type ManifestGenerator interface {
	GenerateManifest(
		deploymentName,
		planID string,
		requestParams map[string]interface{},
		oldManifest []byte,
		previousPlanID *string,
		secretsMap map[string]string,
		logger *log.Logger,
	) (serviceadapter.MarshalledGenerateManifest, error)
}

//go:generate counterfeiter -o fakes/fake_odb_secrets.go . ODBSecrets

type ODBSecrets interface {
	GenerateSecretPaths(deploymentName, manifest string, secretsMap serviceadapter.ODBManagedSecrets) []broker.ManifestSecret
	ReplaceODBRefs(manifest string, secrets []broker.ManifestSecret) string
}

//go:generate counterfeiter -o fakes/fake_bulk_setter.go . BulkSetter

type BulkSetter interface {
	BulkSet([]broker.ManifestSecret) error
}

type deployer struct {
	boshClient        BoshClient
	manifestGenerator ManifestGenerator
	odbSecrets        ODBSecrets
	bulkSetter        BulkSetter
}

func NewDeployer(boshClient BoshClient, manifestGenerator ManifestGenerator, odbSecrets ODBSecrets, bulkSetter BulkSetter) deployer {
	return deployer{
		boshClient:        boshClient,
		manifestGenerator: manifestGenerator,
		odbSecrets:        odbSecrets,
		bulkSetter:        bulkSetter,
	}
}

func (d deployer) Create(deploymentName, planID string, requestParams map[string]interface{}, boshContextID string, logger *log.Logger) (int, []byte, error) {
	err := d.assertNoOperationsInProgress(deploymentName, logger)
	if err != nil {
		return 0, nil, err
	}

	return d.doDeploy(deploymentName, planID, "create", requestParams, nil, nil, boshContextID, nil, logger)
}

func (d deployer) Upgrade(deploymentName, planID string, previousPlanID *string, boshContextID string, logger *log.Logger) (int, []byte, error) {
	err := d.assertNoOperationsInProgress(deploymentName, logger)
	if err != nil {
		return 0, nil, err
	}

	oldManifest, err := d.getDeploymentManifest(deploymentName, logger)
	if err != nil {
		return 0, nil, err
	}

	return d.doDeploy(deploymentName, planID, "upgrade", nil, oldManifest, previousPlanID, boshContextID, nil, logger)
}

func (d deployer) Update(
	deploymentName,
	planID string,
	requestParams map[string]interface{},
	previousPlanID *string,
	boshContextID string,
	oldSecretsMap map[string]string,
	logger *log.Logger,
) (int, []byte, error) {
	if err := d.assertNoOperationsInProgress(deploymentName, logger); err != nil {
		return 0, nil, err
	}

	oldManifest, err := d.getDeploymentManifest(deploymentName, logger)
	if err != nil {
		return 0, nil, err
	}

	if err := d.checkForPendingChanges(deploymentName, previousPlanID, oldManifest, oldSecretsMap, logger); err != nil {
		return 0, nil, err
	}

	return d.doDeploy(deploymentName, planID, "update", requestParams, oldManifest, previousPlanID, boshContextID, oldSecretsMap, logger)
}

func (d deployer) getDeploymentManifest(deploymentName string, logger *log.Logger) ([]byte, error) {
	oldManifest, found, err := d.boshClient.GetDeployment(deploymentName, logger)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, broker.NewDeploymentNotFoundError(fmt.Errorf("bosh deployment '%s' not found", deploymentName))
	}

	return oldManifest, nil
}

func (d deployer) assertNoOperationsInProgress(deploymentName string, logger *log.Logger) error {
	clientTasks, err := d.boshClient.GetTasks(deploymentName, logger)
	if err != nil {
		return broker.NewServiceError(fmt.Errorf("error getting tasks for deployment %s: %s\n", deploymentName, err))
	}

	if incompleteTasks := clientTasks.IncompleteTasks(); len(incompleteTasks) != 0 {
		logger.Printf("deployment %s is still in progress: tasks %s\n", deploymentName, incompleteTasks.ToLog())
		return broker.TaskInProgressError{Message: "task in progress"}
	}

	return nil
}

func (d deployer) checkForPendingChanges(
	deploymentName string,
	previousPlanID *string,
	rawOldManifest RawBoshManifest,
	oldSecretsMap map[string]string,
	logger *log.Logger,
) error {
	regeneratedManifestContent, err := d.manifestGenerator.GenerateManifest(deploymentName, *previousPlanID, map[string]interface{}{}, rawOldManifest, previousPlanID, oldSecretsMap, logger)
	if err != nil {
		return err
	}

	regeneratedManifest, err := marshalBoshManifest([]byte(regeneratedManifestContent.Manifest))
	if err != nil {
		return err
	}
	ignoreUpdateBlock(&regeneratedManifest)

	oldManifest, err := marshalBoshManifest(rawOldManifest)
	if err != nil {
		return err
	}
	ignoreUpdateBlock(&oldManifest)

	manifestsSame := reflect.DeepEqual(regeneratedManifest, oldManifest)

	pendingChanges := !manifestsSame

	if pendingChanges {
		return broker.NewPendingChangesNotAppliedError(errors.New("There are pending changes"))
	}

	return nil
}

func (d deployer) doDeploy(
	deploymentName,
	planID string,
	operationType string,
	requestParams map[string]interface{},
	oldManifest []byte,
	previousPlanID *string,
	boshContextID string,
	oldSecretsMap map[string]string,
	logger *log.Logger,
) (int, []byte, error) {

	generateManifestOutput, err := d.manifestGenerator.GenerateManifest(deploymentName, planID, requestParams, oldManifest, previousPlanID, oldSecretsMap, logger)
	if err != nil {
		return 0, nil, err
	}
	manifest := generateManifestOutput.Manifest

	if d.bulkSetter != nil && !reflect.ValueOf(d.bulkSetter).IsNil() {
		secrets := d.odbSecrets.GenerateSecretPaths(deploymentName, manifest, generateManifestOutput.ODBManagedSecrets)
		if err = d.bulkSetter.BulkSet(secrets); err != nil {
			return 0, nil, err
		}
		manifest = d.odbSecrets.ReplaceODBRefs(generateManifestOutput.Manifest, secrets)
	}

	boshTaskID, err := d.boshClient.Deploy([]byte(manifest), boshContextID, logger, boshdirector.NewAsyncTaskReporter())
	if err != nil {
		return 0, nil, fmt.Errorf("error deploying instance: %s\n", err)
	}
	logger.Printf("Bosh task ID for %s deployment %s is %d\n", operationType, deploymentName, boshTaskID)

	return boshTaskID, []byte(manifest), nil
}

func marshalBoshManifest(rawManifest []byte) (bosh.BoshManifest, error) {
	var boshManifest bosh.BoshManifest
	err := yaml.Unmarshal(rawManifest, &boshManifest)

	if err != nil {
		return bosh.BoshManifest{}, fmt.Errorf("error detecting change in manifest, unable to unmarshal manifest: %s", err)
	}
	return boshManifest, nil
}

func ignoreUpdateBlock(manifest *bosh.BoshManifest) {
	manifest.Update = nil
	for i := range manifest.InstanceGroups {
		manifest.InstanceGroups[i].Update = nil
	}
}
