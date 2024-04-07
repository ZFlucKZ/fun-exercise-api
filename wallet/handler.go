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
	Wallet(userId string) ([]Wallet, error)
	WalletByUserId(userId string) ([]Wallet, error)
}

func New(db Storer) *Handler {
	return &Handler{store: db}
}

type Err struct {
	Message string `json:"message"`
}

// WalletHandler
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

// WalletByWalletTypeHandler
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

func (h *Handler) WalletByUserIdHandler(c echo.Context) error {
	userId := c.Param("user_id")

	wallet, err := h.store.WalletByUserId(userId)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wallet)
}
