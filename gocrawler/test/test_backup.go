package test

import (
	"encoding/gob"
	"fmt"
	"os"
)

func TestBackup() {
	testBackup_Backup()
	testBackup_Load()
}

func testBackup_Load() {
	//doneMap := make(map[string]bool)
	var doneMap map[string]bool
	//var b string
	file, err := os.Open("data/tmp.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	dec := gob.NewDecoder(file)
	dec.Decode(&doneMap)
	//json.Unmarshal([]byte(b), &doneMap)
	fmt.Println("load bak", doneMap)
}

func testBackup_Backup() {
	doneMap := make(map[string]bool)
	doneMap["key14"] = false
	doneMap["key2"] = true
	doneMap["key2中"] = true
	//b, _ := json.Marshal(doneMap)
	file, err := os.Create("data/tmp.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	enc := gob.NewEncoder(file)
	//enc.Encode(string(b))
	enc.Encode(doneMap)
	fmt.Println("backup ok")
}
