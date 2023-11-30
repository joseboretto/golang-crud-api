package controller

import (
	"io"
	"net/http"
)

type HelloController struct {
}

func NewHelloController() *HelloController {
	return &HelloController{}
}

func (controller *HelloController) HelloWorld(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}
