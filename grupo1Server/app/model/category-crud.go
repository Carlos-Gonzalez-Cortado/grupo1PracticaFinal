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

func DeleteCategory(id uint64) error {
	query := `delete from categories where id = ` + strconv.FormatUint(id, 10) + `;`

	_, err := db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func CreateCategory(nombre string, userId uint64) (Tipos, error) {
	var categoria Tipos

	queryString := "insert into categories(name, user_id) values (?, ?)"

	stmt, err := db.Prepare(queryString)

	if err != nil {
		return categoria, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(nombre, userId)

	if err != nil {
		return categoria, err
	}

	categoria = Tipos{
		NOMBRE: nombre,
	}

	return categoria, nil
}

func UpdateCategory(categoria Tipos) (Tipos, error) {
	var updatedCategory Tipos
	var errCategory Tipos

	query := `select name from categories where id = ` + strconv.FormatUint(categoria.ID, 10) + `;`

	stmt, err := db.Prepare(query)

	if err != nil {
		return updatedCategory, err
	}

	defer stmt.Close()

	categoryName := ""

	err = stmt.QueryRow().Scan(&categoryName)

	if err != nil {
		if err == sql.ErrNoRows {
			return updatedCategory, errors.New("No category found")
		}
		return updatedCategory, err
	}

	if categoria.NOMBRE == "" {
		categoria.NOMBRE = categoryName
	}

	updatedCategory = Tipos{
		NOMBRE: categoria.NOMBRE,
	}

	query_update := `update categories set name = "` + updatedCategory.NOMBRE + `" where id =` + strconv.FormatUint(categoria.ID, 10) + `;`

	_, errUpdate := db.Exec(query_update)

	if errUpdate != nil {
		return errCategory, errUpdate
	}

	return updatedCategory, nil
}
