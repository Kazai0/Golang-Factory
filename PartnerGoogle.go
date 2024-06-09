package main

// Google implementation of Stub interface
type Google struct{}

func (g Google) configWallet() InfosWallet {
	return InfosWallet{MERCHANT_GOOGLE, PASSWORD_GOOGLE, CODE_GOOGLE}
}
