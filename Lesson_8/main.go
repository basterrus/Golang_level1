//Разработайте пакет для чтения конфигурации типичного веб-приложения через флаги или переменные окружения.
//Пример конфигурации можно посмотреть здесь. По желанию вы можете задать другие имена полям, сгруппировать их или добавить собственные поля.
//Помимо чтения конфигурации приложение также должно валидировать её - например, все URL’ы должны соответствовать ожидаемым форматам.
//Работу с конфигурацией необходимо вынести в отдельный пакет (не в пакет main).

package main

import (
	"abc/config"
	"fmt"
	"github.com/joho/godotenv"
)

// Функция устанавливает переменное окружение считывая его из .env
// используя библиотеку "godotenv"
func init() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Конфигурационный файл не найден!")
	}
	fmt.Println("Конфигурационный файл успешно загружен!")
}

// В main вызываем импортированную функцию ReadConfig из пакета config
func main() {
	fmt.Println("Чтение данных из виртуального окружения: ")
	configurations := config.ReadConfig()
	fmt.Printf("DB_URL: %s\nDB_PORT: %d\nDB_USER: %s\nDB_PASSWORD: %s\n",
		configurations.DbUrl,
		configurations.DbPort,
		configurations.DbUser,
		configurations.DbPassword)

}
