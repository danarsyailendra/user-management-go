package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Struct untuk representasi data
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func main() {
	// Inisialisasi router menggunakan Mux
	router := mux.NewRouter()

	// Mengatur route untuk endpoint "/users"
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	// Menjalankan server pada port 8080
	http.Handle("/", router)
	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

// Handler untuk mendapatkan semua pengguna
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Handler untuk mendapatkan satu pengguna berdasarkan ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Mendapatkan parameter dari URL
	for _, user := range users {
		if user.ID == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(nil) // Mengirim tanggapan kosong jika tidak ditemukan
}

// Handler untuk membuat pengguna baru
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	json.NewEncoder(w).Encode(users)
}

// Handler untuk memperbarui pengguna berdasarkan ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, user := range users {
		if user.ID == params["id"] {
			users[index] = User{
				ID:       params["id"],
				Username: user.Username,
				Email:    user.Email,
			}
			json.NewEncoder(w).Encode(users)
			return
		}
	}
	json.NewEncoder(w).Encode(users)
}

// Handler untuk menghapus pengguna berdasarkan ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, user := range users {
		if user.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			json.NewEncoder(w).Encode(users)
			return
		}
	}
	json.NewEncoder(w).Encode(users)
}
