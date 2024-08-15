package main

import (
	"github.com/adam-hanna/mock-test/foo"
)

func main() {
	mockedFoo := MockIFoo{}
	doBar(&mockedFoo)
}

func doBar(f foo.IFoo) {
	f.Bar()
}
