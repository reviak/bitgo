package bitgo

import (
	"fmt"
	"time"
)

type WebhookType string

const (
	TransferWebhook            WebhookType = "transfer"
	PendingApprovalWebhook     WebhookType = "pendingapproval"
	AddressConfirmationWebhook WebhookType = "address_confirmation"
	BlockWebhook               WebhookType = "block"
	WalletConfirmationWebhook  WebhookType = "wallet_confirmation"
)

type WebhookState string

const (
	ActiveWebhook    WebhookState = "active"
	SuspendedWebhook WebhookState = "suspended"
	NewWebhook       WebhookState = "new"
)

type WebhookEvent struct {
	Hash     string      `json:"hash"`
	Transfer string      `json:"transfer"`
	Coin     string      `json:"coin"`
	Type     WebhookType `json:"type"`
	Wallet   string      `json:"wallet"`
}

// todo possibly merge with AddWalletWebhookResp
type Webhook struct {
	ID       string      `json:"id"`
	Label    string      `json:"label"`
	Created  time.Time   `json:"created"`
	WalletID string      `json:"walletId"`
	Coin     string      `json:"coin"`
	Type     WebhookType `json:"type"`
	URL      string      `json:"url"`
	Version  int         `json:"version"`
}

type Webhooks []Webhook

type ListWebhooks struct {
	Webhooks Webhooks `json:"webhooks"`
}

type WebhookWalletNotif struct {
	Id        string      `json:"id"`
	Type      WebhookType `json:"type"`
	Wallet    string      `json:"wallet"`
	Url       string      `json:"url"`
	Hash      string      `json:"hash"`
	Coin      string      `json:"coin"`
	State     string      `json:"state"`
	Transfer  string      `json:"transfer"`
	Webhook   string      `json:"webhook"`
	UpdatedAt time.Time   `json:"updatedAt"`
	Version   int32       `json:"version"`
}

type WebhookUserNotif struct {
	Id        string
	Type      WebhookType
	Url       string
	Hash      string
	Coin      string
	State     WebhookState
	Webhook   string
	UpdatedAt time.Time
	Version   int32
}

// List Wallet Webhooks

func (b *BitGo) ListWalletWebhooks(walletId string, params GetWalletParams) (webhooks Webhooks, err error) {
	resp := ListWebhooks{}
	err = b.get(
		fmt.Sprintf("%s/wallet/%s/webhooks", b.coin, walletId),
		nil,
		&resp)
	return resp.Webhooks, err
}

// Add Wallet Webhook
type AddWalletWebhookParams struct {
	Url              string      `json:"url,required"`
	Type             WebhookType `json:"type,required"`
	NumConfirmations int32       `json:"numConfirmations,omitempty"`
	AllToken         bool        `json:"allToken,omitempty"`
}

// todo possibly merge with Webhook
type AddWalletWebhookResp struct {
	Id       string      `json:"id"`
	Created  time.Time   `json:"created"`
	WalletId string      `json:"walletId"`
	Coin     string      `json:"coin"`
	Type     WebhookType `json:"type"`
	Url      string      `json:"url"`
	AllToken bool        `json:"allToken"`
}

func (b *BitGo) AddWalletWebhook(walletId string, params AddWalletWebhookParams) (webhook AddWalletWebhookResp, err error) {
	err = b.post(
		fmt.Sprintf("%s/wallet/%s/webhooks", b.coin, walletId),
		params,
		&webhook)
	return
}

// Remove Wallet Webhook
// it's common for wallet and user webhooks
type RemoveWebhookParams struct {
	Type WebhookType `json:"type,required"`
	Url  string      `json:"url,required"`
}

type RemoveWebhookResp struct {
	Removed int8 `json:"removed"`
}

func (b *BitGo) RemoveWalletWebhook(walletId string, params RemoveWebhookParams) (resp RemoveWebhookResp, err error) {
	err = b.delete(
		fmt.Sprintf("%s/wallet/%s/webhooks", b.coin, walletId),
		params,
		&resp)
	return
}

// Simulate Wallet Webhook
type SimulateWalletWebhookParams struct {
	WebhookId         string `json:"webhookId,required"`
	TransferId        string `json:"transferId,omitempty"`
	PendingApprovalId string `json:"pendingApprovalId,omitempty"`
}

type WalletWebhookNotifs struct {
	WebhookNotifications []*WebhookWalletNotif `json:"webhookNotifications"`
}

func (b *BitGo) SimulateWalletWebhook(walletId string, webhookId string, params SimulateWalletWebhookParams) (resp WalletWebhookNotifs, err error) {
	err = b.post(
		fmt.Sprintf("%s/wallet/%s/webhooks/%s/simulate", b.coin, walletId, webhookId),
		params,
		&resp)
	return
}

// List User Webhooks
type UserWebhook struct {
	Id                       string      `json:"id"`
	Label                    string      `json:"label"`
	Created                  time.Time   `json:"created"`
	Coin                     string      `json:"coin"`
	Type                     WebhookType `json:"type"`
	Url                      string      `json:"url"`
	Version                  int32       `json:"version"`
	State                    string      `json:"state"`
	NumConfirmations         int32       `json:"numConfirmations"`
	SuccessiveFailedAttempts int32       `json:"successiveFailedAttempts"`
}

type UserWebhooks []UserWebhook

type ListUserWebhooks struct {
	Webhooks UserWebhooks `json:"webhooks"`
}

func (b *BitGo) ListUserWebhooks() (webhooks UserWebhooks, err error) {
	resp := ListUserWebhooks{}
	err = b.get(
		fmt.Sprintf("%s/webhooks", b.coin),
		nil,
		&resp)
	return resp.Webhooks, err
}

// Add User Webhook
type AddUserWebhookParams struct {
	Url              string      `json:"url,required"`
	Type             WebhookType `json:"type,required"`
	Label            string      `json:"label"`
	NumConfirmations int32       `json:"numConfirmations"`
}

func (b *BitGo) AddUserWebhook(params AddUserWebhookParams) (webhook UserWebhook, err error) {
	err = b.post(
		fmt.Sprintf("%s/webhooks", b.coin),
		params,
		&webhook)
	return
}

// Remove User Webhook
func (b *BitGo) RemoveUserWebhook(params RemoveWebhookParams) (resp RemoveWebhookResp, err error) {
	err = b.delete(
		fmt.Sprintf("%s/wallet", b.coin),
		params,
		&resp)
	return
}

// Simulate User Webhook
type SimulateUserWebhookParams struct {
	WebhookId string `json:"webhookId,required"`
	BlockId   string `json:"blockId,omitempty"`
}

type UserNotifs []WebhookUserNotif

type UserWebhookNotifs struct {
	WebhookNotifications UserNotifs `json:"webhookNotifications"`
}

func (b *BitGo) SimulateUserWebhook(webhookId string, params SimulateUserWebhookParams) (notifs UserNotifs, err error) {
	resp := UserWebhookNotifs{}
	err = b.post(
		fmt.Sprintf("%s/webhooks/%s/simulate", b.coin, webhookId),
		params,
		&resp)
	return resp.WebhookNotifications, err
}
