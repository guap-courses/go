package repository

import (
	"context"

	"github.com/guap-courses/go/lection06/7_clean_arch/internal/domain"
)

const selectCandlesQuery = `SELECT instrument, period, ts, open, high, low,
		close FROM candles_1m`

func (q *repo) Candles(ctx context.Context) ([]domain.Candle, error) {
	rows, err := q.pool.Query(ctx, selectCandlesQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var candles []domain.Candle
	for rows.Next() {
		var c domain.Candle
		err = rows.Scan(&c.Ticker, &c.Period, &c.TS, &c.Open, &c.High, &c.Low, &c.Close)
		if err != nil {
			return nil, err
		}
		candles = append(candles, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return candles, nil
}

const selectCandlesByTickerQuery = `SELECT instrument, period, ts, open, high, low,
		close FROM candles_1m WHERE instrument = $1`

func (q *repo) CandlesByTicker(ctx context.Context, ticker string) ([]domain.Candle, error) {
	rows, err := q.pool.Query(ctx, selectCandlesByTickerQuery, ticker)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var candles []domain.Candle
	for rows.Next() {
		var c domain.Candle
		err = rows.Scan(&c.Ticker, &c.Period, &c.TS, &c.Open, &c.High, &c.Low, &c.Close)
		if err != nil {
			return nil, err
		}
		candles = append(candles, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return candles, nil
}
