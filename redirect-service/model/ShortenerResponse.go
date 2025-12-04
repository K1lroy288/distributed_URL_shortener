package model

type ShortenerResponse struct {
	LongURL  string `json:"long_url"`
	Owner_id int    `json:"owner_id"`
}
