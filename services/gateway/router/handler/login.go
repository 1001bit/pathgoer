package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/1001bit/pathgoer/services/gateway/userpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoginRequest struct {
	Login string `json:"login"`
}

type OTPRequest struct {
	Otp string `json:"otp"`
}

func LoginEmailHandler(userclient userpb.UserServiceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &LoginRequest{}
		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp, err := userclient.SendOtpEmail(r.Context(), &userpb.SendOtpEmailRequest{Login: req.Login})
		if status.Code(err) == codes.NotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		setTemporaryLoginCookies(w, resp.Email)
	}
}

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

		removeTemporaryLoginCookies(w)
		setAuthCookies(w, tokens.AccessJWT, tokens.RefreshUUID)

		w.WriteHeader(http.StatusOK)
	}
}

func setTemporaryLoginCookies(w http.ResponseWriter, email string) {
	// HACK: Also set temporaryId for better security
	http.SetCookie(w, &http.Cookie{
		Name:     "email",
		Value:    email,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
		Path:     "/login/otp",
	})
}

func removeTemporaryLoginCookies(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "email",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
}
