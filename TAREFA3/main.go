package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Unidade     string `json:"unidade"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
}

type Viacep struct {
	CEP string `json:"cep"`
}

func SearchCEP(cep string) (*Address, error) {

	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var address Address
	err = json.NewDecoder(resp.Body).Decode(&address)
	if err != nil {
		return nil, err
	}

	return &address, nil
}

func ReadCEPCsv(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var ceps []string
	for _, record := range records {
		ceps = append(ceps, record[0])
	}

	return ceps, nil
}

func WriteCEPToCSV(filePath string, addresses []Address) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"CEP", "Logradouro", "Complemento", "Bairro", "Localidade", "UF", "Unidade", "IBGE", "GIA"}
	if err := writer.Write(headers); err != nil {
		return err
	}

	ceps, err := ReadCEPCsv("../CEPs.csv")
	if err != nil {
		return err
	}

	for _, cep := range ceps {
		cep = strings.ReplaceAll(cep, " ", "")

		address, err := SearchCEP(cep)
		if err != nil {
			log.Printf("Erro ao buscar CEP Formato  %s: %v", cep, err)
			continue
		}

		addresses = append(addresses, *address)
	}

	for _, address := range addresses {
		record := []string{address.Cep, address.Logradouro, address.Complemento, address.Bairro, address.Localidade, address.Uf, address.Unidade, address.IBGE, address.GIA}
		err := writer.Write(record)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {

	addresses := []Address{}

	err := WriteCEPToCSV("CEPs_preenchido.csv", addresses)
	if err != nil {
		log.Fatal(err)
	}

}
