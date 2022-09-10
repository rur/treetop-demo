package inline

// Emulate serverside persistence using an encoded cookie
//

import (
	"errors"
	"net/http"
	"time"

	errF "github.com/pkg/errors"
)

const formDataCookieName = "inline-form-data"

func loadFormData(w http.ResponseWriter, req *http.Request) (*FormData, error) {
	var data *FormData

	if cookie, err := req.Cookie(formDataCookieName); err == nil {
		data = &FormData{}
		err := data.UnmarshalBase64([]byte(cookie.Value))
		if err != nil {
			return nil, errF.Wrap(err, "unmarshalling cookie base64 data")
		}
	} else if req.Method == "GET" || req.Method == "HEAD" {
		// There is no form data cookie so set a default (read-only requests)
		fDataCopy := defaultFormData
		data = &fDataCopy
		err := setCookieFormData(w, data)
		if err != nil {
			return nil, errF.Wrap(err, "setting base 64 encoded cookie")
		}
	} else {
		return nil, errors.New("missing cookie, redirect to a get endpoint to initialize the cookie")
	}
	return data, nil
}

// setCookieFormData will set the form data cookie in the response headers.
// Note: like any header this must be called before the headers are written
func setCookieFormData(w http.ResponseWriter, data *FormData) error {
	b64, err := data.MarshalBase64()
	if err != nil {
		return err
	}
	cookie := http.Cookie{
		Name:     formDataCookieName,
		Path:     "/inline",
		Value:    string(b64),
		Expires:  time.Now().Add(1 * time.Hour),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, &cookie)
	return nil
}
