package models

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestPasswordEncryption(t *testing.T) {

	test := User{0, "FooBar", "F001ngwhoS|ag0|@NG24/7"}

	if err := test.EncryptData(); err != nil {
		t.Fatalf(`encryptData() = %q, want "", error`, err)
	}
}

func TestPasswordValidation(t *testing.T) {

	test := User{0, "FizzBuzz", "g0sS1pingAbtg0|@NG24/7"}

	test.EncryptData()
	bcrypt.GenerateFromPassword([]byte("g0sS1pingAbtg0|@NG24/7"), bcrypt.MinCost)

	if err := test.ValidatePassword("g0sS1pingAbtg0|@NG24/7"); err != nil {
		t.Fatalf(`validatePassword() = %q, want "", error`, err)
	}

}
