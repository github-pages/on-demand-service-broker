// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	bosh "github.com/pivotal-cf/on-demand-services-sdk/bosh"
)

func AnyBoshBoshVMs() bosh.BoshVMs {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(bosh.BoshVMs))(nil)).Elem()))
	var nullValue bosh.BoshVMs
	return nullValue
}

func EqBoshBoshVMs(value bosh.BoshVMs) bosh.BoshVMs {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue bosh.BoshVMs
	return nullValue
}