package backend

type user struct {
	Id int64 `json:"id"`
	Login string `json:"login"`
	Password string `json:"password"`
	AccessToken string `json:"access_token"`
	AccessTokenExpire int64 `json:"access_token_expire"`
	RefreshToken string `json:"refresh_token"`
	RefreshTokenExpire int64 `json:"refresh_token_expire"`
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type success struct {
	Success bool `json:"success"`
}

type jwtToken struct {
	Token string `json:"token"`
	ExpiresAt int64 `json:"expiresAt"`
}

type jwtTokensPair struct {
	AccessToken jwtToken `json:"accessToken"`
	RefreshToken jwtToken `json:"refreshToken"`
}

type authBody struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

type developer struct {
    Id int64 `json:"id"`
	Name string `json:"name"`
	Department string `json:"department"`
}
