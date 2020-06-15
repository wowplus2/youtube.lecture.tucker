package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // GET: 쿼리 파라메터 가져오기
	if name == "" {
		name = "World"
	}

	fmt.Fprintf(w, "Hello %s", name)
}

func main() {
	// Mux 인스턴스 생성
	mux := http.NewServeMux()

	// mux instance를 이용하여 라우터 구현
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World - chap02")
	})

	// HandleFunc와 Handler함수가 분리된 형태
	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{})

	http.ListenAndServe(":3000", mux)
}
