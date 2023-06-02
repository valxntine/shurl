package model

// Response is the JSON returned to the client
type Response struct {
	Url string `json:"shortUrl"`
}

// ShortUrlRequest is used to unmarshal the request from the client
type ShortUrlRequest struct {
	OriginalUrl string `json:"originalUrl"`
}
