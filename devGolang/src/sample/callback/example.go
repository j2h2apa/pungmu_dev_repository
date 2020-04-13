package main

import (
	"fmt"
)

//HandlerFunc :
type HandlerFunc func()

//HandlersChain :
type HandlersChain []HandlerFunc

//IRoutes :
type IRoutes interface {
	POSTs(string, ...HandlerFunc) IRoutes
	POST(string, HandlerFunc) IRoutes
}

// CallbackEx :
type CallbackEx struct {
}

// POST :
func (c *CallbackEx) POST(relativePath string, handlers HandlerFunc) IRoutes {
	handlers()
	return c
}

//POSTs :
func (c *CallbackEx) POSTs(relativePath string, handlers ...HandlerFunc) IRoutes {
	//func (c *CallbackEx) POSTs(relativePath string, handlers HandlersChain) IRoutes {
	var k HandlersChain = handlers
	for _, f := range k {
		f()
	}
	f := k.Last()
	f()
	return c
}

//Last :
func (c HandlersChain) Last() HandlerFunc {
	if length := len(c); length > 0 {
		return c[length-1]
	}
	return nil
}

func main() {
	var c *CallbackEx = &CallbackEx{}
	fmt.Println("MAIN")
	c.POST("sample", callbackfn)
	c.POSTs("sample2", callbackfn, callbackfn2)
}

func callbackfn() {
	fmt.Println("CALL CALLBACK")
}

func callbackfn2() {
	fmt.Println("CALL CALLBACK2")
}
