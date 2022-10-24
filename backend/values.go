package backend

const (
	ERR_USER_ALREADY_EXIST    = "user already exist"
	ERR_USER_NOT_EXIST        = "user not exist"
	ERR_WRONG_CREDENTIALS     = "wrong credentials"
	ERR_INVALID_PASSWORD      = "password must be 4 or more characters length"
	ERR_MISSING_AUTH_HEADER   = "missing auth header or wrong header format"
	ERR_INVALID_ACCESS_TOKEN  = "invalid access token"
	ERR_ACCESS_TOKEN_EXPIRED  = "access token expired"
	ERR_INVALID_REFRESH_TOKEN = "invalid refresh token"
	ERR_REFRESH_TOKEN_EXPIRED = "refresh token expired"
	ERR_FAILED_PARSE_TOKEN = "failed parsing auth token"
	ERR_LOGIN_NOT_FOUND = "not found login in jwt"
	ERR_EXPIRE_NOT_FOUND = "not found expire in jwt"
	ACCESS_TOKEN_LIFE_MILLIS  = 1 * 60 * 1000
	REFRESH_TOKEN_LIFE_MILLIS = 5 * 60 * 1000
	SIGN_KEY                  = "some_secret_key"
)

var DEVS = [...]developer {
	{
		Name:       "Oleg",
		Department: "Backend",
	},
	{
		Name:       "Stanislav",
		Department: "Frontend",
	},
	{
		Name:       "Denis",
		Department: "Backend",
	},
	{
		Name:       "Oleg",
		Department: "Design",
	},
	{
		Name:       "Max",
		Department: "Mobile",
	},
	{
		Name:       "Fil",
		Department: "Frontend",
	},
	{
		Name:       "Sergey",
		Department: "Backend",
	},
	{
		Name:       "Andrew",
		Department: "Backend",
	},
	{
		Name:       "Alexey",
		Department: "Frontend",
	},
	{
		Name:       "Georgy",
		Department: "Frontend",
	},
}