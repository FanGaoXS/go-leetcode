package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var input string
	for {
		fmt.Print("request: ")
		// get value from scan
		n, err := fmt.Scanln(&input)
		if err != nil || n == 0 {
			log.Fatalln("failed to get value from scan:", err)
		}

		// cell http client
		strs, err := BcjClient(input)
		if err != nil {
			log.Fatalln("failed to cell BcjClient:", err)
		}
		fmt.Println("response:", strs)
	}
}

func BcjClient(str string) ([]string, error) {
	// 利用http client发起get请求
	// localhost:8080/test?str=5
	res, err := http.Get("http://" + HostAndPort + Path + "?" + Param + "=" + str)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// 从response的body中读内容
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var strs []string
	// 从json转为object
	err = json.Unmarshal(bytes, &strs)
	if err != nil {
		log.Println("failed to unmarshal to object.")
	}

	return strs, nil
}
