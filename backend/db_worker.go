package backend

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	connStr = "user=penachett password=samplesapp908365 dbname=jwt_sample sslmode=disable"
)

var db *sql.DB

func initDB() error {
	for i := 0; i < 5; i++ {
		var err error
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			fmt.Println("sql open: " + err.Error())
		} else {
			return nil
		}
	}
	return errors.New("failed to connect after 5 attempts")
}

func getUser(login string) (*user, error) {
	row := db.QueryRow("select * from users where login = $1", login)
	user := new(user)
	err := row.Scan(&user.Id, &user.Login, &user.Password, &user.AccessToken, &user.AccessTokenExpire, &user.RefreshToken, &user.RefreshTokenExpire)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(ERR_USER_NOT_EXIST)
		}
		fmt.Println("get user err: " + err.Error())
		return nil, err
	} else {
		return user, nil
	}
}

func checkUserExist(login string) (bool, error) {
	row := db.QueryRow("select login from users where login = $1", login)
	var tmpStr string
	err := row.Scan(&tmpStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		fmt.Println("check user exist err: " + err.Error())
		return false, err
	} else {
		return true, nil
	}
}

func registerUser(login, password string) (*user, error) {
	exist, err := checkUserExist(login)
	if err != nil {
		fmt.Println("Registration err (check user exist): " + err.Error())
		return nil, err
	}
	if exist {
		return nil, errors.New(ERR_USER_ALREADY_EXIST)
	}
	token, tokenExpire, refreshToken, refreshTokenExpire, err := generateNewTokens(login)
	if err != nil {
		return nil, err
	}
	res := db.QueryRow("insert into users  values(default,$1, $2, $3, $4, $5, $6) returning id",
		login, password, token, tokenExpire, refreshToken, refreshTokenExpire)
	var userId int64
	err = res.Scan(&userId)
	if err != nil {
		fmt.Println("Registration err: " + err.Error())
		return nil, err
	}
	user := new(user)
	user.Id = userId
	user.Login = login
	user.Password = password
	user.AccessToken = token
	user.AccessTokenExpire = tokenExpire
	user.RefreshToken = refreshToken
	user.RefreshTokenExpire = refreshTokenExpire
	return user, nil
}

func updateUserTokens(id int64, accessToken string, accessTokenExpire int64, refreshToken string, refreshTokenExpire int64) error {
	_, err := db.Exec("update users set access_token = $1, access_token_expire = $2, refresh_token = $3, refresh_token_expire = $4 where id = $5",
		accessToken, accessTokenExpire, refreshToken, refreshTokenExpire, id)
	if err != nil {
		fmt.Println("update tokens err: " + err.Error())
		return err
	}
	return nil
}