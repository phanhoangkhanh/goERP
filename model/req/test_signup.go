package req

type TestReq struct {
	UserName string `validate:"required"`
	Password string `validate:"required"`
	Email    string `validate:"required"`
}
