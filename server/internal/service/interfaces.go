package service

import "context"

// ShortUrlCreator is an interface defining the methods used to create short urls.
type ShortUrlCreator interface {
	// Create takes a long and short url passed from the rest layer and forwards that
	// onto the repository layer for them to be saved in the database.
	Create(context.Context, string, string) error
}

// OriginalUrlGetter is an interface defining methods used to get long urls.
type OriginalUrlGetter interface {
	// Get takes a short url and returns the original, long url
	Get(context.Context, string) (string, error)
	// GetExistingShortUrl takes in a long url, and returns the short url associated with it, if it exists
	GetExistingShortUrl(context.Context, string) (string, error)
}
