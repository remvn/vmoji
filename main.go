package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := "./assets/vmoji/"

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	m := make(map[string]string)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fType := filepath.Ext(file.Name())
		key := strings.TrimSuffix(file.Name(), fType)
		switch fType {
		case
			".webp",
			".svg",
			".gif",
			".ico",
			// types above can't be processed by filter
			".jpg", ".jfif", ".jpeg",
			".png", ".apng":
			m[key] = file.Name()
		default:
			continue
		}
	}

	jsonFile := "./data/vmoji.json"
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(jsonFile, data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s generated with %d items\n", jsonFile, len(m))
}
