package main

import (
	"bitgo"
	"log"
	"time"
)

const (
	//token      = "v2xfbfbbafa909bc4756d17e98a3b24a679706b3a812c3a82295ad621b72e99b1e6"
	token      = "v2xaba3427d8b0fb47c67cf6f70cbf2d2b74361ffce244aaac99005732c5e5582d0"
	enterprise = "5b8647e21ca5ee8203aab3855c962b4f"
)

func main() {
	//b, err := bitgo.New("test", time.Minute)
	b, err := bitgo.New("http://localhost:3080", time.Minute)
	if err != nil {
		log.Fatal(err.Error())
	}
	w, err := b.Token(token).Coin("tbtc").GenerateWallet(bitgo.GenerateWalletParams{
		Label: "auto generated TBTC wallet from SDK",
		Passphrase: "djnnfrjqrhenjqgfhjkm",
		Enterprise: enterprise,
	})

	if err != nil {
		log.Fatalf("%#v\n", err)
	}
	log.Printf("generated wallet ID: %s", w.String())
}