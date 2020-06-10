package scratch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func makeRequest(t *testing.T, r *mux.Router, path string) (int, string) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", path, nil)

	r.ServeHTTP(w, req)
	resp := w.Result()
	bod, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		t.Fatal(readErr)
	}

	fmt.Println("GET", path, resp.StatusCode, string(bod))

	return resp.StatusCode, string(bod)
}

func TestSubMux(t *testing.T) {
	r := mux.NewRouter()

	r.HandleFunc("/foo", func(http.ResponseWriter, *http.Request) {
		fmt.Println("FOOING")
	})

	regSubRouter := r.PathPrefix("/reg_sub").Subrouter()
	regSubRouter.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		fmt.Println("REG SUB SLASH")
	})
	regSubRouter.HandleFunc("", func(http.ResponseWriter, *http.Request) {
		fmt.Println("REG SUB EMPTY")
	})

	plainSubRouter := mux.NewRouter()
	plainSubRouter.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		fmt.Println("Plain SUB SLASH")
	})
	plainSubRouter.HandleFunc("", func(http.ResponseWriter, *http.Request) {
		fmt.Println("Plain SUB Empt")
	})
	plainSubRouter.HandleFunc("/foo", func(http.ResponseWriter, *http.Request) {
		fmt.Println("Plain SUB FOO")
	})

	// r.Handle("/plain", plainSubRouter)
	r.PathPrefix("/plain").Subrouter().Handle("", plainSubRouter)

	fmt.Println("ROUTESONE\n", walkRoutes(t, r))
	fmt.Println("ROUTESTWO\n", walkRoutes(t, plainSubRouter))

	makeRequest(t, r, "/foo")
	makeRequest(t, r, "/")
	// makeRequest(t, r, "/reg_sub/")
	// makeRequest(t, r, "/reg_sub")

	makeRequest(t, r, "/plain")
	makeRequest(t, r, "/plain/foo")

	t.Fatal("MUX")
}

func walkRoutes(t *testing.T, router *mux.Router) string {
	var routes string

	walkErr := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		routes = routes + t + "\n"
		return nil
	})

	if walkErr != nil {
		t.Fatal(walkErr)
	}

	return routes
}
