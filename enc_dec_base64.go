package main

import (
	b64 "encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func encodeBase64(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	sEnc := b64.StdEncoding.EncodeToString(file)
	fmt.Println("\n\n---------------------File Encoded---------------------\n\n ")
	fmt.Println("----------------Encoded String-----------------\n ", sEnc)
	file1, err := os.Create(filename + "_b64.txt")
	if err != nil {
		panic(err)
	}
	io.WriteString(file1, string(sEnc))
	file1.Close()
}

func decodeToStr(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	sDec, _ := b64.StdEncoding.DecodeString(string(file))
	fmt.Println("-------------------Decoded string ------------------")
	fmt.Println(string(sDec) + "\n\n")
}
