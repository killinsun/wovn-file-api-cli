package repository

import (
	"context"
)

type IWovnAPIRepository interface {
	GetTranslated(ctx context.Context, endpoint, apiOptionQuery, apiKey string) (string, error)
	SendReport(endpoint, url, projectToken, apiKey string) (int, error)
}