package model

import "strconv"

type User struct {
	NOMBRE string `json:"nombre"`
	CORREO string `json:"correo"`
	ROL    string `json:"rol"`
	ESTADO bool   `json:"estado"`
	PADRE  string `json:"padre"`
	UID    uint64 `json:"uid"`
}

func GetAllUsers() ([]User, error) {
	var users []User

	query := `select nombre, correo, rol, estado, padre, id from users;`

	rows, err := db.Query(query)
	if err != nil {
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var uid uint64
		var estado bool
		var nombre, correo, rol, padre string

		err := rows.Scan(&nombre, &correo, &rol, &estado, &padre, &uid)
		if err != nil {
			return users, err
		}

		user := User{
			NOMBRE: nombre,
			CORREO: correo,
			ROL:    rol,
			ESTADO: estado,
			PADRE:  padre,
			UID:    uid,
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
		var estado bool
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
