package dto

import model "server/internal/models"

type RegisterRequest struct {
	Account model.Account `json:"account"`
	User    model.User    `json:"user"`
}
