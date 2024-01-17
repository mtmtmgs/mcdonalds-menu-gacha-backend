package values

/*
メールアドレス
*/
type Email struct {
	value string
}

func NewEmail(value string) Email {
	return Email{value: value}
}

func (e *Email) Validate() error {
	return nil
}

func (e *Email) Value() string {
	return e.value
}
