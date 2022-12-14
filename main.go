package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AutoGenerated struct {
	Data []Data `json:"data"`
}
type Data struct {
	IDYear     int    `json:"ID Year"`
	Year       string `json:"Year"`
	Population int    `json:"Population"`
}

func main() {
	res, err := http.Get("https://datausa.io/api/data?drilldowns=Nation&measures=Population")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body) // response body is []byte
	autoGenerated := AutoGenerated{}
	_ = json.Unmarshal(body, &autoGenerated)
	for _, data := range autoGenerated.Data {
		if data.IDYear%2 != 0 {
			fmt.Println(data.Year, data.Population)
		}
	}
}
