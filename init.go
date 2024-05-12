package main

import (
	"flag"
	"fmt"
	"os"
)

var version = flag.Bool("version", false, "测试")
var (
	Version   string
	Branch    string
	Commit    string
	BuildTime string
	GOOS      string
	GOARCH    string
)

func init() {
	flag.Parse()
	if *version {
		fmt.Println("|------------------------------ VERSION INFO -----------------------------|")
		fmt.Printf("VERSION: \t%s\n", Version)
		fmt.Printf("GOOS: \t\t%s\n", GOOS)
		fmt.Printf("GOARCH: \t%s\n", GOARCH)
		fmt.Printf("BRANCH: \t%s\n", Branch)
		fmt.Printf("COMMIT: \t%s\n", Commit)
		fmt.Printf("BUILDTIME: \t%s\n", BuildTime)
		fmt.Println("|-------------------------------------------------------------------------|")

		os.Exit(0)
	}
}
