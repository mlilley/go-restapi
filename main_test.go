package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	thing "github.com/mlilley/go-restapi/thing"
)

var app *App
var srv *httptest.Server

func before() {
	app = NewApp()
	srv = httptest.NewServer(app.router)
	initData(app)
}

func after() {
	srv.Close()
}

func initData(app *App) {
	app.thingRepo.Create(&thing.Thing{ID: 0, Val: 10})
	app.thingRepo.Create(&thing.Thing{ID: 0, Val: 20})
}

func testURL(path string) string {
	return fmt.Sprintf("%s%s", srv.URL, path)
}

func TestThingsList(t *testing.T) {
	before()
	defer after()

	res, err := http.Get(testURL("/things"))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status: %v, got: %v", http.StatusOK, res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("reading body failed: %v", err)
	}
	expected := `[{"id":1,"val":10},{"id":2,"val":20}]`
	if strings.TrimSpace(string(body)) != expected {
		t.Fatalf("expected body: '%v', got: '%v'", expected, strings.TrimSpace(string(body)))
	}
}

func TestThingsCreate(t *testing.T) {
	before()
	defer after()

	res, err := http.Post(testURL("/things"), "application/json", strings.NewReader(`{"id":0,"val":30}`))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("expected status: %v, got: %v", http.StatusCreated, res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("reading body failed: %v", err)
	}
	expected := `{"id":3,"val":30}`
	if strings.TrimSpace(string(body)) != expected {
		t.Fatalf("expected body: '%v', got: '%v'", expected, strings.TrimSpace(string(body)))
	}
}

func TestThingsGet(t *testing.T) {
	before()
	defer after()

	res, err := http.Get(testURL("/things/1"))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status: %v, got: %v", http.StatusOK, res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("reading body failed: %v", err)
	}
	expected := `{"id":1,"val":10}`
	if strings.TrimSpace(string(body)) != expected {
		t.Fatalf("expected body: '%v', got: '%v'", expected, strings.TrimSpace(string(body)))
	}
}

func TestThingsUpdate(t *testing.T) {
	before()
	defer after()

	cli := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, testURL("/things/1"), strings.NewReader(`{"id":0,"val":100}`))
	if err != nil {
		t.Fatalf("creating request failed: %v", err)
	}

	res, err := cli.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status: %v, got: %v", http.StatusOK, res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("reading body failed: %v", err)
	}
	expected := `{"id":1,"val":100}`
	if strings.TrimSpace(string(body)) != expected {
		t.Fatalf("expected body: '%v', got: '%v'", expected, strings.TrimSpace(string(body)))
	}
}

func TestThingsDelete(t *testing.T) {
	before()
	defer after()

	cli := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, testURL("/things/1"), nil)
	if err != nil {
		t.Fatalf("creating request failed: %v", err)
	}

	res, err := cli.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("expected status: %v, got: %v", http.StatusNoContent, res.StatusCode)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("reading body failed: %v", err)
	}
	expected := ""
	if strings.TrimSpace(string(body)) != expected {
		t.Fatalf("expected body: '%v', got: '%v'", expected, strings.TrimSpace(string(body)))
	}
}
