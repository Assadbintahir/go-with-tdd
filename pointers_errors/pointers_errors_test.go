package pointers_errors

import (
	"fmt"
	"go-with-tdd/coverage"
	"testing"
)

func TestMain(t *testing.M) {
	coverage.EnforceCoverage(t, 1, "pointers_errors")
}

func TestWallet(t *testing.T) {
	t.Run("should deposit BTC to wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("should withdraw BTC from wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("should not draw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func TestBitcoin_String(t *testing.T) {
	got := Bitcoin(20)
	expected := "20 BTC"

	if got.String() != expected {
		t.Errorf("got %v but expected %v", got.String(), expected)
	}
}

func ExampleBitcoin_String() {
	got := Bitcoin(20)
	fmt.Println(got.String())
	// Output: 20 BTC
}

func assertBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != expected {
		t.Errorf("got %s but expected %s", got, expected)
	}
}

func assertError(t *testing.T, got error, expected error) {
	t.Helper()
	if got == nil {
		t.Fatal("Expected an error but there was none")
	}

	if got != expected {
		t.Errorf("got %v but expected %v", got.Error(), expected)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Got an error but didn't expected")
	}
}
