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
	b, err := bitgo.New("test", time.Minute)
	if err != nil {
		log.Fatal(err.Error())
	}

	// 0. List keychains
	chains, _ := b.Token(token).Coin("tbtc").ListKeyChains(bitgo.ListKeychainsParams{})
	for _, c := range chains.Keys {
		log.Printf("ListKeyChains: %s\n", c.String())
	}

	// 1. create local keychain
	localKeychain, err := b.Token(token).Coin("tbtc").CreateKeychain([]byte{})
	if err != nil {
		log.Fatalf("CreateKeychain: %#v\n", err)
	}
	log.Printf("Local keychain: %#v\n", localKeychain)

	// 2. Create bitgo keychain
	keychainBitgo, err := b.Token(token).Coin("tbtc").CreateBitgoKeychain(bitgo.CreateBitgoKeychainParams{
		Source:     "bitgo",
		Enterprise: enterprise,
	})
	if err != nil {
		log.Fatalf("CreateBitgoKeychain: %#v\n", err)
	}
	log.Printf("Bitgo keychain: %#v\n", keychainBitgo)

	// 3. create backup keychain
	bkpkch, err := b.Token(token).Coin("tbtc").CreateKeychain([]byte{})
	if err != nil {
		log.Fatalf("CreateBackupKeychain: %#v\n", err)
	}
	log.Printf("Bitgo keychain: %#v\n", bkpkch)
	// Encrypted private key
	// {
	//	"iv":"ggUzzmkL/08MKCvEbpXcdg==",
	//	"v":1,
	//	"iter":10000,
	//	"ks":256,
	//	"ts":64,
	//	"mode":"ccm",
	//	"adata":"",
	//	"cipher":"aes",
	//	"salt":"cFoqYCe+kWE=",
	//	"ct":"tWCCDdeIIVGjRpRJQvBXkCohnfPaoG0vjRFtEz4H22o7GnQ+0MCLnYFhFEvXWQoVq2WFkOVQ/QbQb5GEMu4sBOnjOFuya35m45mn4m3pH7qcc5yrTUiuY5ETd+m3HEomm54OUMxNik9C4crXcCly/xxapa6YOGo="
	//}
	// save
	savedLocalKeyChain, err := b.Token(token).Coin("tbtc").AddKeychain(bitgo.AddKeychainParams{
		Pub: localKeychain.Pub,
		//Priv: localKeychain.Priv,
		Source: "bitgo",
	})
	if err != nil {
		log.Fatalf("AddKeychain for local: %#v\n", err)
	}
	log.Printf("Bitgo keychain: %#v\n", savedLocalKeyChain)

	savedBkpkch, err := b.Token(token).Coin("tbtc").AddKeychain(bitgo.AddKeychainParams{
		Pub: bkpkch.Pub,
		//Priv: bkpkch.Priv,
		Source: "backup",
	})

	if err != nil {
		log.Fatalf("AddKeychain for BKP: %#v\n", err)
	}
	log.Printf("Bitgo keychain: %#v\n", savedBkpkch)
	// Add wallet
	//	Creates the user keychain locally on the machine, and encrypts it with the provided passphrase (skipped if userKey is provided).
	//	Creates the backup keychain locally on the machine.
	//	Uploads the encrypted user keychain and public backup keychain.
	//	Creates the BitGo key (and the backup key if backupXpubProvider is set) on the service.
	//	Creates the wallet on BitGo with the 3 public keys above.
	awp := bitgo.AddWalletParams{
		Label:      "User 5b6aae22df851bc89d267734 TBTC wallet",
		Enterprise: "5b8647e21ca5ee8203aab3855c962b4f",
		M:          2,
		N:          3,
		Keys:       []string{savedLocalKeyChain.Id, savedBkpkch.Id, keychainBitgo.Id},
	}
	log.Printf("Bitgo keychain: %#v\n", awp)
	w, err := b.Token(token).Coin("tbtc").Debug(true).AddWallet(awp)

	if err != nil {
		log.Fatalf("%#v\n", err.(bitgo.Error))
	}
	log.Printf("generated wallet ID: %s", w.Wallet.Wallet.ID)
}