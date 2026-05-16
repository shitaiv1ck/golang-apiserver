package domains

import "time"

type Session struct {
	SessionToken string
	CSRFToken    string
	UserID       int
	CreatedAt    time.Time
	ExpiresAt    time.Time
}

func NewSession(sessionToken string, csfrToken string, userID int, expiresAt time.Time) *Session {
	return &Session{
		SessionToken: sessionToken,
		CSRFToken:    csfrToken,
		UserID:       userID,
		ExpiresAt:    expiresAt,
	}
}
