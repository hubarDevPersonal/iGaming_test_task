package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"iGaming/internal/api/dto"
	"iGaming/internal/api/helpers"
	"iGaming/internal/services"
	"log"
	"net/http"
)

type OpenApiGames interface {
	GameProcessor(w http.ResponseWriter, r *http.Request)
}

type openApiGames struct {
	balanceService services.BalanceService
	db             *sql.DB
	json           *helpers.JsonResponse
}

func NewOpenApiGames(db *sql.DB) OpenApiGames {
	return &openApiGames{
		json:           &helpers.JsonResponse{},
		db:             db,
		balanceService: services.NewBalanceService(db),
	}
}

func (o *openApiGames) GameProcessor(w http.ResponseWriter, r *http.Request) {
	log.Println("GameProcessor")
	req := dto.RequestBalance{}
	rawdata, err := helpers.ReadJSON(w, r, &req)
	if err != nil {
		_ = o.json.WriteJSONError(w, fmt.Errorf("invalid request Paylod %v", err), http.StatusBadRequest)
		return
	}

	receivedSig := r.Header.Get("Sign")
	fmt.Println("Received Signature: ", rawdata)
	expectedSig := helpers.GenerateSignature(rawdata)

	if receivedSig != expectedSig {
		_ = o.json.WriteJSONError(w, errors.New("invalid signature"), http.StatusUnauthorized)
		return
	}

	playerReq := dto.ToPlayer(req)
	resp, err := o.balanceService.GetBalance(req.Api, playerReq)
	if err != nil {
		_ = o.json.WriteJSONError(w, fmt.Errorf("failed to get balance %v", err), http.StatusBadRequest)
		return
	}

	res := &helpers.JsonResponse{
		Error:   false,
		Message: "Processed Game Request",
		Data:    resp,
	}

	_ = res.WriteJSON(w, http.StatusOK, nil)
}
