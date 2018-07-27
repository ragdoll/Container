package container

import "reflect"

type internalBinding struct {
	container *Container
	concrete  interface{}
	shared    bool
}

func (ib *internalBinding) getConcrete(parameters ...interface{}) interface{} {
	spec := reflect.ValueOf(ib.concrete)
	if spec.Kind() == reflect.Func {
		spec = spec.Call(makeArguments(parameters...))[0]
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
