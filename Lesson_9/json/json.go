package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type JSONConfig struct {
	DbURL      string `json:"DB_URL"`
	DbPort     int    `json:"DB_PORT"`
	DbUser     string `json:"DB_USER"`
	DbPassword string `json:"DB_PASSWORD"`
}

func ReadJSONFile() {
	data, err := os.Open("json/data.json")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := data.Close()
		if err != nil {
			fmt.Printf("Не могу закрыть файл %v", err)
		}
	}()

	config := JSONConfig{}

	err = json.NewDecoder(data).Decode(&config)
	if err != nil {
		fmt.Printf("Не удалось декодировать структуру JSON %v", err)
	}
	// Проверка порта
	if 1023 <= config.DbPort && config.DbPort <= 65536 {
		fmt.Println("Порт введен верно!")
	} else {
		fmt.Println("Порт введен не верно!")
	}

	// Проверка не пустого пользователя
	if config.DbUser == "" {
		panic("Имя пользователя не задано!")
	}
	fmt.Println("Имя пользователя считано успешно!")

	// Проверка пустого пароля
	if config.DbPassword == "" {
		fmt.Println("Введите пароль!")
	}
	fmt.Println("Пароль считан успешно!")

	fmt.Printf("DB_URL: %s,\nDB_PORT: %d,\nDB_USER: %s,\nDB_PASSWORD: %s\n",
		config.DbURL,
		config.DbPort,
		config.DbUser,
		config.DbPassword)
}
