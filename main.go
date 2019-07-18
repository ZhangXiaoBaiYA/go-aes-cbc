package main

import (
	"test/utils"
)
import "log"
import "test/tutorials"

//记得引入此库
import "github.com/golang/protobuf/proto"

/**
protobuf + aes
*/
func main() {
	login := tutorials.LoginRequest{}
	login.Username = "zhangsan"
	login.Password = "12345678"
	login.Platform = "platform"
	log.Println(login)

	//将login变成字节数组
	out, err := proto.Marshal(&login)
	if err != nil {
		log.Fatalln(err.Error())
	}
	//加密
	encrypted, err2 := utils.Encrypt(out)
	log.Println(err2)
	log.Println(encrypted)

	s, err := utils.Request(encrypted, "/user/validate-password")
	if err != nil {
		log.Fatalln(err.Error())
	}

	//解密
	decrypted, err3 := utils.Decrypt(s)

	log.Println(decrypted)
	log.Println(err3)

	//使用proto对象解析
	loginResponse := &tutorials.Response{}
	decryptIntoProto := []byte(decrypted)
	proto.Unmarshal(decryptIntoProto, loginResponse)
	log.Println(loginResponse)
}
