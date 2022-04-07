package main

import (
	"C"
	"encoding/json"
	"log"
)
import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

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

//export jsonArray
func jsonArray() *C.char {
	data := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}
	buf := new(bytes.Buffer)
	file, _ := json.MarshalIndent(data, "", " ")
	binary.Write(buf, binary.BigEndian, file)
	str := buf.String()
	fmt.Println(str)
	return C.CString(str)
}

func main() {

}
