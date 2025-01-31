package handler

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1001bit/pathgoer/services/path/pathmodel"
	"github.com/1001bit/pathgoer/services/path/shared/accesstoken"
	"github.com/lib/pq"
)

func (h *Handler) HandleUpdatePath(w http.ResponseWriter, r *http.Request) {
	askerId := "0"
	if claims, ok := accesstoken.GetClaimsFromContext(r.Context()); ok {
		askerId = claims.Id
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	req := pathmodel.Path{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !nameValid(req.Name) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req.Id = r.PathValue("id")
	err = h.pathstore.UpdatePath(r.Context(), req, askerId)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err, ok := err.(*pq.Error); ok && err.Code == "23505" {
		w.WriteHeader(http.StatusConflict)
		return
	} else if err != nil {
		slog.With("err", err).Error("Failed to create path")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
