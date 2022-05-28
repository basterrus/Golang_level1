//Разработайте пакет для чтения конфигурации типичного веб-приложения через флаги или переменные окружения.
//Пример конфигурации можно посмотреть здесь. По желанию вы можете задать другие имена полям, сгруппировать их или добавить собственные поля.
//Помимо чтения конфигурации приложение также должно валидировать её - например, все URL’ы должны соответствовать ожидаемым форматам.
//Работу с конфигурацией необходимо вынести в отдельный пакет (не в пакет main).

package main

import (
	"abc/config"
	"abc/json"
	"abc/yaml"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Функция устанавливает переменное окружение считывая его из .env используя библиотеку "godotenv"
func init() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Конфигурационный файл не найден!")
	}
}

func main() {

	var decodeMethods string
	fmt.Println("Выбирете метод загрузки конфигурации \n" +
		"Введите '1' для загрузки конфигурации из виртуального окружения\n" +
		"Введите '2' для загрузки конфигурации из JSON файла\n" +
		"Введите '3' для загрузки конфигурации из YAML файла\n" +
		"Введите 'exit' для выхода из приложения")
	fmt.Scan(&decodeMethods)

	switch decodeMethods {
	case "1":
		fmt.Println("Чтение данных из виртуального окружения: ")
		config.ReadConfig()
	case "2":
		fmt.Println("Чтение данных из JSON файла: ")
		json.ReadJSONFile()
	case "3":
		fmt.Println("Чтение данных из YAML файла: ")
		yaml.ReadYAMLFile()
	case "exit":
		os.Exit(1)
	default:
		fmt.Println("Вы выбрали не корректную операцию!")
	}

}