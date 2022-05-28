package yaml

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type YAMLConfig struct {
	DbURL      string `yaml:"DB_URL"`
	DbPort     int    `yaml:"DB_PORT"`
	DbUser     string `yaml:"DB_USER"`
	DbPassword string `yaml:"DB_PASSWORD"`
}

func ReadYAMLFile() {
	data, err := os.Open("yaml/data.yaml")
	if err != nil {
		panic(err)
	}

	defer func() {
		err := data.Close()
		if err != nil {
			fmt.Printf("Не могу закрыть файл %v", err)
		}
	}()
	config := YAMLConfig{}

	err = yaml.NewDecoder(data).Decode(&config)
	if err != nil {
		fmt.Printf("Не удалось декодировать структуру YAML %v", err)
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
