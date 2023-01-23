package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"servicediscoverer/dev"
	"servicediscoverer/ent"
	"servicediscoverer/ent/providerregisterdata"
	service2 "servicediscoverer/service"
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
	body, _ := io.ReadAll(asd.Body)
	w.Write(body)
}

func main() {
	service := service2.Service{}
	service.Init()
	defer func(LocalClient *ent.Client) {
		errDb := LocalClient.Close()
		log.Fatal(errDb)
	}(dev.LocalClient)
	mux := http.NewServeMux()
	mux.HandleFunc("/test", service.GetDataHandler)
	mux.HandleFunc("/fillUpTest", FillUpTestData)
	mux.HandleFunc("/DelTest", DelTestData)

	err := http.ListenAndServe(":3333", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}

func DelTestData(writer http.ResponseWriter, request *http.Request) {
	_, err := dev.LocalClient.ProviderRegisterData.Delete().Where(providerregisterdata.Name("Test1")).Exec(dev.Ctx)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}
JSON->Unmarshal -> OBject
func FillUpTestData(writer http.ResponseWriter, request *http.Request) {
	serv, err := dev.LocalClient.ProviderRegisterData.
		Create().
		SetName("Test1").
		SetAddress("www.test.test").
		SetPort("80808080").
		SetDescription("Some Test Service without real data").
		SetLiveInterval(2).
		SetLiveTimeout(4).Save(dev.Ctx)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	end1, err := dev.LocalClient.ProviderEndpoint.
		Create().
		SetName("Endpoint1").
		SetPath("end1").
		SetType("GET").
		SetDescription("ASDASD").
		SetProvider(serv).
		Save(dev.Ctx)

	_, err = dev.LocalClient.EndpointData.
		Create().
		SetDataName("Name").
		SetType("string").
		SetDescription("Same test Name").
		SetEndpointProvided(end1).
		Save(dev.Ctx)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}
