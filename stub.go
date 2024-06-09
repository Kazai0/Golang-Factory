package main

type InfosWallet struct  {
	merchant string
	password string
	code string
}

// Stub interface with configWallet method
type Stub interface {
	configWallet() InfosWallet
}
