package container

// New creates a new container instance
func New() *Container {
	return &Container{
		bindings: make(map[string]binding),
		shared:   make(map[string]interface{}),
	}
}
