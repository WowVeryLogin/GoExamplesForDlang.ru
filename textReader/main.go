package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var (
		readed []byte
		err    error
	)

	if len(os.Args) > 1 {
		readed, err = ioutil.ReadFile(os.Args[1])
	} else {
		readed, err = ioutil.ReadAll(os.Stdin)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes.ReplaceAll(readed, []byte("\n"), []byte(""))))
}
