package wallet

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
)

type StubWallet struct {
	wallet []Wallet
	err error
}

func (s *StubWallet) Wallets() ([]Wallet, error) {
	return s.wallet, s.err
}

func TestWallet(t *testing.T) {


	t.Run("given unable to get wallets should return 500 and error message", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/wallets")

		stubError := StubWallet{
			err: echo.ErrInternalServerError,
		}

		p := New(&stubError)

		err := p.WalletHandler(c)
		if err != nil {
			t.Errorf("expected nil but got %v", err)
		}

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected %v but got %v", http.StatusInternalServerError, rec.Code)
		}

		if rec.Body.String() != "{\"message\":\"code=500, message=Internal Server Error\"}\n" {
			t.Errorf("expected %v but got %v", "{\"message\":\"code=500, message=Internal Server Error\"}\n", rec.Body.String())
		}
	})

	t.Run("given user able to getting wallet should return list of wallets", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/v1/wallets")

		stubSuccess := StubWallet{
			wallet: []Wallet{
				{
					ID: 1,
					UserID: 1,
					UserName: "John Doe",
					WalletName: "John Savings",
					WalletType: "Savings",
					Balance: 1000,
					CreatedAt: time.Now(),
					},
					{
					ID: 2,
					UserID: 1,
					UserName: "John Doe",
					WalletName: "John Credit Card",
					WalletType: "Credit Card",
					Balance: 500,
					CreatedAt: time.Now(),
					},
					{
					ID: 3,
					UserID: 1,
					UserName: "John Doe",
					WalletName: "John Crypto Wallet",
					WalletType: "Crypto Wallet",
					Balance: 100,
					CreatedAt: time.Now(),
					},
					{
					ID: 4,
					UserID: 2,
					UserName: "Jane Doe",
					WalletName: "Jane Savings",
					WalletType: "Savings",
					Balance: 2000,
					CreatedAt: time.Now(),
					},
					{
					ID: 5,
					UserID: 2,
					UserName: "Jane Doe",
					WalletName: "Jane Credit Card",
					WalletType: "Credit Card",
					Balance: 1000,
					CreatedAt: time.Now(),
					},
					{
					ID: 6,
					UserID: 2,
					UserName: "Jane Doe",
					WalletName: "Jane Crypto Wallet",
					WalletType: "Crypto Wallet",
					Balance: 200,
					CreatedAt: time.Now(),
					},
			},
		}

		p := New(&stubSuccess)

		err := p.WalletHandler(c)
		if err != nil {
			t.Errorf("expected nil but got %v", err)
		}

		if rec.Code != http.StatusOK {
			t.Errorf("expected %v but got %v", http.StatusOK, rec.Code)
		}
	})
}
