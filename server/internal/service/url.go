package service

import (
	"context"
	"net/url"
)

// UrlService along with its methods execute the business logic
// for retrieving and adding new short urls. It will call the repositories passed into it
// and orcestrate the dance of the entities to perform some logic
type UrlService struct {
	ShortUrlCreator   ShortUrlCreator
	OriginalUrlGetter OriginalUrlGetter
}

// NewUrlService returns a new instance of UrlService to be injected into the router
func NewUrlService(creator ShortUrlCreator, getter OriginalUrlGetter) UrlService {
	return UrlService{
		ShortUrlCreator:   creator,
		OriginalUrlGetter: getter,
	}
}

// CreateShortUrl passes the HTTP request through to the service layer,
// it does some basic validation such as checking for a scheme on the url and
// checking if we already have that url stored in the database, returning the already
// created short url if so.
func (s *UrlService) CreateShortUrl(ctx context.Context, original string) (string, error) {
	parsed, err := url.Parse(original)
	if err != nil {
		return "", err
	}

	if parsed.Scheme == "" {
		parsed.Scheme = "https"
	}

	longUrl := parsed.String()

	existingShort, err := s.OriginalUrlGetter.GetExistingShortUrl(ctx, longUrl)
	if err != nil {
		return "", err
	}

	if existingShort != "" {
		return existingShort, nil
	}

	short, err := createShortUrl()
	if err != nil {
		return "", err
	}

	err = s.ShortUrlCreator.Create(ctx, longUrl, short)
	if err != nil {
		return "", err
	}
	return short, nil
}

// GetOriginalUrl is a pass-through to the repository layer for getting the long url associated with a short url
func (s *UrlService) GetOriginalUrl(ctx context.Context, short string) (string, error) {
	return s.OriginalUrlGetter.Get(ctx, short)
}
