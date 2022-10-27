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
        Id:         1,
		Name:       "Oleg",
		Department: "Backend",
	},
	{
        Id:         2,
		Name:       "Stanislav",
		Department: "Frontend",
	},
    {
        Id:         3,
        Name:       "Lev",
        Department: "Mobile",
    },
	{
        Id:         4,
		Name:       "Denis",
		Department: "Backend",
	},
	{
        Id:         5,
		Name:       "Oleg",
		Department: "Design",
	},
	{
        Id:         6,
		Name:       "Max",
		Department: "Mobile",
	},
	{
        Id:         7,
		Name:       "Fil",
		Department: "Frontend",
	},
	{
        Id:         8,
		Name:       "Sergey",
		Department: "Backend",
	},
	{
        Id:         9,
		Name:       "Andrew",
		Department: "Backend",
	},
	{
        Id:         10,
		Name:       "Alexey",
		Department: "Frontend",
	},
	{
        Id:         11,
		Name:       "Georgy",
		Department: "Frontend",
	},
}
