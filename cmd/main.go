package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type client struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Port     string    `json:"port"`
	Address  string    `json:"address"`
	check    checkLive `json:"check"`
	testCall string    `json:"testCall"`
}

type checkLive struct {
	interval int `json:"interval"`
	timeout  int `json:"timeout"`
}

func (c *client) call() *http.Response {
	resp, err := http.Get(c.Address + c.Port + c.testCall)
	if err != nil {
		log.Printf("err")
	}
	return resp
}
func getTest(w http.ResponseWriter, r *http.Request) {
	test := client{ID: "1", Name: "Bitcoincharts", Address: "http://api.bitcoincharts.com/v1/", Port: "", check: checkLive{interval: 10, timeout: 30}, testCall: "weighted_prices.json"}
	fmt.Printf("got /test request\n")
	asd := test.call()
	body, _ := ioutil.ReadAll(asd.Body)
	w.Write(body)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", getTest)

	err := http.ListenAndServe(":3333", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
