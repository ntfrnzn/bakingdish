package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gocraft/web"
	"github.com/ntfrnzn/bakingdish/models"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"testing"
)

// Return's the caller's caller info.
func callerInfo() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	parts := strings.Split(file, "/")
	file = parts[len(parts)-1]
	return fmt.Sprintf("%s:%d", file, line)
}

// Make a testing request
func newTestRequest(method, path string) (*httptest.ResponseRecorder, *http.Request) {
	request, _ := http.NewRequest(method, path, nil)
	recorder := httptest.NewRecorder()

	return recorder, request
}

func assertResponse(t *testing.T, rr *httptest.ResponseRecorder, body string, code int) {
	if gotBody := strings.TrimSpace(string(rr.Body.Bytes())); body != gotBody {
		t.Errorf("assertResponse: expected body to be %s but got %s. (caller: %s)", body, gotBody, callerInfo())
	}
	if code != rr.Code {
		t.Errorf("assertResponse: expected code to be %d but got %d. (caller: %s)", code, rr.Code, callerInfo())
	}
}

//----------

func MyNotFoundHandler(rw web.ResponseWriter, r *web.Request) {
	rw.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(rw, "My Not Found")
}

func (c *Context) HandlerWithContext(rw web.ResponseWriter, r *web.Request) {
	rw.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(rw, "My Not Found With Context")
}

func TestNoHandler(t *testing.T) {
	router := SetupTest()

	rw, req := newTestRequest("GET", "/this_path_doesnt_exist")
	router.ServeHTTP(rw, req)
	assertResponse(t, rw, "Not Found", http.StatusNotFound)
}

func TestBadMethod(t *testing.T) {
	router := SetupTest()

	rw, req := newTestRequest("POOP", "/this_path_doesnt_exist")
	router.ServeHTTP(rw, req)
	assertResponse(t, rw, "Not Found", http.StatusNotFound)
}

func TestWithHandler(t *testing.T) {
	router := SetupTest()
	router.NotFound(MyNotFoundHandler)

	rw, req := newTestRequest("GET", "/this_path_doesnt_exist")
	router.ServeHTTP(rw, req)
	assertResponse(t, rw, "My Not Found", http.StatusNotFound)
}

func TestWithRootContext(t *testing.T) {
	router := SetupTest()
	router.NotFound((*Context).HandlerWithContext)

	rw, req := newTestRequest("GET", "/this_path_doesnt_exist")
	router.ServeHTTP(rw, req)
	assertResponse(t, rw, "My Not Found With Context", http.StatusNotFound)
}

//----------

func TestSearchAll(t *testing.T) {
	router := SetupTest()

	request, _ := http.NewRequest("POST", "/search", bytes.NewReader([]byte("{}")))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != 200 {
		t.Errorf("should have gotten 200 OK, got %s.", recorder.Code)
	}

	body := strings.TrimSpace(string(recorder.Body.Bytes()))
	var results []models.RecipeId

	err := json.Unmarshal(recorder.Body.Bytes(), &results)
	if err != nil {
		t.Errorf("could not unmarshal body <<%s>>", body)
	}

	count := 1
	if len(results) < count {
		t.Errorf("expected %d query results, got %d.", count, len(results))
	}

}

func TestSearchDummy(t *testing.T) {
	router := SetupTest()

	request, _ := http.NewRequest("POST", "/search", bytes.NewReader([]byte(`{"id":"dummy_id"}`)))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != 200 {
		t.Errorf("should have gotten 200 OK, got %s.", recorder.Code)
	}

	body := strings.TrimSpace(string(recorder.Body.Bytes()))
	var results []models.RecipeId

	err := json.Unmarshal(recorder.Body.Bytes(), &results)
	if err != nil {
		t.Errorf("could not unmarshal body <<%s>>", body)
	}

	count := 1
	if len(results) != count {
		t.Errorf("expected %d query results, got %d.", count, len(results))
	}

}

func TestSearchNone(t *testing.T) {
	router := SetupTest()

	request, _ := http.NewRequest("POST", "/search", bytes.NewReader([]byte(`{"id":"does_not_exist"}`)))
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != 200 {
		t.Errorf("should have gotten 200 OK, got %s.", recorder.Code)
	}

	body := strings.TrimSpace(string(recorder.Body.Bytes()))
	var results []models.RecipeId

	err := json.Unmarshal(recorder.Body.Bytes(), &results)
	if err != nil {
		t.Errorf("could not unmarshal body <<%s>>", body)
	}

	count := 0
	if len(results) != count {
		t.Errorf("expected %d query results, got %d.", count, len(results))
	}

}
