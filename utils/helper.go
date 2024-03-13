package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)


func ExportJson(i interface{}) {
	file, err := json.MarshalIndent(i, "", " ")
	if err != nil {
		log.Fatal("Error marshalling: ", err)
	}

	_ = ioutil.WriteFile("result.json", file, 0644)
}
