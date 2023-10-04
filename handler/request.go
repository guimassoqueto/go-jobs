package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param %s (type %s) is required", name, typ)
}

// CreateOpening Request Format
type CreateOpeningRequest struct {
	Role string `json:"role"`
	Company string `json:"company"`
	Location string `json:"location"`
	Remote *bool `json:"remote"`
	Link string `json:"link"`
	Salary int64 `json:"salary"`
}

func (cor *CreateOpeningRequest) Validate() error {
	emptyBody := cor.Salary <= 0 && cor.Role == "" && cor.Remote == nil && cor.Location == "" && cor.Link == "" && cor.Company == ""
	if emptyBody {
		return fmt.Errorf("invalid request body")
	}

	if cor.Role == "" {
		return errParamIsRequired("role", "string")
	}

	if cor.Company == "" {
		return errParamIsRequired("company", "string")
	}

	if cor.Location == "" {
		return errParamIsRequired("location", "string")
	}

	if cor.Link == "" {
		return errParamIsRequired("link", "string")
	}

	if cor.Remote == nil {
		return errParamIsRequired("remote", "boolean")
	}

	if cor.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}

	return nil
}