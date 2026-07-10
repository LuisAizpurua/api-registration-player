package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/LuisAizpurua/test/util"
)

func index(w http.ResponseWriter, r *http.Request) {
	util.GenerateCookie(w, r)

	tmpl, err := template.ParseFiles("public/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func submitHandle(w http.ResponseWriter, r *http.Request) {
	_e := util.Env()

	data := util.Data{
		Name:      r.FormValue("name"),
		Email:     r.FormValue("email"),
		Phone:     r.FormValue("phone"),
		SessionID: util.SessionID,
	}

	body, _ := json.Marshal(data)

	authURL := fmt.Sprintf("https://%s.supabase.co/auth/v1/token?grant_type=password", _e.ID)
	token, err := util.FetchToken(authURL, _e.APIKEY)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	url := fmt.Sprintf("https://%s.supabase.co/rest/v1/users", _e.ID)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("apikey", _e.APIKEY)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := util.Client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	fmt.Fprintf(w, "Status: %s", resp.Status)
}

func notification(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("public/notification.html")
	if err != nil {
		http.Error(w, "Plantilla incorrecta", http.StatusBadRequest)
		return
	}

	var msg *util.Message

	switch r.URL.Query().Get("type") {
	case "ntf":
		msg = &util.Message{
			Title:       "¡Inscripción exitosa!",
			Description: "Te enviamos un email con la informacion de la birria. Te esperamos en la cancha.",
		}
	default:
		msg = &util.Message{
			Title:       "¡Ya estas inscrito!",
			Description: "Verificar mensaje en email con la informacion de la birria. Gracias!",
		}
	}
	tmpl.Execute(w, msg)
}

func mapping(next *http.ServeMux) {
	next.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	next.HandleFunc("/", index)
	next.HandleFunc("/submit", submitHandle)
	next.HandleFunc("/notification", notification)
	next.HandleFunc("/health", util.HealthHandler)
}

func middle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s | Desde: %s | UA: %s | Cookie: %v",
			r.Method, r.URL.Path, r.URL.RawQuery, r.RemoteAddr, r.UserAgent(), r.Cookies())

		cookie, err := r.Cookie("session_id")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		util.SessionID = cookie.Value
		resp, _ := util.GetUserBySessionID(cookie.Value)
		if resp == nil {
			next.ServeHTTP(w, r)
			return
		}
		defer resp.Body.Close()

		var users []util.Data
		json.NewDecoder(resp.Body).Decode(&users)

		if len(users) > 0 {
			if r.URL.Path == "/" {
				http.Redirect(w, r, "/notification", http.StatusFound)
				return
			}
		} else if r.URL.Path != "/" && r.URL.Path != "/submit" {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	_e := util.Env()

	mux := http.NewServeMux()
	mapping(mux)

	log.Printf("Conexion con la web mediante http://%s:%s \n", _e.NODE_IP, _e.SVC_PORT)
	log.Println("Servidor iniciado en :8080")

	log.Fatal(http.ListenAndServe(":8080", middle(mux)))
}
