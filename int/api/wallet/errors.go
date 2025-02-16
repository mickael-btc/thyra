package wallet

const (
	errorCodeWalletAlreadyExists    = "Wallet-0001"
	errorCodeWalletWrongPassword    = "Wallet-0002"
	errorCodeGetWallet              = "Wallet-0003"
	errorCodeCanceled               = "Wallet-0004"
	errorCodeWalletCreateNoNickname = "Wallet-1001"
	errorCodeWalletCreateNoPassword = "Wallet-1002"
	errorCodeWalletCreateNew        = "Wallet-1003"
	errorCodeWalletDeleteNoNickname = "Wallet-2001"
	errorCodeWalletDeleteFile       = "Wallet-2002"
	errorCodeWalletImportNew        = "Wallet-3001"
	errorCodeWalletGetWallets       = "Wallet-4001"
	errorCodeWalletGetBalance       = "Wallet-4002"
	errorCodeWalletAlreadyImported  = "Wallet-4003"
)
