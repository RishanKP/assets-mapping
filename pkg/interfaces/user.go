package interfaces

import "asset-mapping/pkg/models"

type UserLoginResponse struct {
	models.User
	Token string `json:"token"`
}

type Dashboard struct {
	EmployeeList []List `json:"employeeList"`
}

type List struct {
	models.User
	AssetCount int64 `json:"assetCount"`
}
