package model

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	NOMBRE   string `json:"nombre"`
	CORREO   string `json:"correo"`
	ROL      string `json:"rol"`
	PASSWORD string `json:"password"`
	ESTADO   uint64 `json:"estado"`
	PADRE    string `json:"padre"`
	UID      uint64 `json:"uid"`
}

type TokenDetails struct {
	TOKEN        string `json:"token"`
	GENERATED_AT string `json:"generated_at"`
	EXPIRES_AT   string `json:"expires_at"`
	USUARIO      User   `json:"usuario"`
}

func GetAllUsers() ([]User, error) {
	var users []User

	query := `select nombre, correo, rol, password, estado, padre, id from users;`

	rows, err := db.Query(query)
	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var uid uint64
		var estado uint64
		var nombre, correo, rol, password, padre string

		err := rows.Scan(&nombre, &correo, &rol, &password, &estado, &padre, &uid)
		if err != nil {
			return users, err
		}

		user := User{
			NOMBRE:   nombre,
			CORREO:   correo,
			ROL:      rol,
			PASSWORD: password,
			ESTADO:   estado,
			PADRE:    padre,
			UID:      uid,
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUser(uid uint64) (User, error) {
	var user User

	query := `select nombre, correo, rol, estado, padre, id from users u where u.id = ` + strconv.FormatUint(uid, 10)

	row, err := db.Query(query)
	if err != nil {
		return user, err
	}

	defer row.Close()

	if row.Next() {
		var uid uint64
		var estado uint64
		var nombre, correo, rol, padre string

		err := row.Scan(&nombre, &correo, &rol, &estado, &padre, &uid)
		if err != nil {
			return user, err
		}

		user = User{
			NOMBRE: nombre,
			CORREO: correo,
			ROL:    rol,
			ESTADO: estado,
			PADRE:  padre,
			UID:    uid,
		}
	}
	return user, nil
}

/*
func CreateUser(user User) error {

	query := `insert into users (nombre, correo, rol, password, estado, padre) values ( "` + user.NOMBRE + `","` + user.CORREO + `","` + user.ROL + `","` + user.PASSWORD + `","` + strconv.FormatUint(user.ESTADO, 10) + `","` + user.PADRE + `");`

	_, err := db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}
*/

func UpdateUser(user User) (User, error) {
	var updatedUser User

	query := `update users set nombre = "` + user.NOMBRE + `",` + `correo = "` + user.CORREO + `",` + `password = "` + user.PASSWORD + `" where id =` + strconv.FormatUint(user.UID, 10) + `;`

	_, err := db.Exec(query)

	if err != nil {
		return updatedUser, err
	}

	updatedUser = User{
		NOMBRE:   user.NOMBRE,
		CORREO:   user.CORREO,
		PASSWORD: user.PASSWORD,
	}

	return updatedUser, nil
}

func DeleteUser(id uint64) error {
	query := `delete from users where id =` + strconv.FormatUint(id, 10)

	_, err := db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func CreateUser(nombre string, correo string, password string) (string, error) {

	queryString := "insert into users(nombre, correo, password) values (?, ?, ?)"

	stmt, err := db.Prepare(queryString)

	if err != nil {
		return "", err
	}

	defer stmt.Close()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	_, err = stmt.Exec(nombre, correo, hashedPassword)

	if err != nil {
		return "", err
	}

	return "Success\r\n", nil

}

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
	generatedAt := ""
	expiresAt := ""

	err = stmt.QueryRow(authToken).Scan(&userId, &username, &userrol, &generatedAt, &expiresAt)

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
	}

	UserToken = TokenDetails{
		TOKEN:        authToken,
		GENERATED_AT: generatedAt,
		EXPIRES_AT:   expiresAt,
		USUARIO:      userDetails,
	}

	return UserToken, nil
}
