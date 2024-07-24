package main

import (
	"s3_file_uploader/commands"
	"s3_file_uploader/config"
	"s3_file_uploader/utils"
)

func main() {
	// Загрузка конфигурации из .env файла
	config.LoadConfig()

	// Инициализация логирования
	utils.InitLogger(config.LogFile)

	// Выполнение команд
	if err := commands.Execute(); err != nil {
		utils.LogError("MongoLogError executing command: ", err)
	}
}
