package container

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

type dummyInterface interface {
	Stub() string
	ParamStub(text string) string
	InvokeStub(d dummyInterface) string
}

type dummyStruct struct{}

func TestImplementation(t *testing.T) {
	t.Run("Implement Contract", func(t *testing.T) {
		_containerType := reflect.TypeOf((*Container)(nil))
		_contractType := reflect.TypeOf((*Contract)(nil))
		if _containerType.Implements(_contractType.Elem()) != true {
			t.Error("Container doesn't implement Contract")
		}
	})
}

func TestProvide(t *testing.T) {
	t.Run("Abstract must be <string|interface|struct>", func(t *testing.T) {
		c := New()
		c.Provide("stuff", "nonsense", false)
		if c.Has("stuff") == false {
			t.Error("cannot bind stuff to container")
		}

		c.Provide(new(dummyInterface), dummyStruct{}, false)
		if c.Has(new(dummyInterface)) == false {
			t.Error("cannot use interface as abstract to bind concrete")
		}

		c.Provide(dummyStruct{}, nil, false)
		if c.Has(dummyStruct{}) == false {
			t.Error("abstract of type struct and nil concrete should return abstract")
		}

		defer func() {
			if recover() == nil {
				t.Error("abstract can only be of type <string|interface|struct>")
			}
		}()
		c.Provide([]string{}, "nonsense", false)
	})

	t.Run("Concrete of abstract<interface> implements interface", func(t *testing.T) {
		defer func() {
			if recover() == nil {
				t.Error("concrete of abstract<interface> must implement interface")
			}
		}()
		New().Provide(new(dummyInterface), "nonsense", false)
	})

	t.Run("abstract<struct> wants concrete<nil>", func(t *testing.T) {
		defer func() {
			if recover() == nil {
				t.Error("abstract<struct> wants concrete<nil>")
			}
		}()
		New().Provide(dummyStruct{}, "nonsense", false)
	})
}

func TestMake(t *testing.T) {
	t.Run("Make binding", func(t *testing.T) {
		c := New()
		c.Provide("string", "lorem", false)
		if value, _ := c.Make("string"); value != "lorem" {
			t.Error("can't make string from container")
		}

		c.Provide("integer", 90210, false)
		if value, _ := c.Make("integer"); value != 90210 {
			t.Error("can't make integer from container")
		}

		c.Provide("boolean", true, false)
		if value, _ := c.Make("boolean"); value != true {
			t.Error("can't make boolean from container")
		}

		c.Provide("map", map[string]string{"key": "value"}, false)
		if value, _ := c.Make("map"); value.(map[string]string)["key"] != "value" {
			t.Error("can't make map from container")
		}

		c.Provide("function", func() string { return "return" }, false)
		if value, _ := c.Make("function"); value != "return" {
			t.Error("can't make function from container")
		}

		c.Provide("struct", dummyStruct{}, false)
		if value, _ := c.Make("struct"); value.(dummyStruct).Stub() != "stub" {
			t.Error("can't make struct from container")
		}

		c.Provide(new(dummyInterface), dummyStruct{}, false)
		if value, _ := c.Make(new(dummyInterface)); value.(dummyInterface).Stub() != "stub" {
			t.Error("can't make new(dummyInterface) from container")
		}

		c.Provide(dummyStruct{}, nil, false)
		if value, _ := c.Make(dummyStruct{}); value.(dummyStruct).Stub() != "stub" {
			t.Error("can't make dummyStruct{} from container")
		}
	})

	t.Run("Make with parameters", func(t *testing.T) {
		c := New()
		c.Provide("stuff", func(a, b int) int { return a + b }, false)
		if value, _ := c.Make("stuff", 10, 5); value != 15 {
			t.Error("can't make function with parameters")
		}
		c.Flush()

		c.Provide("stuff", dummyStruct{}.ParamStub, false)
		if value, _ := c.Make("stuff", "nonsense"); value != "nonsense" {
			t.Error("can't make struct method with parameters")
		}
	})

	t.Run("Make shared binding", func(t *testing.T) {
		c := New()
		c.Provide("stuff", dummyRandomNumber, true)
		c.Singleton("other", dummyRandomNumber)
		if c.Get("stuff") != c.Get("stuff") && c.Get("other") != c.Get("other") {
			t.Error("Shared bindings should only be instantiated once")
		}
	})
}

func TestGet(t *testing.T) {
	t.Run("Panic", func(t *testing.T) {
		defer func() {
			if recover() == nil {
				t.Error("should panic if cannot get from container")
			}
		}()
		New().Get("stuff")
	})
}

func TestAlias(t *testing.T) {
	t.Run("Alias abstract", func(t *testing.T) {
		c := New()
		c.Singleton(new(dummyInterface), dummyStruct{})
		c.Alias(new(dummyInterface), "dummy")
		if c.Get("dummy").(dummyInterface).Stub() != "stub" {
			t.Error("cannot alias abstract")
		}
		if c.Get(new(dummyInterface)).(dummyInterface).Stub() != "stub" {
			t.Error("aliased abstract should still be accessible")
		}
	})

	t.Run("Flush removes alias", func(t *testing.T) {
		defer func() {
			if recover() == nil {
				t.Error("flushing does not remove alias")
			}
		}()

		c := New()
		c.Singleton(new(dummyInterface), dummyStruct{})
		c.Alias(new(dummyInterface), "dummy")
		c.Flush()
		c.Get("dummy")
	})
}

func TestInvoke(t *testing.T) {
	t.Run("Invoke only <func> or <struct>.<func>", func(t *testing.T) {
		defer func() {
			if recover() == nil {
				t.Error("should only invoke function or struct method")
			}
		}()

		c := New()
		c.Singleton(new(dummyInterface), dummyStruct{})
		c.Invoke("string")
	})

	t.Run("Invoke", func(t *testing.T) {
		c := New()
		c.Singleton(new(dummyInterface), dummyStruct{})
		counter := c.Invoke(func(a dummyInterface, e dummyInterface) int {
			return 10
		})
		if counter != 10 {
			t.Error("cannot invoke function")
		}
	})

	t.Run("Invoke struct method", func(t *testing.T) {
		c := New()
		c.Singleton(new(dummyInterface), dummyStruct{})
		stuff := c.Invoke(dummyStruct{}.InvokeStub)
		if stuff != "nonsense" {
			t.Error("cannot invoke struct method")
		}
	})

	// t.Run("Auto-make Get dependencies", func(t *testing.T) {
	// 	c := New()
	// 	c.Singleton(new(dummyInterface), dummyStruct{})
	// 	c.Singleton("stuff", func(d dummyInterface) string { return "nonsense" })
	// 	if c.Get("stuff") != "nonsense" {
	// 		t.Errorf("cannot auto-make Get dependencies")
	// 	}
	// })
}

func TestFlush(t *testing.T) {
	t.Run("Empty Container", func(t *testing.T) {
		c := New()
		c.Provide("stuff", "nonsense", false)
		c.Flush()
		if c.Has("stuff") {
			t.Error("can't flush container")
		}
	})
}

func (ds dummyStruct) Stub() string { return "stub" }

func (ds dummyStruct) ParamStub(text string) string { return text }

func (ds dummyStruct) InvokeStub(d dummyInterface) string { return "nonsense" }

func dummyRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100000)
}
