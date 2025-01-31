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
	Path      pathmodel.FullPath `json:"path"`
	EditRight bool               `json:"editRight"`
}

// single path
func (h *Handler) HandlePath(w http.ResponseWriter, r *http.Request) {
	akserId := "0"
	if claims, ok := accesstoken.GetClaimsFromContext(r.Context()); ok {
		akserId = claims.Id
	}

	if !idValid(r.PathValue("id")) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	path, err := h.pathstore.GetFullPath(r.Context(), akserId, r.PathValue("id"))
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
		EditRight: path.OwnerId == akserId,
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

// Multiple paths
func (h *Handler) HandleUserPaths(w http.ResponseWriter, r *http.Request) {
	askerId := "0"
	if claims, ok := accesstoken.GetClaimsFromContext(r.Context()); ok {
		askerId = claims.Id
	}

	userId := r.PathValue("id")
	if !idValid(userId) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	paths, err := h.pathstore.GetPaths(r.Context(), userId, askerId)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		slog.With("err", err).Error("Failed to get paths")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pathsJson, err := json.Marshal(paths)
	if err != nil {
		slog.With("err", err).Error("Failed to marshal paths")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(pathsJson)
}
