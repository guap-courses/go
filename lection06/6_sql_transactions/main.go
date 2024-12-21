package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	// Создаем пул соединений с PostgreSQL
	dsn := "postgres://username:password@localhost:5442/guap_course?sslmode=disable"
	dbpool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	defer dbpool.Close()

	err = transferMoney(context.Background(), dbpool, 3, 2, 100.0)
	if err != nil {
		log.Fatal("Transfer failed: ", err)
	} else {
		fmt.Println("Transfer completed successfully!")
	}
}

// Функция для перевода денег между пользователями
func transferMoney(ctx context.Context, pool *pgxpool.Pool, senderID, receiverID int, amount float64) error {
	// Начинаем транзакцию
	tx, err := pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	// 1. Проверяем, достаточно ли средств у отправителя
	var senderBalance float64
	err = tx.QueryRow(context.Background(), "SELECT balance FROM users WHERE id = $1", senderID).Scan(&senderBalance)
	if err != nil {
		return fmt.Errorf("failed to get sender balance: %v", err)
	}
	if senderBalance < amount {
		return fmt.Errorf("insufficient funds")
	}

	// 2. Списываем деньги с отправителя
	_, err = tx.Exec(context.Background(), "UPDATE users SET balance = balance - $1 WHERE id = $2", amount, senderID)
	if err != nil {
		return fmt.Errorf("failed to deduct money from sender: %v", err)
	}

	// 3. Зачисляем деньги получателю
	_, err = tx.Exec(context.Background(), "UPDATE users SET balance = balance + $1 WHERE id = $2", amount, receiverID)
	if err != nil {
		return fmt.Errorf("failed to credit money to receiver: %v", err)
	}

	// 4. Создаем запись о транзакции
	_, err = tx.Exec(context.Background(), "INSERT INTO transactions (sender_id, receiver_id, amount, status) VALUES ($1, $2, $3, $4)", senderID, receiverID, amount, "completed")
	if err != nil {
		return fmt.Errorf("failed to insert transaction: %v", err)
	}

	// Если все операции прошли успешно, коммитим транзакцию
	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}
