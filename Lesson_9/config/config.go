package config

import (
	"fmt"
	"os"
	"strconv"
)

//
//type Configuration struct {
//	DbUrl      string
//	DbPort     int64
//	DbUser     string
//	DbPassword string
//}

func ReadConfig() {
	var envUrl string
	var envPort int64
	var envUser string
	var envPassword string

	envUrl = os.Getenv("DB_URL")
	//Валидация URL

	envPort, _ = strconv.ParseInt(os.Getenv("DB_PORT"), 0, 64)
	//Валидация порта
	if 1023 <= envPort && envPort <= 65536 {
		fmt.Println("Порт введен верно!")
	} else {
		fmt.Println("Порт введен не верно!")
	}

	envUser = os.Getenv("DB_USER")
	if envUser == "" {
		fmt.Println("Введите имя пользователя!")
	}
	fmt.Println("Имя пользователя считано успешно!")

	envPassword = os.Getenv("DB_PASSWORD")
	if envPassword == "" {
		fmt.Println("Введите пароль!")
	}
	fmt.Println("Пароль считан успешно!")
	fmt.Printf("DB_URL: %s\nDB_PORT: %d\nDB_USER: %s\nDB_PASSWORD: %s\n",
		envUrl,
		envPort,
		envUser,
		envPassword)

}
