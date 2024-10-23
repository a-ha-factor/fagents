package tg

import (
	"context"
	"fagents/db"
	"fagents/types"
	"os"
	"os/signal"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// Send any text message to the bot after the bot has been started
var FagentsBotToken string

func InitBot() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(searchDialogHandler),
	}

	b, err := bot.New(FagentsBotToken, opts...)
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic("Bot token is incorrect")
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, startHelpHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/help", bot.MatchTypeExact, startHelpHandler)

	b.Start(ctx)
}

func startHelpHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Привет!\nЯ бот для поиска иноагентов в реестре Минюста РФ согласно 255-ФЗ от 14 июля 2022 г.\nДля начала поиска пришли мне сообщение длиной от 4 до 30 символов. Я ищу по ФИО/названию, ИНН, участникам",
	})
}
func searchDialogHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	var (
		faList types.FagentsList
		err    error
		fagent types.Fagent
		msg    string
	)
	if update.Message == nil {
		return
	}

	faList, err = db.FagentsList(update.Message.Text)
	if err == nil && len(faList) > 30 {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Результатов больше 30. Уточните критерий поиска",
		})
		return
	}

	if err != nil || len(faList) == 0 {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Ничего не найдено",
		})
	} else {
		for _, fagent = range faList {
			msg = "<b>" + fagent.FullName + "</b>\n" +
				"<b>№ п/п:</b> " + fagent.Id + "\n" +
				"<b>Дата рождения:</b> " + fagent.Dob + "\n" +
				"<b>ОГРН:</b> " + fagent.Ogrn + "\n" +
				"<b>ИНН:</b> " + fagent.Inn + "\n" +
				"<b>Рег. номер:</b> " + fagent.RegNum + "\n" +
				"<b>СНИЛС:</b> " + fagent.Snils + "\n" +
				"<b>Адрес:</b> " + fagent.Address + "\n" +
				"<b>Инф. ресурсы:</b> " + strings.ReplaceAll(fagent.Resources, ";", "\n") + "\n" +
				"<b>Участники:</b> " + strings.ReplaceAll(strings.ReplaceAll(fagent.Members, ",", ";"), ";", "\n") + "\n" +
				"<b>Основание включения:</b> " + fagent.Law + "\n" +
				"<b>Дата включения:</b> " + fagent.DateIn + "\n" +
				"<b>Дата опубликования:</b> " + fagent.DatePubl + "\n" +
				"<b>Дата исключения:</b> " + fagent.DateOut

			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:    update.Message.Chat.ID,
				Text:      msg,
				ParseMode: models.ParseModeHTML,
			})
		}
	}
}
