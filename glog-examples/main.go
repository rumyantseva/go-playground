package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
)

var (
	exit = flag.Bool("exit", false, "If true, will do glog.Exit (or glog.Fatal otherwise)")
)

// glog.Fatal vs glog.Exit
func main() {
	flag.Parse()
	myfunc()
}

func myfunc() {
	defer func() {
		fmt.Println("Yeah! I was deferred.")
	}()

	if *exit {
		glog.Exit("Trying glog.Exit...")
	} else {
		glog.Fatal("Trying glog.Fatal...")
	}
}
