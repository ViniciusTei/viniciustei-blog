package middlewares

import (
	"net/http"

	"github.com/ViniciusTei/viniciustei-blog/utils"
)

type AuthMidleware struct {
}

func (l *AuthMidleware) ServeHTTP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, userName, err := utils.GetUserFromCookie(r)
		if err != nil {
			http.Redirect(w, r, "/login?error=Please+login+to+access+this+page", http.StatusSeeOther)
		}

		if userId == "" && userName == "" {
			http.Redirect(w, r, "/login?error=Please+login+to+access+this+page", http.StatusSeeOther)
		}

		next.ServeHTTP(w, r)
	})
}

func NewAuthMidleware() *AuthMidleware {
	return &AuthMidleware{}
}
