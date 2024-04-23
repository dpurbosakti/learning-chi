package account

import (
	"context"
	"net/http"

	ma "github.com/dpurbosakti/go-native/internal/models/account"
	aSqlc "github.com/dpurbosakti/go-native/internal/repositories/account/sqlc"
	as "github.com/dpurbosakti/go-native/internal/services/account"
	jh "github.com/dpurbosakti/go-native/pkg/jsonhelper"
)

type AccountHandler struct {
	AccountService *as.AccountService
}

func NewAccountHandler(accountService *as.AccountService) *AccountHandler {
	return &AccountHandler{
		AccountService: accountService,
	}
}

func (h *AccountHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req ma.CreateAccountRequest
	err := jh.ParseJSON(r, &req)
	if err != nil {
		jh.WriteError(w, http.StatusBadRequest, err)
		return
	}

	res, err := h.AccountService.Store.CreateAccount(context.Background(), aSqlc.CreateAccountParams{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		jh.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	response := ma.AccountResponse{
		ID:                res.ID,
		Name:              res.Name,
		Email:             res.Email,
		Phone:             res.Phone,
		PasswordChangedAt: res.PasswordChangedAt.Time,
		CreatedAt:         res.CreatedAt,
		UpdatedAt:         res.UpdatedAt,
		IsEmailVerified:   res.IsEmailVerified,
	}

	jh.WriteJSON(w, http.StatusOK, response)
}

func (h *AccountHandler) GetList(w http.ResponseWriter, r *http.Request) {
	jh.WriteJSON(w, http.StatusOK, "ok")
}

func (h *AccountHandler) GetDetail(w http.ResponseWriter, r *http.Request) {

}

func (h *AccountHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *AccountHandler) Delete(w http.ResponseWriter, r *http.Request) {

}
