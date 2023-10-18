package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
)

type Mapas struct {
	Local                  string
	PopulaçãoNoUltimoCenso string
	IDH                    float64
}

func ReadMAPFile(filePath string) ([]Mapas, error) {
	var list []Mapas

	csvFile, err := os.Open(filePath)
	if err != nil {
		return list, err
	}

	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	csvReader.Comma = ';'
	csvReader.LazyQuotes = true

	csvLines, err := csvReader.ReadAll()
	if err != nil {
		return list, err
	}
	for _, line := range csvLines {

		idh := rand.Float64()*(0.9-0.5) + 0.5
		emp := Mapas{
			Local:                  line[0],
			PopulaçãoNoUltimoCenso: line[1],
			IDH:                    idh,
		}

		list = append(list, emp)
	}

	return list, nil
}

func WriteMAPFile(filePath string, data []Mapas) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	firstLine := true
	for _, value := range data {

		record := []string{value.Local, value.PopulaçãoNoUltimoCenso}
		if firstLine {
			record = append(record, "IDH")
			firstLine = false

		} else {
			record = append(record, fmt.Sprintf("%.2f", value.IDH))
		}
		err := writer.Write(record)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	dados, err := ReadMAPFile("../mapa.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = WriteMAPFile("mapa_com_idh.csv", dados)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Arquivo criado com sucesso!")
}
