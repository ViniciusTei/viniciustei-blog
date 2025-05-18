package handlers

import (
	"log"
	"net/http"

	"github.com/ViniciusTei/viniciustei-blog/internal/usecases"
	"github.com/gorilla/mux"
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

func (uc *UserController) Pages(mux *mux.Router) {
	//mux.HandleFunc("GET /users", uc.HandleApiUsers)
}

func (uc *UserController) Routes(mux *mux.Router) {
	//mux.HandleFunc("POST /api/users", uc.HandleApiUsers)
	mux.HandleFunc("/signin", uc.signIn).Methods("POST")
	mux.HandleFunc("/signout", uc.signOut).Methods("POST")
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
