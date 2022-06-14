package model

type Users struct {
	TOTAL   uint64 `json:"total"`
	USUARIO []User `json:"usuarios"`
}

type User struct {
	NOMBRE   string `json:"nombre"`
	CORREO   string `json:"correo"`
	ROL      string `json:"rol"`
	PASSWORD string `json:"password"`
	ESTADO   uint64 `json:"estado"`
	PADRE    string `json:"padre"`
	UID      uint64 `json:"uid"`
}

type Videos struct {
	TOTAL     uint64      `json:"total"`
	PRODUCTOS []Productos `json:"productos"`
}

type Productos struct {
	ID        uint64 `json:"_id"`
	NOMBRE    string `json:"nombre"`
	USUARIO   User   `json:"usuario"`
	CATEGORIA Tipos  `json:"categoria"`
	URL       string `json:"url"`
}

type SimpleVideo struct {
	ID        uint64 `json:"_id"`
	NOMBRE    string `json:"nombre"`
	USUARIO   uint64 `json:"usuario"`
	CATEGORIA uint64 `json:"categoria"`
	URL       string `json:"url"`
}

type Categorias struct {
	TOTAL      uint64  `json:"total"`
	CATEGORIAS []Tipos `json:"categorias"`
}

type Tipos struct {
	ID      uint64 `json:"_id"`
	NOMBRE  string `json:"nombre"`
	USUARIO User   `json:"usuario"`
}

type TokenDetails struct {
	TOKEN        string `json:"token"`
	GENERATED_AT string `json:"generated_at"`
	EXPIRES_AT   string `json:"expires_at"`
	USUARIO      User   `json:"usuario"`
}
