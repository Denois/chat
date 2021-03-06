package main
import (
  "github.com/Syfaro/telegram-bot-api"
  "log"
)

func main() {
  // подключаемся к боту с помощью токена
  bot, err := tgbotapi.NewBotAPI("87592177:AAE0HqJkmdhwRwWiz2SXk9N2sRzmR2MSRo0")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	err = bot.UpdatesChan(ucfg)
	// читаем обновления из канала
	for {
		select {
		case update := <-bot.Updates:
			// Пользователь, который написал боту
			UserName := update.Message.From.UserName

			// ID чата/диалога.
			// Может быть идентификатором как чата с пользователем
			// (тогда он равен UserID) так и публичного чата/канала
			ChatID := update.Message.Chat.ID

			// Текст сообщения
			Text := update.Message.Text

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

			// Ответим пользователю его же сообщением
			reply := Text
			// Созадаем сообщение
			msg := tgbotapi.NewMessage(ChatID, reply)
			// и отправляем его
			bot.SendMessage(msg)
		}

	}
}
