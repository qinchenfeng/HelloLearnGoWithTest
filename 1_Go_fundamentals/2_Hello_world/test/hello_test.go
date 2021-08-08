package test

import (
	"testing"

	"github.com/qinchenfeng/HelloLearnGoWithTest/1_Go_fundamentals/2_Hello_world/hello"
)

func TestHello(t *testing.T) {
	got := hello.Hello()
	want := "Hello, world"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}