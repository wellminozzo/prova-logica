package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mapas struct {
	Local                  string
	PopulaçãoNoUltimoCenso string
}

func main() {
	dados, err := ReadMAPFile("../mapa.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = WriteSortedCSV("mapa_ordenado.csv", dados)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Novo arquivo CSV criado com os dados ordenados.")
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

		if err != nil {
			return list, err
		}

		emp := Mapas{
			Local:                  line[0],
			PopulaçãoNoUltimoCenso: line[1],
		}

		list = append(list, emp)
	}

	return list, nil
}

func BubbleSort(data []Mapas) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Convertendo as populações para inteiros para comparar
			pop1, _ := strconv.Atoi(strings.TrimSpace(data[j].PopulaçãoNoUltimoCenso))
			pop2, _ := strconv.Atoi(strings.TrimSpace(data[j+1].PopulaçãoNoUltimoCenso))

			if pop1 < pop2 {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func WriteSortedCSV(filePath string, data []Mapas) error {
	// Ordena os dados usando Bubble Sort
	BubbleSort(data)

	// Cria um novo arquivo CSV
	outputFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	csvWriter := csv.NewWriter(outputFile)
	csvWriter.Comma = ';'
	if err := csvWriter.Write([]string{"Local", "População no último censo"}); err != nil {
		return err
	}

	for i, row := range data {
		if i == 0 || i == len(data)-1 {
			continue
		}
		err := csvWriter.Write([]string{row.Local, row.PopulaçãoNoUltimoCenso})
		if err != nil {
			return err
		}
	}

	csvWriter.Flush()

	if err := csvWriter.Error(); err != nil {
		return err
	}

	return nil
}
