package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	filename := "expenses.json"
	expenses, err := loadExpenses(filename)
	if err != nil {
		fmt.Printf("Error loading expenses: %v\n", err)
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		description := addCmd.String("description", "", "Description of the expense")
		amount := addCmd.Float64("amount", 0.0, "Amount of the expense")
		addCmd.Parse(os.Args[2:])

		if *description == "" || *amount <= 0 {
			fmt.Println("Description and positive amount are required")
			os.Exit(1)
		}

		updatedExpenses, newID, err := addExpense(expenses, *description, *amount)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		if err := saveExpenses(filename, updatedExpenses); err != nil {
			fmt.Printf("Error saving expenses: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Expense added successfully (ID: %d)\n", newID)
	case "update":
		updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
		id := updateCmd.Int("id", 0, "ID of the expense to update")
		description := updateCmd.String("description", "", "New description")
		amount := updateCmd.Float64("amount", 0.0, "New amount")
		updateCmd.Parse(os.Args[2:])

		if *id <= 0 || *description == "" || *amount <= 0 {
			fmt.Println("Valid ID, description, and positive amount are required")
			os.Exit(1)
		}

		updatedExpenses, err := updateExpense(expenses, *id, *description, *amount)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		if err := saveExpenses(filename, updatedExpenses); err != nil {
			fmt.Printf("Error saving expenses: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Expense updated successfully")

	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deleteCmd.Int("id", 0, "ID of the expense to delete")
		deleteCmd.Parse(os.Args[2:])

		if *id <= 0 {
			fmt.Println("Valid ID is required")
			os.Exit(1)
		}

		updatedExpenses, err := deleteExpense(expenses, *id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		if err := saveExpenses(filename, updatedExpenses); err != nil {
			fmt.Printf("Error saving expenses: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Expense deleted successfully")

	case "list":
		printExpenses(expenses)

	case "summary":
		summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
		month := summaryCmd.Int("month", 0, "Month to summarize (1-12)")
		summaryCmd.Parse(os.Args[2:])

		total := calculateSummary(expenses, *month)
		if *month == 0 {
			fmt.Printf("Total expenses: $%.2f\n", total)
		} else {
			fmt.Printf("Total expenses for %s: $%.2f\n", time.Month(*month).String(), total)
		}

	default:
		printUsage()
		os.Exit(1)
	}
}

// printExpenses displays all expenses in a table format
func printExpenses(expenses []Expense) {
	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}
	fmt.Printf("%-5s %-10s %-20s %-10s\n", "ID", "Date", "Description", "Amount")
	for _, exp := range expenses {
		fmt.Printf("%-5d %-10s %-20s $%-10.2f\n", exp.ID, exp.Date.Format("2006-01-02"), exp.Description, exp.Amount)
	}
}

// printUsage displays the CLI usage instructions
func printUsage() {
	fmt.Println("Usage: expense-tracker <command> [arguments]")
	fmt.Println("Commands:")
	fmt.Println("  add --description <desc> --amount <amount>")
	fmt.Println("  update --id <id> --description <desc> --amount <amount>")
	fmt.Println("  delete --id <id>")
	fmt.Println("  list")
	fmt.Println("  summary [--month <month>]")
}
