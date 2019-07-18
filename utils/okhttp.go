package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"test/config"
)

func Request(text, url string) (string, error) {

	//请求客户端
	client := &http.Client{}

	req, e := http.NewRequest("POST", config.BASE_URL+url, strings.NewReader(text))
	if e != nil {
		log.Fatalln(e.Error())
	}

	req.Header.Set("Content-Type", "application/x-protobuf")
	req.Header.Set("TRADING_CHAIN_PLATFORM", config.REQ_PLATFORM)

	response, er := client.Do(req)
	if er != nil {
		log.Fatalln(er.Error())
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body), nil
}
