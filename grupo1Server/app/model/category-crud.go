package model

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

func GetAllCategories(limite, desde string) (Categorias, error) {
	var categorias Categorias
	var tipos []Tipos
	var user User
	catCount := 0

	query := `select id, name, user_id from categories limit ` + desde + `,` + limite + `;`

	rows, err := db.Query(query)

	if err != nil {
		return categorias, err
	}

	for rows.Next() {
		var catId uint64 = 0
		catName := ""
		var catUser_id uint64 = 0

		err := rows.Scan(&catId, &catName, &catUser_id)

		if err != nil {

			if err == sql.ErrNoRows {
				return categorias, errors.New("No rows found")
			}

			return categorias, err
		}

		queryUser := `select nombre from users where id = ` + strconv.FormatUint(catUser_id, 10) + `;`

		stmtUser, errUser := db.Prepare(queryUser)

		if errUser != nil {
			return categorias, errUser
		}

		defer stmtUser.Close()

		userName := ""

		errUser = stmtUser.QueryRow().Scan(&userName)

		if errUser != nil {

			if errUser == sql.ErrNoRows {
				return categorias, errors.New("No user found")
			}

			return categorias, errUser
		}

		user = User{
			UID:    catUser_id,
			NOMBRE: userName,
		}

		tipo := Tipos{
			ID:      catId,
			NOMBRE:  catName,
			USUARIO: user,
		}

		tipos = append(tipos, tipo)

		catCount++
		categorias = Categorias{
			TOTAL:      uint64(catCount),
			CATEGORIAS: tipos,
		}
	}

	return categorias, nil

}

func GetAllCategoriesPadre(padre, limite, desde string) (Categorias, error) {
	var categorias Categorias
	var tipos []Tipos
	var user User

	query := `select id, name, user_id from categories where user_id like "` + padre + `" limit ` + desde + `,` + limite + `;`

	rows, err := db.Query(query)

	if err != nil {
		return categorias, err
	}

	for rows.Next() {
		var catId uint64 = 0
		catName := ""
		var catUser_id uint64 = 0

		err := rows.Scan(&catId, &catName, &catUser_id)

		if err != nil {

			if err == sql.ErrNoRows {
				return categorias, errors.New("No rows found")
			}

			return categorias, err
		}

		queryUser := `select nombre from users where id = ` + strconv.FormatUint(catUser_id, 10) + `;`

		stmtUser, errUser := db.Prepare(queryUser)

		if errUser != nil {
			return categorias, errUser
		}

		defer stmtUser.Close()

		userName := ""

		errUser = stmtUser.QueryRow().Scan(&userName)

		if errUser != nil {

			if errUser == sql.ErrNoRows {
				return categorias, errors.New("No user found")
			}

			return categorias, errUser
		}

		queryCount := `select count(*) from categories;`

		stmtCount, errCount := db.Prepare(queryCount)

		if errCount != nil {
			fmt.Print("There is a problem in the category count")
		}

		defer stmtCount.Close()

		catCount := 0

		errCount = stmtCount.QueryRow().Scan(&catCount)

		if errCount != nil {
			fmt.Print("There is a problem in the cateogory count")
		}

		user = User{
			UID:    catUser_id,
			NOMBRE: userName,
		}

		tipo := Tipos{
			ID:      catId,
			NOMBRE:  catName,
			USUARIO: user,
		}

		tipos = append(tipos, tipo)

		categorias = Categorias{
			TOTAL:      uint64(catCount),
			CATEGORIAS: tipos,
		}
	}

	return categorias, nil

}
