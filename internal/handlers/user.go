package handlers

import (
	"log"
	"net/http"

	"github.com/ViniciusTei/viniciustei-blog/internal/usecases"
)

type UserController struct {
	authUseCase *usecases.AuthUseCase
	//TODO: userUserCase *usecases.UserUseCase
}

func NewUserController(authUseCase *usecases.AuthUseCase) *UserController {
	return &UserController{
		authUseCase: authUseCase,
	}
}

func (uc *UserController) Pages(mux *http.ServeMux) {
	//mux.HandleFunc("GET /users", uc.HandleApiUsers)
}

func (uc *UserController) Routes(mux *http.ServeMux) {
	//mux.HandleFunc("POST /api/users", uc.HandleApiUsers)
	mux.HandleFunc("POST /signin", uc.signIn)
	mux.HandleFunc("POST /signout", uc.signOut)
}

func (uc *UserController) signOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "auth_token",
		Value: "",
	})
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

func (h *UserController) signIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	token, err := h.authUseCase.SignIn(username, password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		log.Println("Error signing in:", err)
		return
	}

	// Set the token in the cookie and redirect to the root
	http.SetCookie(w, &http.Cookie{
		Name:  "auth_token",
		Value: token,
	})
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
