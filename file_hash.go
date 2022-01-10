package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"
)

//function to calculate sha256 checksum
func hsha256(filename string) {
	dataBytes, err := ioutil.ReadFile(filename) //read the content of the file directly into byte array
	checkNilError(err)                          //check for any possible error in reading operation

	//computing the checksum and converting into a string to store in a file
	h := sha256.New()
	h.Write(dataBytes)
	hash_data := hex.EncodeToString(h.Sum(nil))

	//storing the hash value in a file
	file, err := os.Create(filename + "_hash256.txt")
	checkNilError(err)
	defer file.Close()
	io.WriteString(file, string(hash_data))
}

//function to calculate sha512 checksum
func hsha512(filename string) {
	dataBytes, err := ioutil.ReadFile(filename)
	checkNilError(err)
	h := sha512.New()
	h.Write(dataBytes)
	hash_data := hex.EncodeToString(h.Sum(nil))
	file, err := os.Create(filename + "_hash512.txt")
	checkNilError(err)
	defer file.Close()
	io.WriteString(file, string(hash_data))
}

//function to calculate md5 checksum
func hmd5(filename string) {
	dataBytes, err := ioutil.ReadFile(filename)
	checkNilError(err)
	md5 := md5.New()
	md5.Write(dataBytes)
	hash_data := hex.EncodeToString(md5.Sum(nil))
	file, err := os.Create(filename + "_md5.txt")
	checkNilError(err)
	defer file.Close()
	io.WriteString(file, string(hash_data))
}

//function to check for errors while read/write operations
func checkNilError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
