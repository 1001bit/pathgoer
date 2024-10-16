package usermodel

import (
	"context"
	"database/sql"
)

type Profile struct {
	Name string
	Date string
}

type Credentials struct {
	Name  string
	Email string
}

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) GetProfile(ctx context.Context, name string) (*Profile, error) {
	profile := &Profile{}

	err := us.db.QueryRowContext(ctx, "SELECT name, date FROM users WHERE LOWER(name) = LOWER($1)", name).Scan(&profile.Name, &profile.Date)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (us *UserStore) GetCredentials(ctx context.Context, login string) (*Credentials, error) {
	creds := &Credentials{}

	err := us.db.QueryRowContext(
		ctx,
		"SELECT name, email FROM users WHERE LOWER(name) = LOWER($1) OR LOWER(email) = LOWER($1)",
		login,
	).Scan(&creds.Name, &creds.Email)

	if err != nil {
		return nil, err
	}

	return creds, nil
}
