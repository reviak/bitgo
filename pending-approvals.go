package bitgo

import (
	"fmt"
	"time"
)

type PendingApproval struct {
	CreateDate time.Time           `json:"createDate"`
	Creator    string              `json:"creator"`
	Enterprise string              `json:"enterprise"`
	Id         string              `json:"id"`
	Info       PendingApprovalInfo `json:"info"`
	Resolvers  []Resolver          `json:"resolvers"`
	State      string
}

type Resolver struct {
	Date time.Time `json:"date"`
	User string    `json:"user"`
}

type PolicyRuleRequest struct {
	Action        string `json:"action"`
	PolicyChanged string `json:"policyChanged"`
	Update        struct {
		Action struct {
			Type string `json:"type"`
		} `json:"action"`
		Condition struct {
			Amount int64 `json:"amount"`
		} `json:"condition"`
		Id   string `json:"id"`
		Type string `json:"type"`
	} `json:"update"`
}

type PendingApprovalInfo struct {
	PolicyRuleRequest PolicyRuleRequest `json:"policyRuleRequest"`
	Type              string            `json:"type"`
}

type PendingApprovals []PendingApproval

// get pending approval
func (b *BitGo) GetPendingApproval(approvalId string) (approval PendingApproval, err error) {
	err = b.get(
		fmt.Sprintf("%s/pendingapprovals/%s", b.coin, approvalId),
		nil,
		&approval)
	return
}

// list pending approvals
type ListPendingApprovalsParams struct {
	WalletID   string `url:"walletId,omitempty"`
	Enterprise string `url:"enterprise,omitempty"`
	AllTokens  bool   `url:"allTokens,omitempty"`
}

func (b *BitGo) ListPendingApprovals(params ListPendingApprovalsParams) (approvals PendingApprovals, err error) {
	err = b.get(
		fmt.Sprintf("%s/pendingapprovals", b.coin),
		nil,
		&approvals)
	return
}

// update pending approval
type UpdatePendingApprovalParams struct {
	State string `json:"state,required"`
	Otp   string `json:"otp,required"`
}

func (b *BitGo) UpdatePendingApproval(approvalId string, params UpdatePendingApprovalParams) (approval PendingApproval, err error) {
	err = b.put(
		fmt.Sprintf("%s/pendingapprovals/%s", b.coin, approvalId),
		params,
		&approval)
	return
}
