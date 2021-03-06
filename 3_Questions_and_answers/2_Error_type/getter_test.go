package error_type

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDumbGetter(t *testing.T) {
	t.Run("when you don't get a 200 you get a status error", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusTeapot)
		}))
		defer svr.Close()

		_, err := DumbGetter(svr.URL)

		if err == nil {
			t.Fatal("expected an error")
		}

		// type assertion
		//got, isStatusErr := err.(BadStatusError)
		// new method
		var got BadStatusError
		isStatusErr := errors.As(err, &got)
		if !isStatusErr {
			t.Fatalf("was not a BadStatusError, got %T", err)
		}
		want := BadStatusError{URL: svr.URL, Status: http.StatusTeapot}

		if got != want {
			t.Errorf(`got "%v", want"%v"`, got, want)
		}
	})
}
