package main

import (
	"fmt"
	"github.com/starine/gonote/note"
)

func sayHello()  {
	fmt.Println("hello world")
}

func main() {
	sayHello()
	note.SayNote()
}
