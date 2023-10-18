package kup

func Name() string {
	return "kup"
}

type KupClient struct {
	RGAPI string
}

func NewKupClient(rgapi string) *KupClient {
	return &KupClient{RGAPI: rgapi}
}
