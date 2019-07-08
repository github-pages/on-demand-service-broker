// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	serviceadapter "github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
)

func AnyServiceadapterPlan() serviceadapter.Plan {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(serviceadapter.Plan))(nil)).Elem()))
	var nullValue serviceadapter.Plan
	return nullValue
}

func EqServiceadapterPlan(value serviceadapter.Plan) serviceadapter.Plan {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue serviceadapter.Plan
	return nullValue
}