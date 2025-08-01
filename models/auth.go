package models

type Auth struct {
	Id       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{
		Username: username,
		Password: password,
	}).First(&auth)
	return auth.Id > 0
}
