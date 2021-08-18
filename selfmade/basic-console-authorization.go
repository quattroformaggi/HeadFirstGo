/*Система авторизации логин-пароль.
WILDCARD: admin-admin
запись в файл? смена пароля? зарегистрировать новый аккаунт?*/
package main

import (
	"fmt"
	"log"
)

type Account struct { //единичный набор логин-пароль
	Login    string
	Password string
}

type Accounts struct { //список всех существующих аккаунтов
	Combinations []Account
}

func check(err error) { //проверка ошибок
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db := Accounts{}
	wildcard := Account{"admin", "admin"}
	db.Combinations = append(db.Combinations, wildcard)
	var log string
	var pass string
	fmt.Print("\nЛогин:\t")
	fmt.Scan(&log)
	fmt.Print("\nПароль:\t")
	fmt.Scan(&pass)
	tmp := Account{log, pass}
	if tmp.Login == wildcard.Login && tmp.Password == wildcard.Password {
		fmt.Println("\nВход успешно выполнен,", log)
	} else if tmp.Password != wildcard.Password && tmp.Login == wildcard.Login {
		fmt.Println("\nНеверный пароль,", log)
	} else if tmp.Login != wildcard.Login && tmp.Password == wildcard.Password {
		fmt.Println("\nНеверный логин.")
	} else {
		fmt.Println("\nНеверные данные.")
	}
	fmt.Println("Существующие аккаунты:", db)
}
