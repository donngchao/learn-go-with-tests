package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()

	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

	wallet2 := Wallet{balance:30}
	wallet2.Deposit(Bitcoin(20))
	got2 := wallet2.Balance()
	want2 := Bitcoin(50)

	if got2 != want2 {
		t.Errorf("got %v want %v",got2,want2)
	}

}
