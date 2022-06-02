package model

import (
	"strconv"
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

func CreateUser(user User) error {

	query := `insert into users (nombre, correo, rol, password, estado, padre) values ( "` + user.NOMBRE + `","` + user.CORREO + `","` + user.ROL + `","` + user.PASSWORD + `","` + strconv.FormatUint(user.ESTADO, 10) + `","` + user.PADRE + `");`

	_, err := db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(user User) error {
	query := `update users set nombre = "` + user.NOMBRE + `",` + `correo = "` + user.CORREO + `",` + `password = "` + user.PASSWORD + `" where id =` + strconv.FormatUint(user.UID, 10) + `;`

	_, err := db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id uint64) error {
	query := `delete from users where id =` + strconv.FormatUint(id, 10)

	_, err := db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}
