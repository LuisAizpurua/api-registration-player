package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

var (
	Client            = &http.Client{Timeout: 10 * time.Second}
	CurrentSession    *SessionResponse
	SessionObtainedAt time.Time
	SessionID         string
)

func GetCredential() Credentials {
	//err := godotenv.Load()
	//if err != nil {
	//	panic("No se pudo cargar .env")
	//}

	return Credentials{
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
	}
}

func Env() Environment {
	//err := godotenv.Load()
	//if err != nil {
	//	panic("No se pudo cargar .env")
	//}

	return Environment{
		APIKEY:   os.Getenv("APIKEY"),
		ID:       os.Getenv("ID"),
		SVC_PORT: os.Getenv("SVC_PORT"),
		NODE_IP:  os.Getenv("NODE_IP"),
	}
}

func GetUserBySessionID(sessionId string) (*http.Response, error) {
	targetUrl := fmt.Sprintf("https://%s.supabase.co/rest/v1/users?session_id=eq.%s", Env().ID, sessionId)
	req, _ := http.NewRequest("GET", targetUrl, nil)
	req.Header.Set("apikey", Env().APIKEY)
	return Client.Do(req)
}

func GenerateCookie(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session_id")
	if err != nil {
		SessionID = uuid.New().String()
		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    SessionID,
			Path:     "/",
			HttpOnly: true,
			Secure:   true,
			MaxAge:   31536000,
			SameSite: http.SameSiteStrictMode,
		})
	}
}

func FetchToken(url string, apikey string) (string, error) {
	if CurrentSession != nil {
		expiresIn := time.Duration(CurrentSession.ExpiresIn) * time.Second
		if time.Since(SessionObtainedAt) < expiresIn {
			return CurrentSession.AccessToken, nil
		}
	}

	body, err := json.Marshal(GetCredential())
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", apikey)

	resp, err := Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var session SessionResponse
	if err := json.NewDecoder(resp.Body).Decode(&session); err != nil {
		return "", err
	}

	CurrentSession = &session
	SessionObtainedAt = time.Now()

	return CurrentSession.AccessToken, nil
}
