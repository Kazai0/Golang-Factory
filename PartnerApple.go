package main

// Apple implementation of Stub interface
type Apple struct{}

func (a Apple) configWallet() InfosWallet {
	return InfosWallet{MERCHANT_APPLE, PASSWORD_APPLE, CODE_APPLE}
}
