package binding

import (
	"testing"
)

type dummyStruct struct{}

func TestBinding(t *testing.T) {
	t.Run("String concrete", func(t *testing.T) {
		b := Binding{Concrete: "nonsense"}
		if b.GetConcrete() != "nonsense" {
			t.Error("cannot get string concrete")
		}
	})

	t.Run("Int concrete", func(t *testing.T) {
		b := Binding{Concrete: 10}
		if b.GetConcrete() != 10 {
			t.Error("cannot get int concrete")
		}
	})

	t.Run("Float concrete", func(t *testing.T) {
		b := Binding{Concrete: 3.142}
		if b.GetConcrete() != 3.142 {
			t.Error("cannot get float concrete")
		}
	})

	t.Run("Boolean concrete", func(t *testing.T) {
		b := Binding{Concrete: true}
		if b.GetConcrete() != true {
			t.Error("cannot get boolean concrete")
		}
	})

	t.Run("Map concrete", func(t *testing.T) {
		b := Binding{Concrete: map[string]string{"key": "value"}}
		concrete := b.GetConcrete().(map[string]string)
		if concrete["key"] != "value" {
			t.Error("cannot get map concrete")
		}
	})

	t.Run("Function concrete", func(t *testing.T) {
		b := Binding{Concrete: func() string { return "nonsense" }}
		if b.GetConcrete() != "nonsense" {
			t.Error("cannot get function concrete")
		}

		b = Binding{Concrete: func(a, b int) int { return a + b }}
		if b.GetConcrete(10, 20) != 30 {
			t.Error("cannot get function concrete with parameters")
		}

		b = Binding{Concrete: func() (int, bool) { return 10, true }}
		if r := b.GetConcrete().([]interface{}); len(r) != 2 {
			t.Error("cannot return multiple function values")
		}

		b = Binding{Concrete: dummyStruct{}.Stub}
		if b.GetConcrete() != "stubbed" {
			t.Error("cannot get struct function concrete")
		}

		b = Binding{Concrete: dummyStruct{}.ParamStub}
		if b.GetConcrete(10, 20) != 30 {
			t.Error("cannot get struct function concrete with parameters")
		}
	})

	t.Run("Struct concrete", func(t *testing.T) {
		b := Binding{Concrete: dummyStruct{}}
		if b.GetConcrete().(dummyStruct).Stub() != "stubbed" {
			t.Error("cannot get struct concrete")
		}
	})
}

func (ds dummyStruct) Stub() string { return "stubbed" }

func (ds dummyStruct) ParamStub(a, b int) int { return a + b }
