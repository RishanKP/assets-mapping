package assets

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"asset-mapping/library/api"
	"asset-mapping/pkg/models"
	"asset-mapping/pkg/repository"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type AssetsHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	repo repository.AssetsRepository
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {
	var asset models.Assets
	var err error
	if err := json.NewDecoder(r.Body).Decode(&asset); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = h.repo.Create(context.TODO(), asset)
	if err != nil {
		log.Err(err).Msg("error creating asset")
		api.NewError(w, http.StatusInternalServerError, errors.New("error creating asset"))
		return
	}

	api.Result(w, http.StatusCreated, "success", struct{}{})
}

func (h handler) Update(w http.ResponseWriter, r *http.Request) {
	var asset models.Assets
	var err error
	if err := json.NewDecoder(r.Body).Decode(&asset); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = h.repo.Update(context.TODO(), asset)
	if err != nil {
		log.Err(err).Msg("error updating user")
		api.NewError(w, http.StatusInternalServerError, errors.New("error updating user"))
		return
	}

	api.Result(w, http.StatusOK, "success", struct{}{})
}

func (h handler) GetById(w http.ResponseWriter, r *http.Request) {
	asset, err := h.repo.GetById(context.TODO(), mux.Vars(r)["id"])
	if err != nil {
		log.Err(err).Msg("failed to get user details")
		api.NewError(w, http.StatusBadRequest, errors.New("failed to fetch details"))
		return
	}

	api.Result(w, http.StatusOK, "success", asset)
	return
}

func (h handler) Get(w http.ResponseWriter, r *http.Request) {
	assets, err := h.repo.Get(context.TODO())
	if err != nil {
		log.Err(err).Msg("failed to get user details")
		api.NewError(w, http.StatusBadRequest, errors.New("failed to fetch details"))
		return
	}

	api.Result(w, http.StatusOK, "success", assets)
	return
}

func (h handler) Delete(w http.ResponseWriter, r *http.Request) {
	err := h.repo.Delete(context.TODO(), mux.Vars(r)["id"])
	if err != nil {
		log.Err(err).Msg("failed to delete user")
		api.NewError(w, http.StatusBadRequest, errors.New("failed to delete user"))
		return
	}

	api.Result(w, http.StatusOK, "success", struct{}{})
	return
}

func Newhandler(repo repository.AssetsRepository) AssetsHandler {
	return handler{
		repo: repo,
	}
}
