package main

import (
	"github.com/grzegorz-zur/realize"
	"log"
)

var r realize.Realize

func main() {
	r.Sync = make(chan string)
	r.Settings.Read(&r)
	if r.Settings.FileLimit != 0 {
		if err := r.Settings.Flimit(); err != nil {
			log.Fatal(err)
		}
	}
	r.Start()
}
