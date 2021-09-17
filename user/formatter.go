package user

type UserFormatter struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	LoginAs  int    `json:"login_as"`
	Token    string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Username: user.Username,
		LoginAs:  user.LoginAs,
		Token:    token,
	}

	return formatter
}
