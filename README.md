# Expense Tracker CLI

A command-line application to manage your expenses, built with Go. Expenses are stored in a JSON file (`expenses.json`) and can be added, updated, deleted, viewed, and summarized by total or monthly spending.

## Features
- Add a new expense with a description and amount
- Update an expense's description and amount
- Delete an expense by ID
- List all expenses in a tabular format
- View a summary of total expenses
- View a summary of expenses for a specific month in the current year
- Persistent storage in a JSON file
- Error handling for invalid inputs (e.g., negative amounts, non-existent IDs)

## Requirements
- Go 1.21 or later

## Installation
1. Clone or download this repository.
2. Navigate to the project directory:
```
cd expense-tracker
```

3. Initialize Go modules if not already done:
```
go mod init expense-tracker
```

4. Build the application:
```
go build -o expense-tracker
```

## Usage

Run the application using the compiled binary (e.g., `./expense-tracker`).

### Commands

| Command                                   | Description                              | Example                                    |
|-------------------------------------------|------------------------------------------|--------------------------------------------|
| `add --description <desc> --amount <amt>` | Add a new expense                        | `./expense-tracker add --description "Lunch" --amount 20` |
| `update --id <id> --description <desc> --amount <amt>` | Update an expense's details              | `./expense-tracker update --id 1 --description "Dinner" --amount 25` |
| `delete --id <id>`                        | Delete an expense                        | `./expense-tracker delete --id 1`          |
| `list`                                    | List all expenses                        | `./expense-tracker list`                   |
| `summary`                                 | Show total expenses                      | `./expense-tracker summary`                |
| `summary --month <month>`                 | Show expenses for a specific month       | `./expense-tracker summary --month 8`      |

### Example Output
```
$ ./expense-tracker add --description "Lunch" --amount 20
Expense added successfully (ID: 1)

$ ./expense-tracker add --description "Dinner" --amount 10
Expense added successfully (ID: 2)

$ ./expense-tracker list
ID    Date       Description           Amount
1     2025-04-14 Lunch                 $20.00
2     2025-04-14 Dinner                $10.00

$ ./expense-tracker summary
Total expenses: $30.00

$ ./expense-tracker summary --month 4
Total expenses for April: $30.00

$ ./expense-tracker delete --id 2
Expense deleted successfully

$ ./expense-tracker summary
Total expenses: $20.00
```
