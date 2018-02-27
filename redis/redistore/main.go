package redistore

import (
	"gopkg.in/boj/redistore.v1"
	"os"
	"net/http"
	"time"
	"github.com/gorilla/sessions"
	"github.com/gorilla/mux"
	"log"
)

var store, err = redistore.NewRediStore(10, "tcp", ":6379", "", []byte(os.Getenv("SESSION_SECRET")))
var users = map[string]string{"andy": "passme", "admin": "password"}

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	if (session.Values["Authenticated"] != nil) && session.Values["Authenticated"] != false {
		w.Write([]byte(time.Now().String()))
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	if originalPassword, ok := users[username]; ok {
		if password == originalPassword {
			session.Values["authenticated"] = true
			session.Save(r, w)
		} else {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}

	w.Write([]byte("Logging Successfully"))
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session.id")
	session.Options.MaxAge = -1
	sessions.Save(r, w)
	w.Write([]byte(""))

}

func main() {

	defer store.Close()
	r := mux.NewRouter()

	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/healthcheck", HealthcheckHandler)
	r.HandleFunc("/logout", LogoutHandler)

	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127,0,0,1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}


