package api

import (
	"encoding/json"
	"net/http"
	"pizza-app/internal/data"
)

// handler for create a new order
func (a *App) newPizzaHandler(w http.ResponseWriter, r *http.Request) {
	newPizza := new(data.NewPizza)
	if err := json.NewDecoder(r.Body).Decode(newPizza); err != nil {
		fail(w, http.StatusBadRequest, err.Error())
		return
	}

	err := a.Repo.NewPizzaDetails(newPizza)
	if err != nil {
		fail(w, http.StatusInternalServerError, err.Error())
		return
	}
	send(w, http.StatusOK, nil)
}
