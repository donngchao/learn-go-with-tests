package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw test again",func(t *testing.T){
		wallet := Wallet{balance:Bitcoin(100)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t,wallet,Bitcoin(90))
	})

	t.Run("Deposit  test again",func(t *testing.T){
		wallet := Wallet{balance:Bitcoin(100)}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t,wallet,Bitcoin(110))
	})


}
