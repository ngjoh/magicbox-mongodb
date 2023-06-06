package model

type RoomType struct {
	Identity           string `json:"Identity"`
	PrimarySmtpAddress string `json:"PrimarySmtpAddress"`
	DisplayName        string `json:"DisplayName"`
	ResourceCapacity   int16  `json:"ResourceCapacity"`
}
