package error

type TelegramTokenNotFound struct{}

func (e TelegramTokenNotFound) Error() string {
	return "unable to find telegram token"
}
