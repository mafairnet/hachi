package main

//APIKey The object that handles the user api keys data
type APIKey struct {
	IDApiKey     int    `json:"id_api_key"`
	AccessKey    string `json:"access_key"`
	AccessSecret string `json:"access_secret"`
	APIKeyUser   User   `json:"user"`
}
