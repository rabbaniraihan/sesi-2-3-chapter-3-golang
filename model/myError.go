package model

type MyError struct {
	Err string `json:"error"`
}

func (me MyError) Error() string {
	return me.Err
}

var (
	ErrorInvalidUsernameOrPassword = MyError{
		Err: "invalid email / password",
	}

	ErrorInvalidToken = MyError{
		Err: "invalid token",
	}
)
