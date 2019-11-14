package main

import (
	"fmt"
	membuff "github.com/orbs-network/membuffers/go/membufc/api"
	"os"
)

func main() {
	conf := membuff.NewConfig()
	conf.Language = "go"
	conf.Mock = true
	conf.LanguageGoCtx = true

	args := os.Args[1:]
	for _, arg := range args {
		conf.Files = append(conf.Files, arg)
	}

	err := membuff.Compile(conf)
	if err != nil {
		fmt.Printf("error %s\n", err.Error())
	}
}
