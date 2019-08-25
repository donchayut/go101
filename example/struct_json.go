package main

import (
	"encoding/json"
	"log"
)

// Generated by https://quicktype.io

func main() {
	jsonBlob := []byte(`{
		"width": 48,
		"length": 64,
		"border": "solid"
}`)

	type rectangle struct {
		Width   int    `json:"width"`
		Length  int    `json:"length"`
		Borders string `json:"border"`
	}

	var rec rectangle
	err := json.Unmarshal(jsonBlob, &rec)
	if err != nil {
		log.Fatal(err)
	}
}

// END OMIT
