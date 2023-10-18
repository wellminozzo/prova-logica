package main

import (
	"testing"
)

func TestSearchCEP(t *testing.T) {
	cep := "22290-240"
	expectedLogradouro := "Avenida Brigadeiro Faria Lima" //04538-133
	address, err := SearchCEP(cep)
	if err != nil {
		t.Errorf("erro ao carregar CEP %s: %v", cep, err)
	}
	if address.Logradouro != expectedLogradouro {
		t.Errorf("logradouro esperado %s, obteve %s", expectedLogradouro, address.Logradouro)
	}
}
