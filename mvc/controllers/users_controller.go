package controllers

import (
	"encoding/json"
	"github.com/AndrewYang17/golang-microservices/mvc/services"
	"github.com/AndrewYang17/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		appErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(appErr)
		res.WriteHeader(appErr.StatusCode)
		_, _ = res.Write(jsonValue)
		return
	}

	user, appErr := services.GetUser(userId)
	if appErr != nil {
		jsonValue, _ := json.Marshal(appErr)
		res.WriteHeader(appErr.StatusCode)
		_, _ = res.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	_, _ = res.Write(jsonValue)

}