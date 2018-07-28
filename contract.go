package container

// Contract is the interface that any container must implement
type Contract interface {
	Provide(abstract interface{}, concrete interface{}, shared bool)
	Singleton(abstract interface{}, concrete interface{})
	Make(abstract interface{}, parameters ...interface{}) (interface{}, error)
	Get(abstract interface{}) interface{}
	Has(abstract interface{}) bool
	Alias(abstract interface{}, alias string)
	Invoke(abstract interface{}) interface{}
	Flush()
}
