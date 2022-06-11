package model

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

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
