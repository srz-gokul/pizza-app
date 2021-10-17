package api

import (
	"log"
	"net/http"
)

func (a *App) statusHandler(w http.ResponseWriter, r *http.Request) {
	pgStatus, err := a.Repo.DBStatusInfo()
	if err != nil {
		log.Println("Database operation failed", err)
		fail(w, http.StatusInternalServerError, "Database operation failed.")
		return
	}
	send(w, http.StatusOK, pgStatus)
}
