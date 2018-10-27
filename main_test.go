package main

// TODO

// func TestGetThings(t *testing.T) {
// 	req, err := http.NewRequest("GET", "localhost:8000/things", nil)
// 	if err != nil {
// 		t.Fatalf("create request failed: %v", err)
// 	}

// 	s := server{router: mux.NewRouter()}
// 	w := httptest.NewRecorder()
// 	s.handleGetThings()(w, req)
// 	res := w.Result()
// 	defer res.Body.Close()

// 	fmt.Printf("xx: %v", res.ContentLength)

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		t.Fatalf("read response failed: %v", err)
// 	}

// 	fmt.Printf("body: %q", body)

// 	fmt.Printf("Status: %v %v\n", res.StatusCode, res.Status)
// 	if res.StatusCode != http.StatusOK {
// 		fmt.Printf("eh?")
// 		t.Fatalf("status code not 200: %v", res.StatusCode)
// 	}

// }
