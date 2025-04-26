package sending

type SendingData struct {
}

func (u *SendingData) TableName() string {
	return "sending"
}
