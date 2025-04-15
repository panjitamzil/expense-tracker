package main

import (
	"testing"
	"time"
)

func TestAddExpense(t *testing.T) {
	expenses := []Expense{}
	updatedExpenses, newID, err := addExpense(expenses, "Lunch", 20.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if newID != 1 {
		t.Errorf("Expected ID 1, got %d", newID)
	}
	if len(updatedExpenses) != 1 || updatedExpenses[0].Description != "Lunch" || updatedExpenses[0].Amount != 20.0 {
		t.Errorf("Expected one expense with description 'Lunch' and amount 20.0, got %v", updatedExpenses)
	}

	// Test negative amount
	_, _, err = addExpense(expenses, "Test", -5.0)
	if err == nil {
		t.Errorf("Expected error for negative amount, got nil")
	}
}

func TestUpdateExpense(t *testing.T) {
	expenses := []Expense{{ID: 1, Description: "Old", Amount: 10.0}}
	updatedExpenses, err := updateExpense(expenses, 1, "New", 15.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if updatedExpenses[0].Description != "New" || updatedExpenses[0].Amount != 15.0 {
		t.Errorf("Expected updated expense with description 'New' and amount 15.0, got %v", updatedExpenses[0])
	}

	// Test ID not found
	_, err = updateExpense(expenses, 999, "Test", 10.0)
	if err == nil {
		t.Errorf("Expected error for non-existent ID, got nil")
	}
}

func TestDeleteExpense(t *testing.T) {
	expenses := []Expense{{ID: 1, Description: "Lunch", Amount: 20.0}}
	updatedExpenses, err := deleteExpense(expenses, 1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(updatedExpenses) != 0 {
		t.Errorf("Expected no expenses, got %v", updatedExpenses)
	}

	// Test ID not found
	_, err = deleteExpense(expenses, 999)
	if err == nil {
		t.Errorf("Expected error for non-existent ID, got nil")
	}
}

func TestCalculateSummary(t *testing.T) {
	expenses := []Expense{
		{ID: 1, Amount: 20.0},
		{ID: 2, Amount: 10.0},
	}
	total := calculateSummary(expenses, 0)
	if total != 30.0 {
		t.Errorf("Expected total 30.0, got %f", total)
	}

	// Test for specific month
	now := time.Now()
	expenses = []Expense{
		{ID: 1, Date: now, Amount: 20.0},
		{ID: 2, Date: now.AddDate(0, -1, 0), Amount: 10.0},
	}
	total = calculateSummary(expenses, int(now.Month()))
	if total != 20.0 {
		t.Errorf("Expected total 20.0 for current month, got %f", total)
	}
}
