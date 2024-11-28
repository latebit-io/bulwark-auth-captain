package bulwarkadmin

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"testing"
)

const baseUri = "http://localhost:8080"

func TestAccountCreate(t *testing.T) {
	client := &http.Client{}
	id := uuid.New()
	captain := NewCaptain(baseUri, client)
	ctx := context.Background()
	err := captain.Account.Create(ctx, fmt.Sprintf("%s@latebit.io", id.String()), true)
	if err != nil {
		t.Error(err)
	}
}
