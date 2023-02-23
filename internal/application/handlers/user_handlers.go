package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/bvaledev/go-expert-commerce-api/internal/domain/contracts"
	"github.com/bvaledev/go-expert-commerce-api/internal/domain/dto"
	"github.com/bvaledev/go-expert-commerce-api/internal/domain/entity"
	pkg "github.com/bvaledev/go-expert-commerce-api/pkg/entity"
	"github.com/go-chi/jwtauth"
)

type UserHandlers struct {
	UserDB contracts.IUserRepository
}

func NewUserHandler(db contracts.IUserRepository) *UserHandlers {
	return &UserHandlers{
		UserDB: db,
	}
}

// Create user godoc
//
//	@Summary		Create user
//	@Description	Create user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body	dto.CreateUserDTO	true	"User registration info"
//	@Success		201
//	@Failure		500	{object}	error
//	@Router			/users [post]
func (h *UserHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.CreateUserDTO
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(userDTO.Name, userDTO.Email, userDTO.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Authenticate user godoc
//
//	@Summary		Authenticate user
//	@Description	Authenticate user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.AuthenticateUserDTO	true	"User credentials"
//	@Success		200		{object}	dto.AuthenticateUserResponse
//	@Failure		404
//	@Failure		401
//	@Failure		500	{object}	error
//	@Router			/users/login [post]
func (h *UserHandlers) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var userAuth dto.AuthenticateUserDTO
	err := json.NewDecoder(r.Body).Decode(&userAuth)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.UserDB.FindByEmail(userAuth.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !user.ValidatePassword(userAuth.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)

	expiresTime := time.Second * time.Duration(jwtExpiresIn)
	mapClaims := map[string]interface{}{
		"sub":        user.ID.String(),
		"user_name":  user.Name,
		"user_email": user.Email,
		"exp":        time.Now().Add(expiresTime).Unix(),
	}
	_, tokenString, err := jwt.Encode(mapClaims)
	if !user.ValidatePassword(userAuth.Password) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var accessToken = dto.AuthenticateUserResponse{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

func (h *UserHandlers) CurrentUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, _ := jwtauth.FromContext(r.Context())
		currentUser := struct {
			Id    string
			Name  string
			Email string
		}{
			Id: claims["sub"].(string),
		}

		validId, err := pkg.ParseID(currentUser.Id)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		user, err := h.UserDB.FindById(validId.String())
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		currentUser.Email = user.Email
		currentUser.Name = user.Name
		ctx := context.WithValue(r.Context(), "user", currentUser)
		log.Println("from middleware", currentUser)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
