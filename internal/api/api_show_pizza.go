package api

import (
	"encoding/json"
	"net/http"
	"pizza-app/internal/data"
)

// handler for show all pizza type

func (a *App) showPizzaDetailsHandler(w http.ResponseWriter, r *http.Request) {
	newPizza := new(data.NewPizza)
	if err := json.NewDecoder(r.Body).Decode(newPizza); err != nil {
		fail(w, http.StatusBadRequest, err.Error())
		return
	}

	err := a.Repo.ShowPizzaDetails(newPizza)
	if err != nil {
		fail(w, http.StatusInternalServerError, err.Error())
		return
	}
	send(w, http.StatusOK, nil)
}

// handler for show all pizza orders from a date
func (a *App) ListAllOrdersHandler(w http.ResponseWriter, r *http.Request) {
	pizza := new(data.NewPizza)
	if err := json.NewDecoder(r.Body).Decode(pizza); err != nil {
		fail(w, http.StatusBadRequest, err.Error())
		return
	}

	err := a.Repo.ListAllOrdersDetails(pizza)
	if err != nil {
		fail(w, http.StatusInternalServerError, err.Error())
		return
	}
	send(w, http.StatusOK, nil)
}
