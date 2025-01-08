package telegramparser

type WebAppUser struct {
	Id                    int64  `json:"id"`
	IsBot                 bool   `json:"is_bot,omitempty"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name,omitempty"`
	Username              string `json:"username,omitempty"`
	LanguageCode          string `json:"language_code,omitempty"`
	IsPremium             bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu bool   `json:"added_to_attachment_menu,omitempty"`
	AllowsWriteToPM       bool   `json:"allows_write_to_pm,omitempty"`
	PhotoURL              string `json:"photo_url,omitempty"`
}

type WebAppChat struct {
	Id       int64  `json:"id"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	Username string `json:"username,omitempty"`
	PhotoURL string `json:"photo_url,omitempty"`
}

type WebAppInitData struct {
	QueryId      string
	User         WebAppUser
	Receiver     WebAppUser
	Chat         WebAppChat
	ChatType     string
	ChatInstance string
	StartParam   string
	CanSendAfter int64
	AuthDate     int64
	Hash         string
	Signature    string
}
