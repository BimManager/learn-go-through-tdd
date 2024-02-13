package hello_test

import (
	"github.com/bimmanager/learn-go-through-tdd/hello"
	"testing"
)

func TestSayHello(t *testing.T) {
	actual := hello.GreetWorld()
	expected := "Hello, World!"
	if actual != expected {
		t.Errorf("expected: %q, got: %q", expected, actual)
	}
}
