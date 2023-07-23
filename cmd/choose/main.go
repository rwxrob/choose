package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/rwxrob/choose"
)

var StringChoices = []string{
	`First choice`,
	`Second choice`,
	`Third choice`,
	`Fourth choice`,
}

var IntChoices = []int{1, 40, 234, 32}

type Entry struct {
	Name  string  `json:"name"`
	Value float32 `json:"value,omitempty"`
}

// MarshalJSON implements json.Marshaler as JSON.
func (e Entry) JSON() ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(e)
	if err != nil {
		return nil, err
	}
	byt := buf.Bytes()
	return byt[:len(byt)-1], err
}

// String implements fmt.Stringer as JSON.
func (e Entry) String() string {
	byt, err := e.JSON()
	if err != nil {
		return "null"
	}
	return string(byt)
}

var EntryChoices = []Entry{
	{Name: "Rob", Value: 23.34},
	{Name: "Doris", Value: 4.23},
	{Name: "<M@xx>", Value: 3234},
}

func main() {
	_, pick, err := choose.From(StringChoices)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	//fmt.Println(pick.Name)
	fmt.Println(pick)
}
