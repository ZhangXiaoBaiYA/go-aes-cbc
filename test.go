package main

import "fmt"
import "./utils"
import "log"

func main() {
	fmt.Println("abc")
	to_encrypt := []byte("i want an apple")
	aa, e := utils.Encrypt(to_encrypt)
	log.Println(aa)
	log.Println(e)

	//解密
	bb, ee := utils.Decrypt(aa)
	log.Println(bb)
	log.Println(ee)
}
