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
	var categoria Tipos

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

		/// -->
		queryCount := `select count(*) from videos;`

		stmtCount, errCount := db.Prepare(queryCount)

		if errCount != nil {
			fmt.Print("There is a problem in the video count")
		}

		defer stmtCount.Close()

		videoCount := 0

		errCount = stmtCount.QueryRow().Scan(&videoCount)

		if errCount != nil {
			fmt.Print("There is a problem in the video count")
		}

		/// <-

		user = User{
			UID:    uint64(userId),
			NOMBRE: userName,
		}

		categoria = Tipos{
			ID:      uint64(catId),
			NOMBRE:  catName,
			USUARIO: user,
		}

		producto := Productos{
			ID:        uint64(videoId),
			NOMBRE:    videoName,
			USUARIO:   user,
			CATEGORIA: categoria,
			URL:       videoUrl,
		}

		productos = append(productos, producto)

		//videoCount++

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
	var categoria Tipos

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

		queryCount := `select count(*) from videos;`

		stmtCount, errCount := db.Prepare(queryCount)

		if errCount != nil {
			fmt.Print("There is a problem in the video count")
		}

		defer stmtCount.Close()

		videoCount := 0

		errCount = stmtCount.QueryRow().Scan(&videoCount)

		if errCount != nil {
			fmt.Print("There is a problem in the video count")
		}

		user = User{
			UID:    uint64(userId),
			NOMBRE: userName,
		}

		categoria = Tipos{
			ID:     uint64(catId),
			NOMBRE: catName,
		}

		producto := Productos{
			ID:        uint64(videoId),
			NOMBRE:    videoName,
			USUARIO:   user,
			CATEGORIA: categoria,
			URL:       videoUrl,
		}

		productos = append(productos, producto)

		videos = Videos{
			TOTAL:     uint64(videoCount),
			PRODUCTOS: productos,
		}
	}
	return videos, nil
}

func DeleteVideo(id uint64) error {
	query := `delete from videos where id = ` + strconv.FormatUint(id, 10) + `;`

	_, err := db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}

func CreateVideo(nombre string, url string, user_id uint64, category_id uint64) (SimpleVideo, error) {
	var video SimpleVideo

	queryString := "insert into videos(name, url, user_id, category_id) values (?, ?, ?, ?)"

	stmt, err := db.Prepare(queryString)

	if err != nil {
		return video, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(nombre, url, user_id, category_id)

	if err != nil {
		return video, err
	}

	video = SimpleVideo{
		NOMBRE:    nombre,
		URL:       url,
		USUARIO:   user_id,
		CATEGORIA: category_id,
	}
	return video, nil
}

func UpdateVideo(video SimpleVideo) (SimpleVideo, error) {
	var updatedVideo SimpleVideo
	var errVideo SimpleVideo

	query := `select name, url, category_id from videos where id = ` + strconv.FormatUint(video.ID, 10) + `;`

	stmt, err := db.Prepare(query)

	if err != nil {
		return updatedVideo, err
	}

	defer stmt.Close()

	videoName := ""
	videoUrl := ""
	videoCategory := 0

	err = stmt.QueryRow().Scan(&videoName, &videoUrl, &videoCategory)

	if err != nil {
		if err == sql.ErrNoRows {
			return updatedVideo, errors.New("No user found")
		}
		return updatedVideo, err
	}

	if video.NOMBRE == "" {
		video.NOMBRE = videoName
	}

	if video.URL == "" {
		video.URL = videoUrl
	}

	if video.CATEGORIA < 0 {
		video.CATEGORIA = uint64(videoCategory)
	}

	updatedVideo = SimpleVideo{
		NOMBRE:    video.NOMBRE,
		URL:       video.URL,
		CATEGORIA: video.CATEGORIA,
	}

	query_update := `update videos set name = "` + updatedVideo.NOMBRE + `",` + `url = "` + updatedVideo.URL + `",` + `category_id = "` + strconv.FormatUint(updatedVideo.CATEGORIA, 10) + `" where id =` + strconv.FormatUint(video.ID, 10) + `;`

	_, errUpdate := db.Exec(query_update)

	if errUpdate != nil {
		return errVideo, errUpdate
	}

	return updatedVideo, nil
}
