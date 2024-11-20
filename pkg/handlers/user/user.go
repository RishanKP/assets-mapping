package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"asset-mapping/library/api"
	"asset-mapping/library/jwt"
	"asset-mapping/library/utils"
	"asset-mapping/pkg/interfaces"
	"asset-mapping/pkg/models"
	"asset-mapping/pkg/repository"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	repo repository.UserRepository
}

func (h handler) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var err error
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	existingUser, _ := h.repo.GetByEmail(context.TODO(), user.Email)
	if existingUser.Email != "" {
		api.NewError(w, http.StatusBadRequest, errors.New("email id exists"))
		return
	}

	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		log.Err(err).Msg("error hashing password")
		api.NewError(w, http.StatusInternalServerError, errors.New("unexpected server error"))
		return
	}

	err = h.repo.Create(context.TODO(), user)
	if err != nil {
		log.Err(err).Msg("error creating user")
		api.NewError(w, http.StatusInternalServerError, errors.New("error creating user"))
		return
	}

	api.Result(w, http.StatusCreated, "success", struct{}{})
}

func (h handler) Update(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var err error
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = h.repo.Update(context.TODO(), user)
	if err != nil {
		log.Err(err).Msg("error updating user")
		api.NewError(w, http.StatusInternalServerError, errors.New("error updating user"))
		return
	}

	api.Result(w, http.StatusCreated, "success", struct{}{})
}

func (h handler) Login(w http.ResponseWriter, r *http.Request) {

	var req interfaces.LoginCredentials

	var err error
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := h.repo.GetByEmail(context.TODO(), req.Email)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			api.NewError(w, http.StatusBadRequest, errors.New("account not found"))
			return
		}
		api.NewError(w, http.StatusInternalServerError, errors.New("failed to fetch details"))
		return
	}

	if !utils.ComparePassword(req.Password, user.Password) {
		api.NewError(w, http.StatusUnauthorized, errors.New("invalid password"))
		return
	}

	token, err := jwt.CreateToken(jwt.Claims{
		UserId:   user.ID.Hex(),
		Username: user.Email,
	})

	if err != nil {
		api.NewError(w, http.StatusInternalServerError, errors.New("failed to generate token"))
		return
	}

	res := interfaces.UserLoginResponse{
		User:  user,
		Token: token,
	}

	api.Result(w, http.StatusOK, "success", res)
}

func (h handler) GetById(w http.ResponseWriter, r *http.Request) {
	user, err := h.repo.GetById(context.TODO(), mux.Vars(r)["id"])
	if err != nil {
		log.Err(err).Msg("failed to get user details")
		api.NewError(w, http.StatusBadRequest, errors.New("failed to fetch details"))
		return
	}

	api.Result(w, http.StatusOK, "success", user)
	return
}

func (h handler) Get(w http.ResponseWriter, r *http.Request) {
	users, err := h.repo.Get(context.TODO())
	if err != nil {
		log.Err(err).Msg("failed to get user details")
		api.NewError(w, http.StatusBadRequest, errors.New("failed to fetch details"))
		return
	}

	api.Result(w, http.StatusOK, "success", users)
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

func Newhandler(repo repository.UserRepository) UserHandler {
	return handler{
		repo: repo,
	}
}
