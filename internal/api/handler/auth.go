package handler

import (
	"encoding/json"
	"net/http"
)

// RegisterHandler - Xử lý đăng ký người dùng
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Giải mã JSON từ body request
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Kiểm tra thông tin và lưu người dùng vào cơ sở dữ liệu (ở đây chỉ trả về success)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Registration successful")
}
