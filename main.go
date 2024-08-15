package main

import (
	"github.com/adam-hanna/mock-test/mock"
)

func main() {
	mockedFoo := MockIFoo{}
	doBar(&mockedFoo)
}

func doBar(f mock.IFoo) {
	f.Bar()
}
