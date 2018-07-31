package utils

import "reflect"

func GetKey(a interface{}) string {
	abstract := reflect.TypeOf(a)
	switch {
	case IsImplements(a, new(reflect.Type)):
		abstract = a.(reflect.Type)
	case IsInterface(a):
		abstract = abstract.Elem()
	case abstract.Kind() == reflect.String:
		return a.(string)
	}
	return abstract.PkgPath() + "." + abstract.Name()
}

func IsInterface(abstract interface{}) bool {
	t := reflect.TypeOf(abstract)
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Interface
}

func IsFunc(abstract interface{}) bool {
	return reflect.TypeOf(abstract).Kind() == reflect.Func
}

func IsStruct(abstract interface{}) bool {
	return reflect.TypeOf(abstract).Kind() == reflect.Struct
}

func IsImplements(concrete interface{}, abstract interface{}) bool {
	return reflect.TypeOf(concrete).Implements(reflect.TypeOf(abstract).Elem())
}
