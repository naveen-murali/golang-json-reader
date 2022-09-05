package main

import (
	"fmt"

	readjson "read-json/read-json"
)

type Folder struct {
	Path string `json:"path"`
	Name string `json:"name"`
}

type DataFile struct {
	Folders []Folder `json:"folders"`
}

func main() {
	var dataStruct DataFile
	readjson.ReadJson("./assets/data.json", &dataStruct)

	fmt.Printf("%+v\n", dataStruct)
	readjson.PrintJsonInFormat(dataStruct)
}
