package wallet

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	store Storer
}

type Storer interface {
	Wallets() ([]Wallet, error)
	CreateWallet(wallet *Wallet) (id int, err error)
	UpdateWallet(wallet *Wallet) (err error)
	Wallet(userId string) ([]Wallet, error)
	WalletByUserId(userId string) ([]Wallet, error)
	DeleteWalletByUserId(userId string) (err error)
}

func New(db Storer) *Handler {
	return &Handler{store: db}
}

type Err struct {
	Message string `json:"message"`
}

//  WalletHandler
//	@Summary		Get all wallets
//	@Description	Get all wallets
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/wallets [get]
//	@Failure		500	{object}	Err
func (h *Handler) WalletHandler(c echo.Context) error {
	wallets, err := h.store.Wallets()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wallets)
}

//  CreateWalletHandler
//	@Summary		Create a wallet
//	@Description	Create a wallet
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/wallets [post]
//	@Failure		500	{object}	Err
func (h *Handler) CreateWalletHandler(c echo.Context) error {
	wallet := new(Wallet)
	if err := c.Bind(wallet); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	id, err := h.store.CreateWallet(wallet)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	wallet.ID = id

	return c.JSON(http.StatusOK, wallet)
}

//  UpdateWalletHandler
//	@Summary		Update a wallet
//	@Description	Update a wallet
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/wallets [patch]
//	@Failure		500	{object}	Err
func (h *Handler) UpdateWalletHandler(c echo.Context) error {
	wallet := new(Wallet)
	if err := c.Bind(wallet); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	err := h.store.UpdateWallet(wallet)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, "Wallet updated")
}

//  WalletByWalletTypeHandler
//	@Summary		Get a wallet by Wallet type
//	@Description	Get a wallet by Wallet type
//  @Param			wallet-type	path	string	true	"wallet-type"
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/wallets/:wallet-type [get]
//	@Failure		500	{object}	Err
func (h *Handler) WalletByWalletTypeHandler(c echo.Context) error {
	walletType := c.Param("wallet_type")

	wallet, err := h.store.Wallet(walletType)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wallet)
}

//  WalletByUserIdHandler
//	@Summary		Get a wallet by User Id
//	@Description	Get a wallet by User Id
//  @Param			user_id	path	string	true	"user_id"
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/users/:user_id/wallets [get]
//	@Failure		500	{object}	Err
func (h *Handler) WalletByUserIdHandler(c echo.Context) error {
	userId := c.Param("user_id")

	wallet, err := h.store.WalletByUserId(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wallet)
}

//  DeleteWalletByUserIdHandler
//	@Summary		Delete a wallet by User Id
//	@Description	Delete a wallet by User Id
//  @Param			user_id	path	string	true	"user_id"
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/users/:user_id/wallets [delete]
//	@Failure		500	{object}	Err
func (h *Handler) DeleteWalletByUserIdHandler(c echo.Context) error {
	userId := c.Param("user_id")

	err := h.store.DeleteWalletByUserId(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, "Wallet deleted")
}
