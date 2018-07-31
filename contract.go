package container

// Contract is the interface that any container must implement
type Contract interface {
	Bind(abstract interface{}, concrete interface{})
	BindShared(abstract interface{}, concrete interface{})
	Singleton(abstract interface{}, concrete interface{})

	Alias(abstract interface{}, alias string)

	Make(abstract interface{}, parameters ...interface{}) (interface{}, error)
	// Invoke(abstract interface{}) interface{}
	// Get(abstract interface{}) interface{}

	Has(abstract interface{}) bool
}
