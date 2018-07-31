package container

import (
	"go.rafdel.co/akisa/container/internal/pkg/binding"
	"go.rafdel.co/akisa/container/internal/pkg/utils"
)

// Container implements the container contract
type Container struct {
	bindings map[string]binding.Binding
	shared   map[string]interface{}
	aliases  map[string]string
}

// New creates a new container instance
func New() *Container {
	return &Container{
		bindings: make(map[string]binding.Binding),
		shared:   make(map[string]interface{}),
		aliases:  make(map[string]string),
	}
}

// Bind makes an abstract have a concrete when invoked from the container
func (c Container) Bind(abstract interface{}, concrete interface{}) {
	c.provide(abstract, concrete, false)
}

// BindShared makes an abstract have a concrete when invoked from the container
func (c Container) BindShared(abstract interface{}, concrete interface{}) {
	c.provide(abstract, concrete, true)
}

// Singleton makes an abstract have a concrete when invoked from the container
func (c Container) Singleton(abstract interface{}, concrete interface{}) {
	c.BindShared(abstract, concrete)
}

func (c Container) provide(abstract interface{}, concrete interface{}, shared bool) {
	var finalConcrete interface{}
	if utils.IsInterface(abstract) && !utils.IsImplements(concrete, abstract) {
		panic(ErrInterfaceMismatch)
	}
	if utils.IsStruct(abstract) {
		if concrete != nil {
			panic(ErrAbstractStructConcreteNotNil)
		}
		finalConcrete = abstract
	} else {
		finalConcrete = concrete
	}
	c.bindings[utils.GetKey(abstract)] = binding.New(finalConcrete, shared)
}

// Alias changes the name of the abstract
func (c Container) Alias(abstract interface{}, alias string) {
	if _, err := c.getBinding(abstract); err != nil {
		panic(ErrAliasAbstractMissing)
	}
	c.aliases[utils.GetKey(alias)] = utils.GetKey(abstract)
}

// Make finds an entry of the container by its identifier and returns it.
func (c Container) Make(abstract interface{}, parameters ...interface{}) (interface{}, error) {
	if shared, ok := c.shared[utils.GetKey(abstract)]; ok {
		return shared, nil
	}

	binding, err := c.getBinding(abstract)
	if err != nil {
		return nil, err
	}

	concrete := binding.GetConcrete(parameters...)
	if binding.Shared {
		c.shared[utils.GetKey(abstract)] = concrete
	}

	return concrete, nil
}

// // Invoke auto injects dependencies
// func (c *Container) Invoke(abstract interface{}) interface{}

// // Get finds a binding and returns the concretion or panic
// func (c *Container) Get(abstract interface{}) interface{}

// Has determine if the given key type has been bound.
func (c Container) Has(abstract interface{}) bool {
	_, err := c.getBinding(abstract)
	return err == nil
}

func (c Container) getBinding(abstract interface{}) (binding.Binding, error) {
	key := utils.GetKey(abstract)
	if alias, found := c.aliases[key]; found {
		key = alias
	}
	if binding, found := c.bindings[key]; found {
		return binding, nil
	}
	return binding.Binding{}, &BindingMissingError{abstract}
}
