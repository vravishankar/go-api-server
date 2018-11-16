package handler

import (
	"net/http"

	"../helper"
	repo "../repository"
)

// GetAll -
func GetAll(w http.ResponseWriter, r *http.Request) {
	payments, err := repo.GetAll()

	if err != nil {
		helper.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, payments)
}
