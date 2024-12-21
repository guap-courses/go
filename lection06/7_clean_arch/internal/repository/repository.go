package repository

import (
	"context"

	"github.com/guap-courses/go/lection06/7_clean_arch/internal/domain"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type repo struct {
	pool   *pgxpool.Pool
	logger logrus.FieldLogger
}

func NewRepository(pgxPool *pgxpool.Pool, logger logrus.FieldLogger) Repository {
	return &repo{
		pool:   pgxPool,
		logger: logger,
	}
}

type Repository interface {
	Candles(context.Context) ([]domain.Candle, error)
	CandlesByTicker(ctx context.Context, ticker string) ([]domain.Candle, error)
}
