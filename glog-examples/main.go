package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
	"github.com/sirupsen/logrus"
)

var (
	exit      = flag.Bool("exit", false, "If true, will do glog.Exit (or glog.Fatal otherwise)")
	useLogrus = flag.Bool("logrus", false, "If true, will use logrus (or glog otherwise)")
)

// This application shows how defer work with logger.Fatal or logger.Error.
// Explanation: in case of logger.Fatal/logger.Error defer will not be done
// (application will os.Exit immediately).

// How to run this app:
// try logrus.Exit: go run main.go --logrus --exit
// try logrus.Fatal: go run main.go --logrus
// try glog.Exit: go run main.go --exit
// try glog.Fatal: go run main.go
func main() {
	flag.Parse()
	myfunc()
}

func myfunc() {
	defer func() {
		fmt.Println("Yeah! I was deferred.")
	}()

	if *useLogrus {
		if *exit {
			logrus.Exit(255)
		} else {
			logrus.Fatal("Trying glog.Exit...")
		}
	} else {
		if *exit {
			glog.Exit("Trying glog.Exit...")
		} else {
			glog.Fatal("Trying glog.Fatal...")
		}
	}
}
