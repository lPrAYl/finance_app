package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var id = 1

type Transaction struct {
	ID       int
	Amount   float64
	Category string
	Date     string
	Type     string
}

func validateInput(transType string, amount float64) error {
	if amount < 0 {
		return fmt.Errorf("сумма не может быть отрицательной")
	}

	if transType != "income" && transType != "expense" {
		return fmt.Errorf("тип транзакции должен быть 'income' или 'expense'")
	}

	return nil
}

func createTransaction(scanner *bufio.Scanner) *Transaction {

	for {
		fmt.Print("Введите тип операции (income/expense): ")
		scanner.Scan()
		transType := strings.TrimSpace(scanner.Text())

		fmt.Print("Введите сумму: ")
		scanner.Scan()
		amountStr := strings.TrimSpace(scanner.Text())
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			fmt.Println("Ошибка: введите корректное число для суммы")
			continue
		}

		fmt.Print("Введите категорию: ")
		scanner.Scan()
		category := strings.TrimSpace(scanner.Text())

		err = validateInput(transType, amount)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			continue
		}

		return &Transaction{
			ID:       id,
			Amount:   amount,
			Category: category,
			Date:     "now",
			Type:     transType,
		}
	}
}

func main() {
	fmt.Println("Добро пожаловать в приложение для учета финансов!")

	scanner := bufio.NewScanner(os.Stdin)

	transaction := createTransaction(scanner)

	fmt.Printf("\nСоздана транзакция:\nID: %d\nСумма: %f\nКатегория: %s\nДата: %s\nТип: %s\n",
		transaction.ID, transaction.Amount, transaction.Category, transaction.Date, transaction.Type)
}
