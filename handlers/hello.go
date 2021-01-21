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

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello there\n"))
	h.l.Print("Hello world")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		h.l.Print(err)
		http.Error(w, "Oops", http.StatusBadRequest)
		return
	}
	h.l.Printf("Data %s", d)
	fmt.Fprintf(w, "Hello %s\n", d)
}
