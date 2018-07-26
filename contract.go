package container

// Contract is the interface that any container must implement
type Contract interface {
	Bind(key string, concrete interface{}, shared bool)
	Make(key string, parameters ...interface{}) (interface{}, error)
	Get(key string) interface{}
	Bound(key string) bool
	Singleton(key string, concrete interface{})
	Flush()
}
