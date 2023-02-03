package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name             string
	url              string
	method           string
	params           []postData
	expectStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/", "GET", []postData{}, http.StatusOK},
	{"contact", "/", "GET", []postData{}, http.StatusOK},
	{"generals-quarters", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"majors-suite", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"invlaid", "/nothing_valid", "GET", []postData{}, http.StatusNotFound},
	{"post search availablity", "/search-availability", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2021-01-01"},
	}, http.StatusOK},
	{"post search availablity json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2021-01-01"},
	}, http.StatusOK},
	{"post make reservation", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "a name"},
		{key: "last_name", value: "lat"},
		{key: "email", value: "me@a.com"},
		{key: "phone", value: "546546513246"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			} else {
				if resp.StatusCode != e.expectStatusCode {
					t.Errorf("for %s expect code %d but fot %d", e.name, e.expectStatusCode, resp.StatusCode)
				}
			}
		} else {
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			} else {
				if resp.StatusCode != e.expectStatusCode {
					t.Errorf("for %s expect code %d but fot %d", e.name, e.expectStatusCode, resp.StatusCode)
				}
			}
		}
	}
}
