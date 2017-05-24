package main

import (
	"io/ioutil"
	"log"
)

func main() {

	/*
	   Hızlıca Dosya Yazma (Quick Write to File)
	*/

	err := ioutil.WriteFile("demo.txt", []byte("Hi!\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
