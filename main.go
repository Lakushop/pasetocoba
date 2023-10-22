package pasetocoba

import (
	"fmt"
	"net/http"
)

func main() {
	// Menangani permintaan untuk halaman login
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/success", successHandler)

	// Menjalankan server HTTP pada port 8080
	http.ListenAndServe(":6400", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if isValidLogin(username, password) {
			http.Redirect(w, r, "/success", http.StatusFound)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	http.ServeFile(w, r, "login.html")
}

func successHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login success!")
}

func isValidLogin(username, password string) bool {
	// Implementasi validasi login di sini
	// Contoh sederhana: jika username dan password adalah "admin", maka login berhasil
	return username == "admin123" && password == "admin1324"
}
