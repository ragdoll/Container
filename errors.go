package container

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrInterfaceMismatch            = errors.New("Concrete must implement abstract contract")
	ErrAbstractStructConcreteNotNil = errors.New("abstract<struct> wants concrete<nil>")
	ErrAbstractNotInvocable         = errors.New("Can only invoke function or struct method")
	ErrAliasAbstractMissing         = errors.New("trying to alias missing abstract")
)

type BindingMissingError struct {
	abstract interface{}
}

func (bme *BindingMissingError) Error() string {
	name := reflect.TypeOf(bme.abstract).Name()
	return fmt.Sprintf("Binding [%s] not found in container", name)
}
