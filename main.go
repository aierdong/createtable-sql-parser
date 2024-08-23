package main

import (
	"encoding/json"
	"github.com/aierdong/createtable-sql-parser/visitor"
	"log"
)

func main() {
	table, err := visitor.ParseMySql(exampleMySQL)
	if err != nil {
		log.Fatal(err)
	}

	bs, _ := json.MarshalIndent(table, "", " ")
	log.Println(string(bs))
}
