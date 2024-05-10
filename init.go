package main

import (
	"flag"
)

var testis = flag.Bool("test", false, "测试")

func init() {
	flag.Parse()
}
