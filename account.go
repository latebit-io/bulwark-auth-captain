package bulwarkadmin

import (
	"context"
	"fmt"
	"net/http"
)

const (
	createUrl = "accounts/create"
	readUrl   = "accounts/read"
)

// Account is used for account bulwark-auth tasks, but it's preferable to use it via the Guard struct.
type Account struct {
	client  *http.Client
	baseURL string
}

// NewAccountClient creates a client for account tasks
func NewAccountClient(baseURL string, client *http.Client) *Account {
	return &Account{
		baseURL: baseURL,
		client:  client,
	}
}

// Create will create a user account and send a verification email
func (a Account) Create(ctx context.Context, email string, isVerified bool) error {
	payload := struct {
		Email      string `json:"email"`
		IsVerified bool   `json:"isVerified"`
	}{
		Email:      email,
		IsVerified: isVerified,
	}

	err := doPost(ctx, fmt.Sprintf("%s/%s", a.baseURL, createUrl), payload, nil, a.client)

	if err != nil {
		return err
	}

	return nil
}

func (a Account) Exists(ctx context.Context, email string) bool {
	payload := struct {
		Email string `json:"email"`
	}{
		Email: email,
	}

	err := doPost(ctx, fmt.Sprintf("%s/%s", a.baseURL, readUrl), payload, nil, a.client)

	if err != nil {
		return false
	}

	return true
}
