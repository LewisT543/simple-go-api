package handlers

import (
	"encoding/json"
	"github.com/LewisT543/simple-go-api/api"
	"github.com/LewisT543/simple-go-api/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetCoinBalance(writer http.ResponseWriter, request *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, request.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(writer)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*tokenDetails).Coins,
		Code:    http.StatusOK,
	}

	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(writer)
		return
	}
}
