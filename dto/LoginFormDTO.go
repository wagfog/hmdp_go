package dto

type LoginFormDTO struct {
	phone    string
	code     string
	password string
}

type LoginFormDTO2 struct {
	Phone string
	Code  string
}

func (loginFOrmDTO *LoginFormDTO2) GetPhone() string {
	return loginFOrmDTO.Phone
}

func (loginFOrmDTO *LoginFormDTO2) GetCode() string {
	return loginFOrmDTO.Code
}

func (loginFOrmDTO *LoginFormDTO) GetPhone() string {
	return loginFOrmDTO.phone
}

func (loginFOrmDTO *LoginFormDTO) GetCode() string {
	return loginFOrmDTO.code
}

func (loginFOrmDTO *LoginFormDTO) GetPassword() string {
	return loginFOrmDTO.password
}

func NewLoginFOrmDTO(p string, c string, pass string) LoginFormDTO {
	return LoginFormDTO{
		phone:    p,
		code:     c,
		password: pass,
	}
}
