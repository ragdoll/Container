package container

// Contract is the interface that any container must implement
type Contract interface {
	Provide(abstract interface{}, concrete interface{}, shared bool)
	Has(abstract interface{}) bool
	Make(abstract interface{}, parameters ...interface{}) (interface{}, error)
	Singleton(abstract interface{}, concrete interface{})
	Get(abstract interface{}) interface{}
	Alias(abstract interface{}, alias string)
	Invoke(abstract interface{}) interface{}
	Flush()
}
