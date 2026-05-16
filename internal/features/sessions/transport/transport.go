package sessions_transport

import (
	"apiserver/internal/core/domains"
	core_request "apiserver/internal/core/transport/request"
	core_response "apiserver/internal/core/transport/response"
	"net/http"
)

type SessionsService interface {
	Authenticate(email string, password string) (int, error)
	Create(userID int) (*domains.Session, error)
}

type SessionsTransport struct {
	sessionService SessionsService
}

func NewTransport(sessionService SessionsService) *SessionsTransport {
	return &SessionsTransport{
		sessionService: sessionService,
	}
}

func (t *SessionsTransport) CreateSessionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseHandler := core_response.NewResponseHandler(w)

		var request CreateSessionRequest
		if err := core_request.DecodeAndValidate(r, &request); err != nil {
			responseHandler.ErrorResponse("decode and validate", err)

			return
		}

		userID, err := t.sessionService.Authenticate(request.Email, request.Password)
		if err != nil {
			responseHandler.ErrorResponse("authenticate user", err)

			return
		}

		session, err := t.sessionService.Create(userID)
		if err != nil {
			responseHandler.ErrorResponse("create session", err)

			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    session.SessionToken,
			Expires:  session.ExpiresAt,
			HttpOnly: true,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "csrf_token",
			Value:    session.CSRFToken,
			Expires:  session.ExpiresAt,
			HttpOnly: false,
		})

		w.WriteHeader(http.StatusOK)
	}
}
