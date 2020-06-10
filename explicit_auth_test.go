package scratch

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func checkSession(sessionId string) (user, error) {

	if sessionId == "7" {
		return "joebob", nil
	}

	return "", errors.New("Unknown / Invalid Session")
}

type authedHandler interface {
	ServeAuthedHTTP(username user, w http.ResponseWriter, r *http.Request)
}

func withAuth(h authedHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("MIDDLEWARE RUNNING")

		sessionToken := r.Header.Get("auth-token")
		fmt.Println("token", sessionToken)

		// get the session if it's valid....
		username, sessionErr := checkSession(sessionToken)
		if sessionErr != nil {
			http.Error(w, "Bad news bears, session bad", http.StatusBadRequest)
			return
		}

		h.ServeAuthedHTTP(username, w, r)

	})
}

type user string

type saveHandler struct {
	db string
}

func newSaveHandler(db string) saveHandler {
	return saveHandler{
		db,
	}
}

func (h saveHandler) ServeAuthedHTTP(username user, w http.ResponseWriter, r *http.Request) {

	fmt.Println("LOGGED IN USER: ", username)

	w.Write([]byte("SUCCESS"))
}

func TestAuthIsPassed(t *testing.T) {

	fullHandler := withAuth(newSaveHandler("bad db string"))

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/save", strings.NewReader(`{"body": "mod"}`))
	r.Header.Add("auth-token", "7")

	fullHandler.ServeHTTP(w, r)
	goodResp := w.Result()

	if goodResp.StatusCode != http.StatusOK {
		t.Fatal("failed with a valid user")
	}

	badW := httptest.NewRecorder()
	badR := httptest.NewRequest("POST", "/save", strings.NewReader(`{"body": "mod"}`))
	r.Header.Add("auth-token", "4")

	fullHandler.ServeHTTP(badW, badR)
	badResp := badW.Result()

	if badResp.StatusCode != http.StatusBadRequest {
		t.Fatal("succeeded with an invalid user")
	}

	t.Fatal("no")
}
