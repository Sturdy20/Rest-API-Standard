package models

type ResponseCompany struct {
	CpnName    string `json:"cpn_name"`
	CpnAddress string `json:"cpn_address"`
}

type ResponseEmployee struct {
	EpyID       string `json:"epy_id"`
	EpyName     string `json:"epy_name"`
	EpyPosition string `json:"epy_position"`
	EpyEmail    string `json:"epy_email"`
	EpyPhone    string `json:"epy_phone"`
	CpnID       string `json:"cpn_id"`
}
