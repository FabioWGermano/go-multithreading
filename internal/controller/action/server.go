package controller

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/FabioWGermano/go-multithreading/internal/controller/response"
	"github.com/FabioWGermano/go-multithreading/internal/model"
)

type Endereco struct {
	Canal  string `json:"canal"`
	Cep    string `json:"cep"`
	Rua    string `json:"endereco"`
	Bairro string `json:"bairro"`
	Cidade string `json:"cidade"`
	Estado string `json:"estado"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var cep = r.URL.Query().Get("cep")
	var msgBrasilApi = make(chan model.Brasilapi)
	var msgViaCep = make(chan model.Viacep)

	go BuscarViaCep(r.Context(), cep, msgViaCep)
	go BuscarBrasilApi(r.Context(), cep, msgBrasilApi)

	for {
		select {
		case msg := <-msgBrasilApi:
			msgTratada := Endereco{
				Canal:  msg.Canal,
				Cep:    msg.Cep,
				Rua:    msg.Street,
				Bairro: msg.Neighborhood,
				Cidade: msg.City,
				Estado: msg.State,
			}
			log.Print(msgTratada)
			response.NewSucess(msgTratada, http.StatusAccepted).Send(w)
			return
		case msg := <-msgViaCep:
			msgTratada := Endereco{
				Canal:  msg.Canal,
				Cep:    msg.Cep,
				Rua:    msg.Logradouro,
				Bairro: msg.Bairro,
				Cidade: msg.Localidade,
				Estado: msg.Uf,
			}
			log.Print(msgTratada)
			response.NewSucess(msgTratada, http.StatusAccepted).Send(w)
			return
		case <-time.After(time.Second * 1):
			log.Printf("Timeout")
			response.NewError(errors.New("Timeout"), http.StatusRequestTimeout).Send(w)
			return
		}
	}
}

func BuscarViaCep(ctx context.Context, cep string, msgViaCep chan model.Viacep) error {
	var c model.Viacep
	if err := c.NewTaxaViacep(ctx, cep); err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	msgViaCep <- c
	return nil
}

func BuscarBrasilApi(ctx context.Context, cep string, msgViaCep chan model.Brasilapi) error {
	var c model.Brasilapi
	if err := c.NewTaxaBrasilapi(ctx, cep); err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	msgViaCep <- c
	return nil
}
