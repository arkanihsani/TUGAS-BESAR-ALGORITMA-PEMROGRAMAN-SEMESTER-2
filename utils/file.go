package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"pemilu/data"
)

const calegFile = "caleg.json"

func SaveCaleg(caleg data.ListCaleg, n int) {
	file, err := os.Create(calegFile)
	if err != nil {
		fmt.Println("Gagal menyimpan data:", err)
		return
	}
	defer file.Close()

	// potong array sesuai jumlah caleg
	toSave := caleg[:n]

	encoder := json.NewEncoder(file)
	err = encoder.Encode(toSave)
	if err != nil {
		fmt.Println("Gagal menyimpan data:", err)
	}
}

func LoadCaleg() (data.ListCaleg, int) {
	var caleg data.ListCaleg
	var slice []data.Caleg

	file, err := os.Open(calegFile)
	if err != nil {
		return caleg, 0 // file belum ada
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&slice)
	if err != nil {
		return caleg, 0
	}

	count := len(slice)
	for i := 0; i < count; i++ {
		caleg[i] = slice[i]
	}
	return caleg, count
}
