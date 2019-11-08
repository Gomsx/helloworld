package helloworld

import "testing"

func TestHelloWorld(t *testing.T) {
	var hw HelloWorld
	var want = "hello\nworld\n"
	if got := hw.Print(); got != want {
		t.Errorf("hw.Print() = %v,want %v", got, want)
	}
}
