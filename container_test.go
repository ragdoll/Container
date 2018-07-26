package container

import (
	"math/rand"
	"reflect"
	"strings"
	"testing"
	"time"
)

var (
	_containerType = reflect.TypeOf((*Container)(nil))
	_contractType  = reflect.TypeOf((*Contract)(nil))
)

func TestContainer(t *testing.T) {
	t.Run("Implement Contract", func(t *testing.T) {
		if _containerType.Implements(_contractType.Elem()) != true {
			t.Errorf("Container doesn't implement Contract")
		}
	})

	t.Run("Binding", func(t *testing.T) {
		c := New()
		c.Bind("stuff", "nonsense", false)
		if c.Bound("stuff") == false {
			t.Errorf("cannot bind item to container")
		}
	})

	t.Run("Retrieve from container", func(t *testing.T) {
		c := New()
		c.Bind("stuff", "nonsense", false)

		if c.Get("stuff") != "nonsense" {
			t.Errorf("cannot get item in container")
		}
	})

	t.Run("Bound function gets called", func(t *testing.T) {
		c := New()
		c.Singleton("stuff", func() string {
			return "nonsense"
		})

		if c.Get("stuff") != "nonsense" {
			t.Errorf("resolving function bound to key should automatically be invoked")
		}
	})

	t.Run("Singleton or shared item in container", func(t *testing.T) {
		c := New()
		c.Bind("random", dummyRandomNumber, true)
		if c.Get("random") != c.Get("random") {
			t.Errorf("cannot bind shared item to container")
		}

		c.Singleton("check", dummyRandomNumber)
		if c.Get("check") != c.Get("check") {
			t.Errorf("Singleton should only be instantiated once")
		}
	})

	t.Run("Make dependencies", func(t *testing.T) {
		c := New()
		c.Singleton("stuff", func(t string) string { return strings.ToUpper(t) })
		if v, _ := c.Make("stuff", "nonsense"); v != "NONSENSE" {
			t.Errorf("cannot make bindings with arguments")
		}
	})

	t.Run("Empty container", func(t *testing.T) {
		c := New()
		c.Singleton("stuff", "nonsense")
		if c.Flush(); c.Bound("cli") {
			t.Errorf("cannot flush container")
		}
	})
}

func dummyRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000000)
}
