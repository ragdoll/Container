package container

import (
	"fmt"
	"reflect"
)

// Container implements the container contract
type Container struct {
	bindings map[string]internalBinding
	shared   map[string]interface{}
}

// New creates a new container instance
func New() *Container {
	c := &Container{}
	c.Flush()
	return c
}

// Provide makes an abstract have a concrete when invoked from the container
func (c *Container) Provide(abstract interface{}, concrete interface{}, shared bool) {
	var finalConcrete interface{}
	_tOA := reflect.TypeOf(abstract)
	_tOC := reflect.TypeOf(concrete)
	if isInterface(_tOA) && !_tOC.Implements(_tOA.Elem()) {
		panic("Concrete must implement abstract contract")
	}
	if _tOA.Kind() == reflect.Struct {
		if concrete != nil {
			panic("abstract<struct> wants concrete<nil>")
		}
		finalConcrete = abstract
	} else {
		finalConcrete = concrete
	}

	c.bindings[getKey(abstract)] = internalBinding{c, finalConcrete, shared}
}

// Has determine if the given key type has been bound.
func (c *Container) Has(abstract interface{}) bool {
	_, found := c.bindings[getKey(abstract)]
	return found
}

// Make finds an entry of the container by its identifier and returns it.
func (c *Container) Make(abstract interface{}, parameters ...interface{}) (interface{}, error) {
	key := getKey(abstract)
	if concrete, ok := c.shared[key]; ok {
		return concrete, nil
	}
	if binding, ok := c.bindings[key]; ok {
		concrete := binding.getConcrete(parameters...)
		if binding.shared {
			c.shared[key] = concrete
		}
		return concrete, nil
	}
	return nil, fmt.Errorf("Binding [%s] not found in container", key)
}

// Singleton register a shared binding in the container.
func (c *Container) Singleton(abstract interface{}, concrete interface{}) {
	c.Provide(abstract, concrete, true)
}

// Get finds a binding and returns the concretion or panic
func (c *Container) Get(key interface{}) interface{} {
	binding, err := c.Make(key)
	if err != nil {
		panic(err)
	}
	return binding
}

// Flush remove all bindings from Container
func (c *Container) Flush() {
	c.bindings = make(map[string]internalBinding)
	c.shared = make(map[string]interface{})
}

func isInterface(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Interface
}

func getKey(abstract interface{}) string {
	a := reflect.TypeOf(abstract)
	switch {
	case a.Kind() == reflect.String:
		return abstract.(string)
	case isInterface(a):
		return makeName(a.Elem())
	case a.Kind() == reflect.Struct:
		return makeName(a)
	}
	panic("Abstract much be a string, struct or interface")
}

func makeName(abstract reflect.Type) string {
	return abstract.PkgPath() + "." + abstract.Name()
}
