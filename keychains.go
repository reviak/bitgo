package bitgo

import (
	"encoding/json"
	"fmt"
	"github.com/tyler-smith/go-bip32"
	"math/rand"
)

const MaxUint32 = ^uint32(0)

// List Keychains
type ListKeychainsParams struct {
	Limit  int32  `url:"limit,omitempty"`
	PrevId string `url:"prevId,omitempty"`
}

// keychain key
type Key struct {
	EncryptedPrv string   `json:"encryptedPrv"`
	Id           string   `json:"id"`
	Pub          string   `json:"pub"`
	Users        []string `json:"users"`
}

func (k Key) String() string {
	b, _ := json.MarshalIndent(&k, "", "  ")
	return string(b)
}

type ListKeychainsResp struct {
	Keys            []Key  `json:"keys"`
	Limit           int32  `json:"limit"`
	NextBatchPrevId string `json:"nextBatchPrevId"`
}

func (b *BitGo) ListKeyChains(params ListKeychainsParams) (resp ListKeychainsResp, err error) {
	err = b.get(
		fmt.Sprintf("%s/key", b.coin),
		params,
		&resp)
	return
}

// Get Keychain
type KeychainResp struct {
	Id    string `json:"id"`
	Pub   string `json:"pub"`
	Users []struct {
		User         string `json:"user"`
		EncryptedPrv string `json:"encryptedPrv"`
	} `json:"users"`
}

func (b *BitGo) GetKeychain(id string) (resp KeychainResp, err error) {
	err = b.get(
		fmt.Sprintf("%s/key/%s", b.coin, id),
		nil,
		&resp)
	return
}

// Create Keychain
type LocalKeychain struct {
	Pub  string `json:"pub"`
	Priv string `json:"priv"`
}

func (b *BitGo) CreateKeychain(seed []byte) (chain LocalKeychain, err error) {
	// todo refactor function
	if len(seed) == 0 {
		seed, err = bip32.NewSeed()
		if err != nil {
			return
		}
	}
	privKey, _ := bip32.NewMasterKey(seed)
	chain.Priv = privKey.String()
	var pubKey *bip32.Key
	var e error
	// iterating until we get proper result in case of error
	for i := 0; i < int(MaxUint32); i++ {
		rnd := rand.Uint32()
		pubKey, e = privKey.NewChildKey(rnd)
		if e == nil {
			chain.Pub = pubKey.String()
			return
		}
	}
	return
}

// Create BitGo Keychain
// Creates a new keychain on BitGo’s servers and returns
// the public keychain to the caller.

// todo maybe add specific factory or so
type CreateBitgoKeychainParams struct {
	Source        string `json:"source,required"`         //The origin of the keychain. Must be bitgo for a BitGo key.
	Enterprise    string `json:"enterprise,required"`     //(only for Eth)	The enterprise id of the BitGo key
	NewFeeAddress bool   `json:"newFeeAddress,omitempty"` //Create a new keychain instead of fetching enterprise key (only for Ethereum)
}

// todo maybe merge with Key struct
type KeychainBitgo struct {
	Id      string   `json:"id"`
	Users   []string `json:"users,omitempty"`
	Pub     string   `json:"pub,omitempty"`
	IsBitGo bool     `json:"isBitGo"`
}

func (b *BitGo) CreateBitgoKeychain(params CreateBitgoKeychainParams) (chain KeychainBitgo, err error) {
	err = b.post(
		fmt.Sprintf("%s/key", b.coin),
		params,
		&chain)
	return
}

// Create Backup Keychain
// Creates a new backup keychain on a third party specializing in key
// recovery services. This keychain will be stored on the third party
// service and usable for recovery purposes only.
type CreateBackupKeychain struct {
	Source   string `json:"source,required"`   //String	Yes	The origin of the keychain. Must be backup for a backup key.
	Provider string `json:"provider,required"` //String	Yes	The backup provider for the keychain, e. g. cme.
}

type BackupKeychain struct {
	Id  string `json:"id"`
	Pub string `json:"pub"`
}

func (b *BitGo) CreateBackupKeychain(params CreateBackupKeychain) (chain BackupKeychain, err error) {
	err = b.post(
		fmt.Sprintf("%s/key", b.coin),
		params,
		&chain)
	return
}

// Add Keychain
// This API call allows you to add a public keychain on BitGo’s server.
// Only the public key parameter is required. If using the 'Create Keychain’
// API call, you do not need to include the source parameter. You must add
// a keychain to BitGo before a wallet can be created with a keychain.
type AddKeychainParams struct {
	Pub          string `json:"pub,required"` //String	Yes	The keychain’s public key.
	EncryptedPrv string `json:"encryptedPrv"` //String	No	The keychain’s encrypted private key.
	Source       string `json:"source"`       //String	No	The origin of the keychain, e. g. bitgo or backup
}

type AddKeychainResp struct {
	Id   string   `json:"id"`
	User []string `json:"user"`
	Pub  string   `json:"pub"`
}

func (b *BitGo) AddKeychain(params AddKeychainParams) (chain AddKeychainResp, err error) {
	err = b.post(
		fmt.Sprintf("%s/key", b.coin),
		params,
		&chain)
	return
}
