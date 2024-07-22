package model

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Brasilapi struct {
	Canal        string `json:"canal"`
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type Viacep struct {
	Canal       string `json:"canal"`
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (c *Brasilapi) NewTaxaBrasilapi(ctx context.Context, cep string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprint("https://brasilapi.com.br/api/cep/v1/", cep), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	//time.Sleep(time.Second * 2)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		return err
	}

	c.Canal = "BrasilApi"

	return nil
}

func (c *Viacep) NewTaxaViacep(ctx context.Context, cep string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&c); err != nil {
		return err
	}

	c.Canal = "ViaCEP"

	return nil
}
