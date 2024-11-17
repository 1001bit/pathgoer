package handler

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/1001bit/pathgoer/services/path/pathmodel"
	"github.com/1001bit/pathgoer/services/path/shared/accesstoken"
)

type PathResponse struct {
	Path      pathmodel.Path `json:"path"`
	EditRight bool           `json:"editRight"`
}

func (h *Handler) HandlePath(w http.ResponseWriter, r *http.Request) {
	pathId := r.PathValue("id")
	akserId := "0"
	if claims, ok := accesstoken.GetClaimsFromContext(r.Context()); ok {
		akserId = claims.Id
	}

	path, ownerId, err := h.pathstore.GetPathAndOwnerId(r.Context(), akserId, pathId)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		slog.With("err", err).Error("Failed to get path")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := PathResponse{
		Path:      path,
		EditRight: ownerId == akserId,
	}

	pathJson, err := json.Marshal(resp)
	if err != nil {
		slog.With("err", err).Error("Failed to marshal path")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(pathJson)
}
