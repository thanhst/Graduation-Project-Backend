package dto

type UpdateUserRequest struct {
	Fullname     string `json:"fullname"`
	OldPassword  string `json:"oldPassword"`
	NewPassword  string `json:"newPassword"`
	ImagePreview string `json:"imagePreview"`
}
