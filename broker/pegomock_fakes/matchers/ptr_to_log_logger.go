// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	log "log"
)

func AnyPtrToLogLogger() *log.Logger {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*log.Logger))(nil)).Elem()))
	var nullValue *log.Logger
	return nullValue
}

func EqPtrToLogLogger(value *log.Logger) *log.Logger {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *log.Logger
	return nullValue
}