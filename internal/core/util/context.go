package util

import(
	"net/http"
)

func GetUserID(w http.ResponseWriter, r *http.Request) (string, bool) {
	userID, ok := r.Context().Value("userID").(string)
	if !ok || userID == "" {
		SendErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return "", false
	}
	return userID, true
}
