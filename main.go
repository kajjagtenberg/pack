package main

import (
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
)

func init() {
	log.SetFormatter(&logrus.JSONFormatter{})
}

func main() {

}
