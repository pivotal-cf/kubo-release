package jsonReader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	. "osl-collector/oslStructs"
)

func ReadFiles(files []string) [][]byte {
	var jsonData [][]byte
	for _, fileName := range files {
		bytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}
		jsonData = append(jsonData, bytes)
	}

	return jsonData
}

func ParseOSLData(inputs [][]byte) []OSLData {
	var data []OSLData
	for _, input := range inputs {
		var datum OSLData
		err := json.Unmarshal(input, &datum)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, datum)
	}

	return data
}

func MyFunction(folder string, jsonFile string) {
	files := FindFiles(folder, jsonFile)
	contents := ReadFiles(files)
	rawOslData := ParseOSLData(contents)
	flattened := FlattenPackages(rawOslData)
	output, err := json.Marshal(flattened)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(output) + "\n")
}
