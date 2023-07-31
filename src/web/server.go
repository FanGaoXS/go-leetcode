package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	s := newService()
	http.HandleFunc(Path, s.add())
	log.Printf("start http server on %v\n", HostAndPort)
	err := http.ListenAndServe(HostAndPort, nil)
	if err != nil {
		log.Fatalln("failed to start http server: ", err)
	}
}

type service struct {
	cache []string
}

func newService() *service {
	return &service{cache: []string{}}
}

func (s *service) add() func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		// Get请求的所有query
		values := req.URL.Query()

		// 根据param获取value
		str := values.Get(Param)
		log.Printf("get value: %v\n", str)

		// 追加到cache中
		s.cache = append(s.cache, str)

		// 转为json
		bytes, err := json.Marshal(s.cache)
		if err != nil {
			log.Println("failed to marshal to json.")
		}

		// 写入response的body
		result, err := fmt.Fprintln(res, string(bytes))
		if err != nil || result == 0 {
			log.Println("failed to write into body of response.")
		}
	}
}
