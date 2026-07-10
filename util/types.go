package util

type Message struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Health struct {
	Status string `json:"status"`
}

type Data struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	SessionID string `json:"session_id"`
}

type Environment struct {
	APIKEY   string
	ID       string
	SVC_PORT string
	NODE_IP  string
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SessionResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}
