package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IPInfoResponse struct {
	Ip       string `json:"ip,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	City     string `json:"city,omitempty"`
	Region   string `json:"region,omitempty"`
	Country  string `json:"country,omitempty"`
	Loc      string `json:"loc,omitempty"`
	Org      string `json:"org,omitempty"`
	Postal   string `json:"postal,omitempty"`
	Timezone string `json:"timezone,omitempty"`
	Readme   string `json:"readme,omitempty"`
}

func simpleHttpGet(url string) string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(contents)
}

func main() {
	ipInfoURL := "https://ipinfo.io"
	response := simpleHttpGet(ipInfoURL)

	fmt.Println(response)
	fmt.Println("######")

	ipInfoResponse := IPInfoResponse{}
	json.Unmarshal([]byte(response), &ipInfoResponse)
	fmt.Println(ipInfoResponse)
}
