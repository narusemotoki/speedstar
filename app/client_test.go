package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"testing"
)

var URL = "http://localhost:8080"
var BODY_TYPE = "application/json"

type Part struct {
	Num      float32 `json:"num"`
	Operator string  `json:"operator"`
}

type Request struct {
	Parts []Part `json:"parts"`
}

type Result struct {
	Num float32 `json:"num"`
}

type Response struct {
	Result Result `json:"result"`
}

type Expect struct {
	Num     float32
	Request Request
}

var expects []Expect = []Expect{
	Expect{2, Request{[]Part{
		Part{2, "add"},
		Part{1, "sub"},
	}}},
	Expect{55, Request{[]Part{
		Part{2, "add"},
		Part{3, "add"},
		Part{4, "add"},
		Part{5, "add"},
		Part{6, "add"},
		Part{7, "add"},
		Part{8, "add"},
		Part{9, "add"},
		Part{10, "add"},
	}}},
	Expect{-9.9, Request{[]Part{
		Part{2, "sub"},
		Part{-1, "multi"},
		Part{-0.1, "div"},
		Part{-0.1, "sub"},
	}}},
}

func _post(expect Expect) {
	j, _ := json.Marshal(expect.Request)
	res, err := http.Post(URL, BODY_TYPE, strings.NewReader(string(j)))
	if err == nil {
		b, _ := ioutil.ReadAll(res.Body)
		var res Response
		json.Unmarshal(b, &res)
		if expect.Num != res.Result.Num {
			fmt.Println("Response is wrong:", expect.Request, "->", string(b))
		}
	} else {
		fmt.Println(err)
	}
}

func post(wg *sync.WaitGroup) {
	for _, expect := range expects {
		_post(expect)
	}
	wg.Done()
}

func Benchmark(b *testing.B) {
	b.N = 100
	for i := 0; i < b.N; i++ {
		forTest(i)
	}
}

func forTest(i int) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go post(&wg)
	}
	wg.Wait()
	fmt.Println(i)
}
