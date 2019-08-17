package main

type User struct {
	IdUser     int    `json:"id_user"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
	Mail       string `json:"mail"`
	IdUserType int    `json:"id_user_type"`
}
