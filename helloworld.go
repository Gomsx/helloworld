package helloworld

import (
	"fmt"

	"github.com/gomsx/hello"
	"github.com/gomsx/world/v2"
)

type HelloWorld struct {
	h hello.Hello
	w world.World
}

func (hw HelloWorld) Print() string {
	s := hw.h.Print() + hw.w.Print()
	fmt.Println(s)
	return s
}
