package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// handler for delete a  order

func (a *App) deleteStatusHandler(w http.ResponseWriter, r *http.Request) {
	orderID, err := strconv.Atoi(chi.URLParam(r, "order_id"))
	if err != nil {
		fail(w, http.StatusBadRequest, err.Error())
		return
	}
	err = a.Repo.DeleteOrderDetails(orderID)
	if err != nil {
		fail(w, http.StatusInternalServerError, err.Error())
		return
	}
	send(w, http.StatusOK, nil)
}
