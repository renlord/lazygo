package foo

import (
	"testing"
)

func TestFoo(t *testing.T) {
	if Foo() != "foo" {
		t.Fatalf("not foo")
	}
}
