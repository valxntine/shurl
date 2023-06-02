package rest

import "context"

type Services struct {
	ShortUrlCreater   ShortUrlCreator
	OriginalUrlGetter OriginalUrlGetter
}

type ShortUrlCreator interface {
	CreateShortUrl(context.Context, string) (string, error)
}

type OriginalUrlGetter interface {
	GetOriginalUrl(context.Context, string) (string, error)
}
