package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/api/services"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/api/status"
	"github.com/victormagalhaess/origin-backend-take-home-assignment/pkg/model"
)

// @Description Risk endpoint
// @Summary Endpoint to calculate the risk of an insurance policy related to a given customer
// @Accept  application/json
// @Produce  application/json
// @Router /api/v1/risk [POST]
// @Param body body model.UserPersonalInformation true "User Personal Information"
// @Success 200 {object} model.RiskProfile
// @Failure 400
// @Failure 500
// @Tags Risk
func Risk(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		status.BadRequest(w, err)
		return
	}

	userInfo := &model.UserPersonalInformation{}
	err = json.Unmarshal(body, userInfo)
	if err != nil {
		status.BadRequest(w, err)
		return
	}

	response, err := services.Risk(*userInfo)
	if err != nil {
		status.ServerError(w, err)
		return
	}

	status.Success(w, response)
}
