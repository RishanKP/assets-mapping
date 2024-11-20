package mappings

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
	"go.mongodb.org/mongo-driver/mongo"
)

type MappingHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	repo      repository.MappingRepository
	assetRepo repository.AssetsRepository
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {
	var mapping models.Mapping
	var err error
	if err := json.NewDecoder(r.Body).Decode(&mapping); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err = h.assetRepo.GetById(context.TODO(), mapping.AssetId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			api.NewError(w, http.StatusBadRequest, errors.New("invalid asset id"))
		}

		api.NewError(w, http.StatusInternalServerError, errors.New("error fetching asset details"))
		return
	}

	err = h.repo.Create(context.TODO(), mapping)
	if err != nil {
		log.Err(err).Msg("error mapping asset")
		api.NewError(w, http.StatusInternalServerError, errors.New("error mapping asset"))
		return
	}

	api.Result(w, http.StatusCreated, "success", struct{}{})
}

func (h handler) Get(w http.ResponseWriter, r *http.Request) {
	mappings, err := h.repo.Get(context.TODO(), mux.Vars(r)["userId"])
	if err != nil {
		log.Err(err).Msg("failed to get mapping details")
		api.NewError(w, http.StatusBadRequest, errors.New("failed to fetch mapping details"))
		return
	}

	api.Result(w, http.StatusOK, "success", mappings)
	return
}

func (h handler) Delete(w http.ResponseWriter, r *http.Request) {
	err := h.repo.Delete(context.TODO(), mux.Vars(r)["id"])
	if err != nil {
		log.Err(err).Msg("failed to delete")
		api.NewError(w, http.StatusBadRequest, errors.New("failed to delete"))
		return
	}

	api.Result(w, http.StatusOK, "success", struct{}{})
	return
}

func Newhandler(repo repository.MappingRepository, asserRepo repository.AssetsRepository) MappingHandler {
	return handler{
		repo:      repo,
		assetRepo: asserRepo,
	}
}
