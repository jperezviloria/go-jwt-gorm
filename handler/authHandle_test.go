package handler

import (
	"fmt"
	"gorm.io/gorm/utils/tests"
	"testing"
)

func TestGetUserByUsername(t *testing.T) {
	const username = "msesarego"
	user, _ := GetUserByUsername(username)
	fmt.Println(user)
	fmt.Println(user.Username)
	fmt.Println(user.Password)
	tests.AssertEqual(t, user.Username, username)

}

func TestIntegral(t *testing.T) {
	const password = "julio123"
	const password2 = "julio123"
	token, _ := HashPassword(password)
	compare := CheckPasswordHash(password2, string(token))
	tests.AssertEqual(t, compare, true)
}

func TestIntegral2(t *testing.T) {
	const password = "julio123"
	const password2 = "julio123"
	const username = "msesarego"
	user, _ := GetUserByUsername(username)
	//token, _ := HashPassword(password)
	compare := CheckPasswordHash(password2, user.Password)
	tests.AssertEqual(t, compare, true)
}
