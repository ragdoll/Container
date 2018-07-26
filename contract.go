package container

// Contract is the interface that any container must implement
type Contract interface {
	Bind(key string, concrete interface{}, shared bool)
	Bound(key string) bool
	Find(key string) (interface{}, error)
	Get(key string) interface{}
	Singleton(key string, concrete interface{})
	Flush()
}
