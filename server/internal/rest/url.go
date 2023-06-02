package rest

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/valxntine/shurl/internal/model"
)

// CreateShortUrl is the request handler for POST requests,
// it parses the body for a long url and passes that on to the service layer
// to get back a short url.
func (h *Handlers) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	body, err := parseRequest(r)

	var u model.ShortUrlRequest
	if err = json.Unmarshal(body, &u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	short, err := h.Services.ShortUrlCreater.CreateShortUrl(r.Context(), u.OriginalUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := model.Response{
		Url: short,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// GetOriginalUrl is the request handler for GET requests,
// it uses the short url string in the URL path to pass to the service layer
// to get back the original url and redirect the client to that page.
// If no long url exists we return a 5xx error.
func (h *Handlers) GetOriginalUrl(w http.ResponseWriter, r *http.Request) {
	p := chi.URLParam(r, "url")

	original, err := h.Services.OriginalUrlGetter.GetOriginalUrl(r.Context(), p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, original, http.StatusSeeOther)
}
