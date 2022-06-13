package model

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(nombre string, password string) (TokenDetails, error) {
	var tokenDetails TokenDetails

	queryString := "select id, rol, password, estado, padre, correo from users where nombre = ?"

	stmt, err := db.Prepare(queryString)

	if err != nil {
		return tokenDetails, err
	}

	defer stmt.Close()

	userId := 0
	accountPassword := ""
	userRol := ""
	userEstado := 0
	userPadre := ""
	userCorreo := ""

	err = stmt.QueryRow(nombre).Scan(&userId, &userRol, &accountPassword, &userEstado, &userPadre, &userCorreo)

	if err != nil {

		if err == sql.ErrNoRows {
			return tokenDetails, errors.New("Invalid username or password.\r\n")
		}

		return tokenDetails, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(accountPassword), []byte(password))

	if err != nil {
		return tokenDetails, errors.New("Invalid username or password.\r\n")
	}

	queryString = "insert into authentication_tokens(user_id, auth_token, generated_at, expires_at) values (?, ?, ?, ?)"
	stmt, err = db.Prepare(queryString)

	if err != nil {
		return tokenDetails, err
	}

	defer stmt.Close()

	randomToken := make([]byte, 32)

	_, err = rand.Read(randomToken)

	if err != nil {
		return tokenDetails, err
	}

	authToken := base64.URLEncoding.EncodeToString(randomToken)

	const timeLayout = "2006-01-02 15:04:05"

	dt := time.Now()
	expirtyTime := time.Now().Add(time.Minute * 60)

	generatedAt := dt.Format(timeLayout)
	expiresAt := expirtyTime.Format(timeLayout)

	_, err = stmt.Exec(userId, authToken, generatedAt, expiresAt)

	if err != nil {
		return tokenDetails, err
	}

	userDetails := User{
		NOMBRE:   nombre,
		CORREO:   userCorreo,
		ROL:      userRol,
		PASSWORD: accountPassword,
		ESTADO:   uint64(userEstado),
		PADRE:    userPadre,
		UID:      uint64(userId),
	}

	tokenDetails = TokenDetails{
		TOKEN:        authToken,
		GENERATED_AT: generatedAt,
		EXPIRES_AT:   expiresAt,
		USUARIO:      userDetails,
	}

	return tokenDetails, nil
}

func ValidateToken(authToken string) (TokenDetails, error) {
	var UserToken TokenDetails

	queryString := `select 
						users.id,
						users.nombre,
						users.rol,
						users.padre,
						authentication_tokens.generated_at,
						authentication_tokens.expires_at                         
					from authentication_tokens
					left join users
						on authentication_tokens.user_id = users.id
					where authentication_tokens.auth_token = ?`

	stmt, err := db.Prepare(queryString)

	if err != nil {
		return UserToken, err
	}

	defer stmt.Close()

	userId := 0
	username := ""
	userrol := ""
	userPadre := ""
	generatedAt := ""
	expiresAt := ""

	err = stmt.QueryRow(authToken).Scan(&userId, &username, &userrol, &userPadre, &generatedAt, &expiresAt)

	if err != nil {

		if err == sql.ErrNoRows {
			return UserToken, errors.New("Invalid access token.\r\n")
		}

		return UserToken, err
	}

	const timeLayout = "2006-01-02 15:04:05"

	expiryTime, _ := time.Parse(timeLayout, expiresAt)
	currentTime, _ := time.Parse(timeLayout, time.Now().Format(timeLayout))

	if expiryTime.Before(currentTime) {
		return UserToken, errors.New("The token is expired.\r\n")
	}

	userDetails := User{
		UID:    uint64(userId),
		NOMBRE: username,
		ROL:    userrol,
		PADRE:  userPadre,
	}

	UserToken = TokenDetails{
		TOKEN:        authToken,
		GENERATED_AT: generatedAt,
		EXPIRES_AT:   expiresAt,
		USUARIO:      userDetails,
	}

	return UserToken, nil
}
