package models

type RequestRegister struct {
	// Company
	CpnName    string `json:"cpn_name"`
	CpnAddress string `json:"cpn_address"`
	// Employee
	EmpName     string `json:"epy_name"`
	EmpPosition string `json:"epy_position"`
	EmpEmail    string `json:"epy_email"`
	EmpPhone    string `json:"epy_phone"`
}

type RequestUpdateEmployee struct {
	EpyID       string `json:"epy_id"`
	EpyName     string `json:"epy_name"`
	EpyPosition string `json:"epy_position"`
	EpyEmail    string `json:"epy_email"`
	EpyPhone    string `json:"epy_phone"`
}
