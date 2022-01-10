package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("----------------Welcome-------------\n ")
	var ch int
	var choice int
	for true {
		fmt.Println("1 For Operations\n2 For exit")
		fmt.Scanln(&ch)
		if ch == 1 {
			fmt.Println("\n-------------------Enter the file name----------\n ")
			var filename string
			fmt.Scanln(&filename)
			fmt.Println("\n\n 1 For base64 encryption and decryption\n 2 For sha256 hash generation\n 3 For sha512 hash generation\n 4 For md5 hash generation\n 5 For validation \n 6 For Exit")
			fmt.Println("\n-----------------Enter choice-------------\n ")
			fmt.Scanln(&choice)
			switch choice {
			case 1:
				var ch1 int
				fmt.Println("1 For Encryption\n2 For Decryption\n--------------------------\nEnter Choice:-------------------------")
				fmt.Scanln(&ch1)
				switch ch1 {
				case 1:
					encodeBase64(filename)
				case 2:
					decodeToStr(filename + "_b64.txt")
				default:
					fmt.Println("--------------Wrong choice-------------------")
				}

			case 2:
				hsha256(filename)
				fmt.Println("\n\n----------------------------\nHash Generated\n-----------------------------\n\n ")
			case 3:
				hsha512(filename)
				fmt.Println("\n\n----------------------------\nHash Generated\n------------------------------\n\n ")
			case 4:
				hmd5(filename)
				fmt.Println("\n\n----------------------------\nHash Generated\n------------------------------\n\n ")
			case 5:
				fmt.Println("Check For\n 1:    sha256\n 2:   sha512\n 3:    md5")
				var choice1 int
				fmt.Println("\n-----------Enter choice-----------")
				fmt.Scanln(&choice1)
				switch choice1 {
				case 1:
					calc_sha(filename, filename+"_hash256.txt", 1)
				case 2:
					calc_sha(filename, filename+"_hash512.txt", 2)
				case 3:
					calc_sha(filename, filename+"_md5.txt", 3)
				default:
					fmt.Println("--------------------Wrong choice-----------------------")
				}

			case 6:
				os.Exit(0)
			default:
				fmt.Println("---------------------Wrong choice----------------------")
			}

		} else {
			os.Exit(0)
		}

	}

}
