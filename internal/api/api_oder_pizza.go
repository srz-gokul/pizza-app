package api

import (
	"encoding/json"
	"net/http"
	"pizza-app/internal/data"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

// handler for create a new order
func (a *App) orderPizzaHandler(w http.ResponseWriter, r *http.Request) {
	orderData := new(data.OrderData)
	if err := json.NewDecoder(r.Body).Decode(orderData); err != nil {
		fail(w, http.StatusBadRequest, err.Error())
		return
	}

	startTime := time.Now()
	orderData.CookingStage = data.CookingStage1
	orderData.StartTime = &startTime
	orderData.IsActive = true

	err := a.Repo.AddOrderDetails(orderData)
	if err != nil {
		fail(w, http.StatusInternalServerError, err.Error())
		return
	}
	send(w, http.StatusOK, orderData)
}

// handler for update the order details
func (a *App) updateOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderData := new(data.OrderData)
	if err := json.NewDecoder(r.Body).Decode(orderData); err != nil {
		fail(w, http.StatusBadRequest, err.Error())
		return
	}
	if orderData.CookingStage == data.CookingStage5 {
		endTime := time.Now()
		orderData.EndTime = &endTime
		orderData.IsActive = false
		a.Msg.SendMessage("Pizza ready to deliver!!")
	} else {
		orderData.IsActive = true
	}
	err := a.Repo.UpdateOrderDetails(orderData)
	if err != nil {
		fail(w, http.StatusInternalServerError, err.Error())
		return
	}
	send(w, http.StatusOK, nil)
}

// handler for get the order status
func (a *App) orderStatusHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		fail(w, http.StatusBadRequest, err.Error())
		return
	}
	orderDetails, err := a.Repo.GetOrderDetails(userID)
	if err != nil {
		fail(w, http.StatusInternalServerError, err.Error())
		return
	}
	send(w, http.StatusOK, orderDetails)
}
