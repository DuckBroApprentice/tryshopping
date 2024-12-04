package param

//手機號+驗證碼登錄時 參數傳遞
type SmsLoginParma struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
