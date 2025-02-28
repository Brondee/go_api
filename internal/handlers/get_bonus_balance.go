package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/brondee/go_api/api"
	"github.com/brondee/go_api/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetBonusBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.BonusBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.BonusDetails
	tokenDetails = (*database).GetUserBonuses(params.Username)
	if tokenDetails != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.BonusBalanceResponse{
		Balance: (*tokenDetails).Bonuses,
		Code: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
