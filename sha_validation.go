package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func calc_sha(fpath, hash_path string, choice int) {
	//creating the sha256 hash
	f, err := os.Open(fpath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	var h hash.Hash
	if choice == 1 {
		h = sha256.New()
	} else if choice == 2 {
		h = sha512.New()
	} else if choice == 3 {
		h = md5.New()
	}

	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	hash := h.Sum(nil)
	var s string = hex.EncodeToString(hash[:])

	//reading the hash file to compare
	file, err := ioutil.ReadFile(hash_path)
	if err != nil {
		panic(err)
	}
	if s == string(file) {
		fmt.Println("\n\n\n--------------------\nFile integrity is maintained\n--------------------\n\n ")
	} else {
		fmt.Println("\n\n\n--------------------\nFile has been damaged\n--------------------\n\n ")
	}
}
