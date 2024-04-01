package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				// Kalau ada variabel yg di sharing/diakses oleh beberapa goroutine
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter =", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	// Lock & Unlock untuk write
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	// Rlock & RUnlock untuk read
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
                account.AddBalance(1)
				fmt.Println(account.GetBalance())
            }
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance =", account.GetBalance())
}

// ----------------------------------------------------------------

type UserBalance struct {
	sync.Mutex
	Name string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
    user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

/* Kode mengalami deadlock
func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)
	
	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}
*/

// kode sudah tidak deadlock
func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	// Mengunci dalam urutan yang ditentukan
	if user1.Name < user2.Name {
		user1.Lock()
		user2.Lock()
	} else {
		user2.Lock()
		user1.Lock()
	}

	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name: "Fadli", 
		Balance: 1000000,
	}

    user2 := UserBalance{
		Name: "Darusalam", 
		Balance: 1000000,
	}

    go Transfer(&user1, &user2, 100000)
    go Transfer(&user2, &user1, 200000)

    time.Sleep(10 * time.Second)

	fmt.Println("User", user1.Name, "Balance :", user1.Balance)
	fmt.Println("User", user2.Name, "Balance :", user2.Balance)

	// Terjadi deadlock dimana baris 113 me-lock user 1, kemudian seharusnya berlanjut me-lock user 2, namun user 2 di-lock baris 114 lebih dahulu. Begitupula sebaliknya
}