package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/esuEdu/casa-oliveira/internal/util"
	jwt "github.com/golang-jwt/jwt/v5"
)

type key string

const userIDKey key = "user_id"
const roleKey key = "user_role"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			util.RespondWithJSON(w, http.StatusUnauthorized, util.APIResponse{
				Error: "token not found",
			})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == authHeader {
			util.RespondWithJSON(w, http.StatusUnauthorized, util.APIResponse{
				Error: "token format invalid",
			})
			return
		}

		token, err := util.ValidateToken(tokenStr)
		if err != nil {
			util.RespondWithJSON(w, http.StatusUnauthorized, util.APIResponse{
				Error: "token invalid or expired",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			util.RespondWithJSON(w, http.StatusUnauthorized, util.APIResponse{
				Error: "error processing token",
			})
			return
		}

		userID, ok := claims["user_id"]
		if !ok {
			util.RespondWithJSON(w, http.StatusUnauthorized, util.APIResponse{
				Error: "user_id missing in token",
			})
			return
		}

		userRole, ok := claims["user_role"]
		if !ok {
			util.RespondWithJSON(w, http.StatusUnauthorized, util.APIResponse{
				Error: "user_role missing in token",
			})
			return
		}

		// Add user ID and Role to the context
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		ctx = context.WithValue(r.Context(), roleKey, userRole)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Helper to retrieve user ID from context
func GetUserID(r *http.Request) (interface{}, bool) {
	userID := r.Context().Value(userIDKey)
	return userID, userID != nil
}

func GetUserRole(r *http.Request) (interface{}, bool) {
	userRole := r.Context().Value(roleKey)
	return userRole, userRole != nil
}
