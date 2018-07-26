package container

import (
	"reflect"
)

type binding struct {
	container      *Container
	concrete       interface{}
	shared         bool
	sharedConcrete interface{}
}

func (b *binding) getConcrete(parameters ...interface{}) interface{} {
	spec := reflect.ValueOf(b.concrete)
	if spec.Kind() == reflect.Func {
		spec = spec.Call(makeArguments(parameters...))[0]
	}
	if spec.Kind() == reflect.Struct {
		return spec.Elem().Interface()
	}
	return spec.Interface()
}

func makeArguments(parameters ...interface{}) []reflect.Value {
	arguments := make([]reflect.Value, len(parameters))
	for index, parameter := range parameters {
		arguments[index] = reflect.ValueOf(parameter)
	}
	return arguments
}
