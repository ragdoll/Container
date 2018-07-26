package container

import "fmt"

// Container implements the container contract
type Container struct {
	bindings map[string]binding
	shared   map[string]interface{}
}

// Bind register a binding with the container.
func (c *Container) Bind(key string, value interface{}, shared bool) {
	c.dropStaleInstance(key)
	c.bindings[key] = binding{
		container: c,
		concrete:  value,
		shared:    shared,
	}
}

// Make finds an entry of the container by its identifier and returns it.
func (c *Container) Make(key string, parameters ...interface{}) (interface{}, error) {
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

// Get finds a binding and returns the concretion or panics
func (c *Container) Get(key string) interface{} {
	binding, err := c.Make(key)
	if err != nil {
		panic(err)
	}
	return binding
}

// Bound determine if the given key type has been bound.
func (c *Container) Bound(key string) bool {
	_, ok := c.bindings[key]
	return ok
}

// Singleton register a shared binding in the container.
func (c *Container) Singleton(key string, value interface{}) {
	c.Bind(key, value, true)
}

// Flush remove all bindings from Container
func (c *Container) Flush() {
	c.bindings = make(map[string]binding)
	c.shared = make(map[string]interface{})
}

func (c *Container) dropStaleInstance(key string) {
	delete(c.bindings, key)
	delete(c.shared, key)
}
