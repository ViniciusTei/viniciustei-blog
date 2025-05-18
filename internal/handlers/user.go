package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ViniciusTei/viniciustei-blog/internal/middlewares"
	"github.com/ViniciusTei/viniciustei-blog/internal/repositories"
	"github.com/gorilla/mux"
)

type UserController struct {
	authUseCase *repositories.AuthRepositoryImpl
	//TODO: userUserCase *usecases.UserUseCase
}

func NewUserController(authUseCase *repositories.AuthRepositoryImpl) *UserController {
	return &UserController{
		authUseCase: authUseCase,
	}
}

func (uc *UserController) Pages(prefix string, mux *mux.Router) {
	authMiddleware := middlewares.NewAuthMidleware()
	mux.Handle(fmt.Sprintf("%s/profile", prefix), authMiddleware.ServeHTTP(http.HandlerFunc(uc.profile)))
}

func (uc *UserController) Routes(prefix string, mux *mux.Router) {
	mux.HandleFunc(fmt.Sprintf("%s/signin", prefix), uc.signIn).Methods(http.MethodPost)
	mux.HandleFunc(fmt.Sprintf("%s/singout", prefix), uc.signOut).Methods(http.MethodPost)
}

func (uc *UserController) profile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User Profile"))
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
		http.Redirect(w, r, "/login?error=Invalid+credentials", http.StatusSeeOther)
		log.Printf("Error signing in user %s\n", username)
		log.Println("> Error:", err)
		return
	}

	// Set the token in the cookie and redirect to the root
	http.SetCookie(w, &http.Cookie{
		Name:  "auth_token",
		Value: token,
	})
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
