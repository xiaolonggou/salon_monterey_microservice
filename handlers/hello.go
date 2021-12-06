package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("hello world")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "ooh ooops", http.StatusBadRequest)
		return
	}
	h.l.Println(fmt.Sprintf("serving %s", d))

	fmt.Fprintf(rw, "hey %s \n", d)
}