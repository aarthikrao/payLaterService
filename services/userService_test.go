package services

import "testing"

//Negative testcases
func TestCreateUserInvalidAmountN(t *testing.T) {
	actual := US.CreateUser([]string{"userTest", "test@test.com", "abc"})
	if actual {
		t.Errorf("User created despite invalid amount")
	}
}

func TestCreateUserIncorrectArgsN(t *testing.T) {
	actual := US.CreateUser([]string{"userTest", "test@test.com"})
	if actual {
		t.Errorf("User created despite invalid args")
	}
}

func TestCreateUserIncorrectPaybackN(t *testing.T) {
	actual := US.Payback([]string{"userTest", "0"})
	if actual {
		t.Errorf("User created despite invalid payback amount")
	}
}
