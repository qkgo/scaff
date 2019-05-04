package dto

type SignData struct {
	Id           int      `json:"id" `
	Email        string   `json:"email" `
	Custom       int64    `json:"custom" `
	CustomerCode string   `json:"customerCode" `
	Granted      string   `json:"granted" `
	Level        int64    `json:"level" `
	PhoneNumber  string   `json:"phoneNumber"`
	RoleCode     []string `json:"role" `
	GroupCode    []string `json:"groupCode"`
}
