package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
}

// loadExpenses reads expenses from a JSON file
func loadExpenses(filename string) ([]Expense, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Expense{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var expenses []Expense
	if err := json.NewDecoder(file).Decode(&expenses); err != nil {
		return nil, err
	}
	return expenses, nil
}

// saveExpenses writes expenses to a JSON file
func saveExpenses(filename string, expenses []Expense) error {
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// getNextID determines the next available ID
func getNextID(expenses []Expense) int {
	if len(expenses) == 0 {
		return 1
	}
	maxID := 0
	for _, exp := range expenses {
		if exp.ID > maxID {
			maxID = exp.ID
		}
	}
	return maxID + 1
}

// addExpense adds a new expense to the list
func addExpense(expenses []Expense, description string, amount float64) ([]Expense, int, error) {
	if amount <= 0 {
		return nil, 0, fmt.Errorf("amount must be positive")
	}
	newID := getNextID(expenses)
	newExpense := Expense{
		ID:          newID,
		Date:        time.Now(),
		Description: description,
		Amount:      amount,
	}
	return append(expenses, newExpense), newID, nil
}

// updateExpense updates an existing expense by ID
func updateExpense(expenses []Expense, id int, description string, amount float64) ([]Expense, error) {
	if amount <= 0 {
		return nil, fmt.Errorf("amount must be positive")
	}
	for i, exp := range expenses {
		if exp.ID == id {
			expenses[i].Description = description
			expenses[i].Amount = amount
			expenses[i].Date = time.Now()
			return expenses, nil
		}
	}
	return nil, fmt.Errorf("expense with ID %d not found", id)
}

// deleteExpense removes an expense by ID
func deleteExpense(expenses []Expense, id int) ([]Expense, error) {
	for i, exp := range expenses {
		if exp.ID == id {
			return append(expenses[:i], expenses[i+1:]...), nil
		}
	}
	return nil, fmt.Errorf("expense with ID %d not found", id)
}

// calculateSummary computes the total expenses, optionally for a specific month
func calculateSummary(expenses []Expense, month int) float64 {
	total := 0.0
	currentYear := time.Now().Year()
	for _, exp := range expenses {
		if month == 0 || (exp.Date.Year() == currentYear && int(exp.Date.Month()) == month) {
			total += exp.Amount
		}
	}
	return total
}
