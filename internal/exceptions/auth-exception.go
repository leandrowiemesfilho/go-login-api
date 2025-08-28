package exceptions

type InvalidCredentialsErr struct {
}

func (e InvalidCredentialsErr) Error() string {
	return "Invalid Credentials"
}
