package main

//User The object that handles the user data
type User struct {
	IDUser       int      `json:"id_user"`
	UserName     string   `json:"username"`
	Password     string   `json:"password"`
	Mail         string   `json:"mail"`
	UserUserType UserType `json:"user_type"`
}
