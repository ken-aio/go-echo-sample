package main

import (
	"fmt"
	"log"
	"time"

	resty "gopkg.in/resty.v1"
)

// Req is request struct
type Req struct {
	Query string `json:"query"`
}

// Resp is response struct
type Resp struct {
	Success string `json:"success"`
	Message string `json:"message"`
	Param   string `json:"param"`
}

func main() {
	//simpleTest()
	//structBindTest()
	retryTest()
}

func simpleTest() {
	resp, _ := resty.R().Get("http://localhost:1314")
	fmt.Printf("Get resp = %+v\n", resp)
	resp, _ = resty.R().Post("http://localhost:1314")
	fmt.Printf("Post resp = %+v\n", resp)
	resp, _ = resty.R().Put("http://localhost:1314")
	fmt.Printf("Put resp = %+v\n", resp)
	resp, _ = resty.R().Delete("http://localhost:1314")
	fmt.Printf("Delete resp = %+v\n", resp)
}

func structBindTest() {
	resp, err := resty.
		R().
		SetQueryParam("query", "this is query param").
		Get("http://localhost:1314/param")
	if err != nil {
		log.Printf("get err: %+v", err)
	}
	fmt.Printf("Get resp = %+v\n", resp)

	resp, err = resty.
		R().
		SetBody(&Req{Query: "this is body param"}).
		Post("http://localhost:1314/param")
	if err != nil {
		log.Printf("get err: %+v", err)
	}
	fmt.Printf("Get resp = %+v\n", resp)
}

func retryTest() {
	log.Println("start")
	retryNum := 5
	resty.
		SetRetryCount(retryNum).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(5 * time.Second).
		R().
		Get("http://dummyurl")
	log.Println("end")
}
