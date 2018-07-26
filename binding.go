package container

import "reflect"

type binding struct {
	container      *Container
	concrete       interface{}
	shared         bool
	sharedConcrete interface{}
}

func (b *binding) getConcrete(parameters ...interface{}) interface{} {
	spec := reflect.ValueOf(b.concrete)
	if spec.Kind() == reflect.Struct {
		return spec.Elem().Interface()
	}
	if spec.Kind() != reflect.Func {
		return spec.Interface()
	}

	concrete := spec.Call(nil)[0]
	return b.maybeReturnShared(concrete)
}

func (b *binding) maybeReturnShared(concrete reflect.Value) interface{} {
	switch {
	case b.shared && b.sharedConcrete == nil:
		b.sharedConcrete = concrete.Interface()
		fallthrough
	case b.shared:
		return b.sharedConcrete
	default:
		return concrete.Interface()
	}
}
