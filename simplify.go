package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
	jsonFile, err := os.Open("fosmmpl.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var dat map[string]interface{}
	if err := json.Unmarshal(byteValue, &dat); err != nil {
        panic(err)
    }
	fmt.Println(dat)

}
