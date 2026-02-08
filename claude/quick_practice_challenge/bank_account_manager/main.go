package main

import (
	"bufio"
	"fmt"
	"os"
)

// Challenge: Create a bank account system using pointers:
// Create a struct:
// type BankAccount struct {
//     accountNumber string
//     holderName    string
//     balance       float64
// }
// Functions (all use pointers):

// deposit(account *BankAccount, amount float64)
// withdraw(account *BankAccount, amount float64) error - check sufficient funds
// transfer(from, to *BankAccount, amount float64) error
// displayAccount(account *BankAccount)

// Menu:

// Create account
// Deposit
// Withdraw
// Transfer between accounts
// Display account
// Exit

// Bonus: Maintain multiple accounts using a slice/map of pointers!

type BankAccount struct {
	accountNumber string
	holderName    string
	balance       float64
}

func getAmount() float64 {
	var amount float64
	fmt.Print("Enter amount: ")
	fmt.Scan(&amount)
	return amount
}

func getAccountNumber(reader *bufio.Reader) string {
	var accountNumber string
	fmt.Print("Enter account number: ")
	accountNumber, _ = reader.ReadString('\n')
	return accountNumber
}

func getHolderName(reader *bufio.Reader) string {
	var holderName string
	fmt.Print("Enter holder name: ")
	holderName, _ = reader.ReadString('\n')
	return holderName
}

func createAccount(reader *bufio.Reader) *BankAccount {
	accountNumber := getAccountNumber(reader)
	holderName := getHolderName(reader)
	balance := getAmount()

	return &BankAccount{
		accountNumber: accountNumber,
		holderName:    holderName,
		balance:       balance,
	}
}

func deposit(account *BankAccount) {
	amount := getAmount()
	account.balance += amount
}

func withdraw(account *BankAccount) error {
	amount := getAmount()
	if account.balance < amount {
		return fmt.Errorf("insufficient funds")
	}
	account.balance -= amount
	return nil
}

func transfer(from, to *BankAccount) error {
	amount := getAmount()
	if from.balance < amount {
		return fmt.Errorf("insufficient funds")
	}
	from.balance -= amount
	to.balance += amount
	return nil
}

func displayAccount(account *BankAccount) {
	fmt.Println("Account Number:", account.accountNumber)
	fmt.Println("Holder Name:", account.holderName)
	fmt.Println("Balance:", account.balance)
}

func main() {
	accounts := make(map[string]*BankAccount)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("1. Create account")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Transfer")
		fmt.Println("5. Display account")
		fmt.Println("6. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			account := createAccount(reader)
			accounts[account.accountNumber] = account
		case 2:
			accountNumber := getAccountNumber(reader)
			if account, ok := accounts[accountNumber]; ok {
				deposit(account)
			} else {
				fmt.Println("Account not found")
			}
		case 3:
			accountNumber := getAccountNumber(reader)
			if account, ok := accounts[accountNumber]; ok {
				err := withdraw(account)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Account not found")
			}
		case 4:
			from := getAccountNumber(reader)
			to := getAccountNumber(reader)
			if fromAccount, ok := accounts[from]; ok {
				if toAccount, ok := accounts[to]; ok {
					err := transfer(fromAccount, toAccount)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println("To account not found")
				}
			} else {
				fmt.Println("From account not found")
			}
		case 5:
			accountNumber := getAccountNumber(reader)
			if account, ok := accounts[accountNumber]; ok {
				displayAccount(account)
			} else {
				fmt.Println("Account not found")
			}
		case 6:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
