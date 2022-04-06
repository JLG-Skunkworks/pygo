package main

import (
	"C"
	"encoding/json"
	"log"
)

//export helloWorld
func helloWorld() {
	log.Println("Hello World")
}

//export jsonDict
func jsonDict() {
	fileCount := map[string]int{
		"cpp": 10,
		"js":  8,
		"go":  10,
	}
	bytes, _ := json.Marshal(fileCount)
	log.Println(string(bytes))
}

func main() {

}
