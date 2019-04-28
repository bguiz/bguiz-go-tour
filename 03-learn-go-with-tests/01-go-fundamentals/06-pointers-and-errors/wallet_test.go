package wallet

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, w *Wallet, want Ether) {
		t.Helper()
		got := w.Balance()

		if got != want {
			// t.Errorf("got %.2f want %.2f", got, want)
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, err error, want error) {
		t.Helper()

		if err == nil {
			t.Fatal("expected error, but did not get one")
		}

		if err != want {
			t.Errorf("want '%s' got '%s'", want, err)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Ether(10.0))

		assertBalance(t, &wallet, Ether(10.0))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 43.21}
		err := wallet.Withdraw(Ether(10.21))

		assertBalance(t, &wallet, Ether(33.00))
		if err != nil {
			t.Errorf("expected no error, but got %s", err)
		}
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: 12.34}
		err := wallet.Withdraw(Ether(123.45))

		assertBalance(t, &wallet, Ether(12.34))
		assertError(t, err, ErrInsufficientFunds)
	})
}

// NOTE
// - switch from  `%.2f` to `%s`, because we have implemented the `Stringer`
//   interface for the `Ether` type
// - the error `err := wallet.Withdraw(Ether(123.45))` shows up as
//   `wallet.Withdraw(Ether(123.45)) used as value` is oddly phrased,
//   but is basically means that there was no return value defined for this
//   function
// - calling `t.Fatal()` stops the rest of the test from running immediately
