package model

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

func GetAllVideos(limite string, desde string) (Videos, error) {
	var videos Videos
	var productos []Productos
	var user User
	var categoria tipos
	videoCount := 0

	query := `select id, name, url, user_id, category_id from videos limit ` + desde + `, ` + limite + `;`

	rows, err := db.Query(query)

	if err != nil {
		return videos, err
	}

	defer rows.Close()

	for rows.Next() {
		videoId := 0
		videoName := ""
		videoUrl := ""
		var user_id uint64 = 0
		var category_id uint64 = 0

		err := rows.Scan(&videoId, &videoName, &videoUrl, &user_id, &category_id)
		fmt.Println("* Scan de videos ")

		if err != nil {

			if err == sql.ErrNoRows {
				return videos, errors.New("No rows found")
			}

			return videos, err
		}

		// Query para sacar los datos del usuario relacionados con el vídeo
		queryUser := `select id, nombre from users where id =` + strconv.FormatUint(user_id, 10) + `;`

		stmtUser, errUser := db.Prepare(queryUser)

		if errUser != nil {
			return videos, errUser
		}

		defer stmtUser.Close()

		userId := 0
		userName := ""

		errUser = stmtUser.QueryRow().Scan(&userId, &userName)
		fmt.Println("* Scan de usuarios ")

		if errUser != nil {

			fmt.Println("* Estoy en el primer if de errUser")
			if errUser == sql.ErrNoRows {
				return videos, errors.New("No user found")
			}

			fmt.Println("* Voy a ejecutar el return porque " + errUser.Error())
			return videos, errUser
		}

		//Query para sacar los datos de las categorios relacionados con el vídeo
		fmt.Println("* Me he quedado justo antes de la query de categorias")
		queryCat := `select id, name from categories where id =` + strconv.FormatUint(category_id, 10) + `;`

		stmtCat, errCat := db.Prepare(queryCat)

		if errCat != nil {
			return videos, errCat
		}

		defer stmtCat.Close()

		catId := 0
		catName := ""

		fmt.Println("* Me he quedado justo antes del scan de categorias")
		errCat = stmtCat.QueryRow().Scan(&catId, &catName)
		fmt.Println("* Scan de categorias ")

		if errCat != nil {

			if errCat == sql.ErrNoRows {
				return videos, errors.New("No category found")
			}

			return videos, errCat
		}

		user = User{
			UID:    uint64(userId),
			NOMBRE: userName,
		}

		categoria = tipos{
			_ID:    uint64(catId),
			NOMBRE: catName,
		}

		producto := Productos{
			_ID:       uint64(videoId),
			NOMBRE:    videoName,
			USUARIO:   user,
			CATEGORIA: categoria,
			URL:       videoUrl,
		}

		productos = append(productos, producto)

		videoCount++
		videos = Videos{
			TOTAL:     uint64(videoCount),
			PRODUCTOS: productos,
		}
	}
	return videos, nil
}

func GetAllVideosPadre(padre string, limite string, desde string) (Videos, error) {
	var videos Videos
	var productos []Productos
	var user User
	var categoria tipos
	videoCount := 0

	fmt.Println("* El padre del video es " + padre)
	query := `select id, name, url, user_id, category_id from videos where user_id like "` + padre + `" limit ` + desde + `, ` + limite + `;`

	rows, err := db.Query(query)

	if err != nil {
		return videos, err
	}

	defer rows.Close()

	for rows.Next() {
		videoId := 0
		videoName := ""
		videoUrl := ""
		var user_id uint64 = 0
		var category_id uint64 = 0

		err := rows.Scan(&videoId, &videoName, &videoUrl, &user_id, &category_id)

		if err != nil {

			if err == sql.ErrNoRows {
				return videos, errors.New("No rows found")
			}

			return videos, err
		}

		// Query para sacar los datos del usuario relacionados con el vídeo
		queryUser := `select id, nombre from users where id =` + strconv.FormatUint(user_id, 10) + `;`

		stmtUser, errUser := db.Prepare(queryUser)

		if errUser != nil {
			return videos, errUser
		}

		defer stmtUser.Close()

		userId := 0
		userName := ""

		errUser = stmtUser.QueryRow().Scan(&userId, &userName)

		if errUser != nil {

			if errUser == sql.ErrNoRows {
				return videos, errors.New("No user found")
			}
			return videos, errUser
		}

		//Query para sacar los datos de las categorios relacionados con el vídeo
		queryCat := `select id, name from categories where id =` + strconv.FormatUint(category_id, 10) + `;`

		stmtCat, errCat := db.Prepare(queryCat)

		if errCat != nil {
			return videos, errCat
		}

		defer stmtCat.Close()

		catId := 0
		catName := ""

		errCat = stmtCat.QueryRow().Scan(&catId, &catName)

		if errCat != nil {

			if errCat == sql.ErrNoRows {
				return videos, errors.New("No category found")
			}

			return videos, errCat
		}

		user = User{
			UID:    uint64(userId),
			NOMBRE: userName,
		}

		categoria = tipos{
			_ID:    uint64(catId),
			NOMBRE: catName,
		}

		producto := Productos{
			_ID:       uint64(videoId),
			NOMBRE:    videoName,
			USUARIO:   user,
			CATEGORIA: categoria,
			URL:       videoUrl,
		}

		productos = append(productos, producto)

		videoCount++
		videos = Videos{
			TOTAL:     uint64(videoCount),
			PRODUCTOS: productos,
		}
	}
	return videos, nil
}
