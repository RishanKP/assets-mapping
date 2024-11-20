package dashboard

import (
	"context"
	"errors"
	"net/http"

	"asset-mapping/library/api"
	"asset-mapping/pkg/interfaces"
	"asset-mapping/pkg/repository"
)

type DashboardHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	repo        repository.UserRepository
	mappingRepo repository.MappingRepository
}

func (h handler) Get(w http.ResponseWriter, r *http.Request) {
	var response interfaces.Dashboard
	users, err := h.repo.Get(context.TODO())
	if err != nil {
		api.NewError(w, http.StatusInternalServerError, errors.New("error fetching user data"))
	}

	for _, u := range users {
		var list interfaces.List
		list.User = u
		list.AssetCount = h.mappingRepo.GetCountByUserId(context.TODO(), u.ID.Hex())

		response.EmployeeList = append(response.EmployeeList, list)
	}

	api.Result(w, http.StatusOK, "success", response)
	return
}

func Newhandler(repo repository.UserRepository, mappingRepo repository.MappingRepository) DashboardHandler {
	return handler{
		repo:        repo,
		mappingRepo: mappingRepo,
	}
}
