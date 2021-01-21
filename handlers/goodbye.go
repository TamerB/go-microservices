package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("See you later\n"))
	g.l.Print("see you later")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		g.l.Print(err)
		http.Error(w, "Oops", http.StatusBadRequest)
		return
	}
	g.l.Printf("Data %s", d)
	fmt.Fprintf(w, "See you later %s\n", d)
}
