package bitgo

type Error struct {
	Message   string `json:"error"`
	RequestId string `json:"requestId"`
	Name      string `json:"name"`
}

func (e Error) Error() string {
	return e.Message
}

func IsWebhookTypeUnsupported(e error) bool {
	err, ok := e.(Error)
	if !ok {
		return false
	}
	return err.Message == "address confirmation webhooks are only supported for account-based coins"
}