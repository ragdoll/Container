package container

import (
	"math/rand"
	"testing"
	"time"

	"go.rafdel.co/akisa/container/internal/pkg/utils"
)

func TestContract(t *testing.T) {
	t.Run("Implement Contract", func(t *testing.T) {
		if !utils.IsImplements(Container{}, new(Contract)) {
			t.Error("Container doesn't implement Contract")
		}
	})
}

func TestBind(t *testing.T) {
	t.Run("can bind", func(t *testing.T) {
		c := New()
		if c.Bind("stuff", "nonsense"); !c.Has("stuff") {
			t.Error("cannot bind stuff to the container")
		}
	})

	t.Run("concrete of abstract<interface> implements interface", func(t *testing.T) {
		defer func() {
			if recover() == nil {
				t.Error("concrete of abstract<interface> must implement interface")
			}
		}()
		New().Bind(new(Contract), "nonsense")
	})

	t.Run("abstract<struct> wants concrete<nil>", func(t *testing.T) {
		defer func() {
			if recover() == nil {
				t.Error("abstract<struct> wants concrete<nil>")
			}
		}()
		New().Bind(Container{}, "nonsense")
	})

	t.Run("shared", func(t *testing.T) {
		c := New()
		c.BindShared("stuff", "nonsense")
		if c.Has("stuff") == false {
			t.Error("shared binding not binding")
		}

		c.Singleton("singleton", "nonsense")
		if c.Has("singleton") == false {
			t.Error("singleton helper function not binding")
		}
	})
}

func TestAlias(t *testing.T) {
	t.Run("can alias", func(t *testing.T) {
		c := New()
		c.Bind(new(Contract), Container{})
		if c.Alias(new(Contract), "alias"); c.Has("alias") == false {
			t.Error("cannot alias abstract in container")
		}
	})

	t.Run("alias must alias only existing abstract", func(t *testing.T) {
		defer func() {
			if recover() == nil {
				t.Error("does not account for trying to alias missing abstract")
			}
		}()
		New().Alias(new(Contract), "alias")
	})
}

func TestMake(t *testing.T) {
	t.Run("can make concrete<interface{}> from abstract", func(t *testing.T) {
		c := New()
		c.Bind(new(Contract), Container{})
		if _, err := c.Make(new(Contract)); err != nil {
			t.Error("cannot make binding from container")
		}
	})

	t.Run("can make concrete<interface{}> with alias", func(t *testing.T) {
		c := New()
		c.Bind(new(Contract), Container{})
		c.Alias(new(Contract), "container")
		if _, err := c.Make("container"); err != nil {
			t.Error("cannot make concrete<interface{}> with alias")
		}
	})

	t.Run("call concrete<func> during make", func(t *testing.T) {
		c := New()
		c.Bind("stuff", func() string { return "nonsense" })
		if value, _ := c.Make("stuff"); value != "nonsense" {
			t.Error("does not call concrete<func> during make")
		}
	})

	t.Run("call concrete<func> during make with parameters", func(t *testing.T) {
		c := New()
		c.Bind("stuff", func(a, b int) int { return a + b })
		if value, _ := c.Make("stuff", 10, 15); value != 25 {
			t.Error("does not call concrete<func> during make with parameters")
		}
	})

	t.Run("can make shared binding", func(t *testing.T) {
		c := New()
		c.BindShared("stuff", dummyRandomNumber)
		a1, _ := c.Make("stuff")
		a2, _ := c.Make("stuff")
		if a1 != a2 {
			t.Error("cannot make shared binding")
		}

		c.BindShared("stuff", dummyRandomNumber)
		b1, _ := c.Make("stuff")
		b2, _ := c.Make("stuff")
		if b1 != b2 {
			t.Error("cannot make shared binding via singleton")
		}
	})
}

func TestHas(t *testing.T) {
	t.Run("find binding", func(t *testing.T) {
		c := New()
		c.Bind("stuff", "nonsense")

		if c.Has("stuff") == false {
			t.Error("container should have binding [stuff]")
		}

		if c.Has("notstuff") == true {
			t.Error("container should not have binding [notstuff]")
		}
	})
}

func dummyRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100000)
}
