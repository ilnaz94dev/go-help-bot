package endpoint

import (
	tlg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
	"go-help-bot/internal/app/commands"
	"go-help-bot/internal/app/common"
	"os"
	"strconv"
	"strings"
)

var bot *tlg.BotAPI
var lockCommands map[string]bool

func StartTlgBot(c echo.Context) error {
	key := os.Getenv("TELEGRAM_API_KEY")
	var err error
	bot, err = tlg.NewBotAPI(key)
	if err != nil {
		return err
	}
	bot.Debug = true
	lockCommands = make(map[string]bool)

	updateConfig := tlg.NewUpdate(0)
	updateConfig.Timeout = 60
	updates := bot.GetUpdatesChan(updateConfig)

	allowChats := getAllowChatsMap()

	for update := range updates {
		inMsg := update.Message
		if inMsg == nil || !allowChats[inMsg.Chat.ID] {
			continue
		}
		var outText string
		if inMsg.IsCommand() {
			params := strings.Split(inMsg.Text, " ")
			switch inMsg.Command() {
			case "ping":
				if isLockCommand("check_site") {
					outText = "Site check already started"
					break
				}
				// TODO stop check
				// TODO service to struct
				pingParams, err := commands.GetPingParamsByText(inMsg.Text)
				if err != nil {
					outText = err.Error()
					break
				}
				outText = "Site check started"
				lockCommand("check_site")

				go pingParams.PingSite(func(isAvailable bool) {
					if isAvailable {
						sendTlgMsg(inMsg.Chat.ID, "Site available! "+pingParams.Url)
					}
					unlockCommand("check_site")
				})
			case "vpn":
				filePath := os.Getenv("VPN_FILE_PATH")
				if params[1] == "add" {
					err = commands.AddTextToFile(params[2], filePath)
				} else if params[1] == "remove" {
					err = commands.RemoveTextFromFile(params[2], filePath)
				}
			case "currency":
				outText, err = commands.GetCurrency(params[1])
			case "ticker":
				outText, err = commands.GetTicker(params[1])
			}
		} else {
			outText = inMsg.Text + " lol"
		}
		if err != nil {
			common.HandleError("Command execution error", err)
		}
		sendTlgMsg(inMsg.Chat.ID, outText)
	}
	return nil
}

func isLockCommand(command string) bool {
	return lockCommands[command]
}

func lockCommand(command string) {
	lockCommands[command] = true
}

func unlockCommand(command string) {
	delete(lockCommands, command)
}

func sendTlgMsg(chatId int64, text string) {
	msg := tlg.NewMessage(chatId, text)
	if _, err := bot.Send(msg); err != nil {
		common.DebugMsg(err.Error())
	}
}

func getAllowChatsMap() map[int64]bool {
	chats := strings.Split(os.Getenv("TELEGRAM_ALLOW_CHATS"), ";")
	chatsMap := make(map[int64]bool)
	for _, chatId := range chats {
		k, err := strconv.ParseInt(chatId, 10, 64)
		if err == nil {
			chatsMap[k] = true
		}
	}
	return chatsMap
}
