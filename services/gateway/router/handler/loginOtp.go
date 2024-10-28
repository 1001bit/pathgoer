package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/1001bit/pathgoer/services/gateway/userpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func LoginOTPHandler(userclient userpb.UserServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &OTPRequest{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Extract email from cookie
		emailCookie, err := r.Cookie("email")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokens, err := userclient.VerifyOtp(r.Context(), &userpb.VerifyOtpRequest{
			Email: emailCookie.Value,
			Otp:   req.Otp,
		})
		if status.Code(err) == codes.NotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		removeCookie(w, r, "email", "/login/otp")
		setAuthCookies(w, tokens.AccessJWT, tokens.RefreshUUID)

		w.WriteHeader(http.StatusOK)
	}
}

func setAuthCookies(w http.ResponseWriter, access, refresh string) {
	// INSECURE
	http.SetCookie(w, &http.Cookie{
		Name:     "access",
		Value:    access,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
		Path:     "/",
	})

	// INSECURE
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh",
		Value:    refresh,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
		Path:     "/auth",
		Expires:  time.Now().Add(30 * 24 * time.Hour),
	})
}