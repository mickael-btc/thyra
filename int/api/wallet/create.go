package wallet

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/go-openapi/runtime/middleware"
	"github.com/massalabs/thyra/api/swagger/server/models"
	"github.com/massalabs/thyra/api/swagger/server/restapi/operations"
	"github.com/massalabs/thyra/pkg/node/base58"
	"github.com/massalabs/thyra/pkg/wallet"
)

func NewCreate(walletStorage *sync.Map) operations.MgmtWalletCreateHandler {
	return &walletCreate{walletStorage: walletStorage}
}

type walletCreate struct {
	walletStorage *sync.Map
}

func (c *walletCreate) Handle(params operations.MgmtWalletCreateParams) middleware.Responder {
	if params.Body.Nickname == nil || len(*params.Body.Nickname) == 0 {
		return operations.NewMgmtWalletCreateBadRequest().WithPayload(
			&models.Error{
				Code:    errorCodeWalletCreateNoNickname,
				Message: "Error: nickname field is mandatory.",
			})
	}

	_, ok := c.walletStorage.Load(*params.Body.Nickname)
	if ok {
		return operations.NewMgmtWalletCreateInternalServerError().WithPayload(
			&models.Error{
				Code:    errorCodeWalletAlreadyExists,
				Message: "Error: a wallet with the same nickname already exists.",
			})
	}

	if params.Body.Password == nil || len(*params.Body.Password) == 0 {
		return operations.NewMgmtWalletCreateBadRequest().WithPayload(
			&models.Error{
				Code:    errorCodeWalletCreateNoPassword,
				Message: "Error: password field is mandatory.",
			})
	}

	newWallet, err := wallet.New(*params.Body.Nickname)
	if err != nil {
		return operations.NewMgmtWalletCreateInternalServerError().WithPayload(
			&models.Error{
				Code:    errorCodeWalletCreateNew,
				Message: err.Error(),
			})
	}

	return CreateNewWallet(params.Body.Nickname, params.Body.Password, c.walletStorage, newWallet)
}

func CreateNewWallet(
	nickname *string, password *string,
	storage *sync.Map,
	newWallet *wallet.Wallet,
) middleware.Responder {
	err := newWallet.Protect(*password, 0)
	if err != nil {
		return operations.NewMgmtWalletCreateInternalServerError().WithPayload(
			&models.Error{
				Code:    errorCodeWalletCreateNew,
				Message: err.Error(),
			})
	}

	bytesOutput, err := json.Marshal(newWallet)
	if err != nil {
		return operations.NewMgmtWalletCreateInternalServerError().WithPayload(
			&models.Error{
				Code:    errorCodeWalletCreateNew,
				Message: err.Error(),
			})
	}

	err = os.WriteFile(wallet.GetWalletFile(*nickname), bytesOutput, fileModeUserRW)
	if err != nil {
		return operations.NewMgmtWalletCreateInternalServerError().WithPayload(
			&models.Error{
				Code:    errorCodeWalletCreateNew,
				Message: err.Error(),
			})
	}

	storage.Store(newWallet.Nickname, newWallet)

	privK := base58.CheckEncode(newWallet.KeyPairs[0].PrivateKey)
	pubK := base58.CheckEncode(newWallet.KeyPairs[0].PublicKey)
	salt := base58.CheckEncode(newWallet.KeyPairs[0].Salt[:])
	nonce := base58.CheckEncode(newWallet.KeyPairs[0].Nonce[:])

	return operations.NewMgmtWalletCreateOK().WithPayload(
		&models.Wallet{
			Nickname: &newWallet.Nickname,
			Address:  &newWallet.Address,
			KeyPairs: []*models.WalletKeyPairsItems0{{
				PrivateKey: &privK,
				PublicKey:  &pubK,
				Salt:       &salt,
				Nonce:      &nonce,
			}},
			Balance: 0,
		})
}
