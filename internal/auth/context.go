package auth

import "net/http"

func GetUserID(r *http.Request) int64 {
	if id, ok := r.Context().Value("userID").(int64); ok {
		return id
	}
	return 0
}
