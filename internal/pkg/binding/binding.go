package binding

import (
	"reflect"
)

// Binding struct
type Binding struct {
	Concrete interface{}
	Shared   bool
}

// New creates a new Binding instance
func New(concrete interface{}, shared bool) Binding {
	return Binding{concrete, shared}
}

// GetConcrete returns binding concrete. If binding concrete is function,
// the binding is invoked with or without parameters.
func (b *Binding) GetConcrete(parameters ...interface{}) interface{} {
	spec := reflect.ValueOf(b.Concrete)
	if spec.Kind() == reflect.Func {
		values := spec.Call(parameterValues(parameters...))
		return interfaceReturnValues(values)
	}
	return spec.Interface()
}

func parameterValues(parameters ...interface{}) []reflect.Value {
	arguments := make([]reflect.Value, len(parameters))
	for index, parameter := range parameters {
		arguments[index] = reflect.ValueOf(parameter)
	}
	return arguments
}

func interfaceReturnValues(called []reflect.Value) interface{} {
	args := make([]interface{}, len(called))
	for i, value := range called {
		args[i] = value.Interface()
	}
	if len(args) == 1 {
		return args[0]
	}
	if len(args) > 1 {
		return args
	}
	return nil
}
