package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type SafeConn struct {
	conn sqlx.DB;
	err error;
}

func (conn *SafeConn) checkAllIsOK() {
	if conn.err != nil { panic(conn.err) };
}

func _main() {
	conn, err := sqlx.Connect("mysql", "root:dummy@tcp(localhost:3306)/story")
	if err != nil {
		panic(err)
	}
	res, err := conn.Exec("INSERT INTO users (name) VALUES(\"Peter\")")
	if err != nil { panic(err) }
	id, err := res.LastInsertId()
	if err != nil { panic(err) }
	fmt.Printf("Created user with id:%d", id)
	var user User
	err = conn.Get(&user, "select * from users where id=?", id)
	if err != nil { panic(err) }
	_, err = conn.Exec("UPDATE users set name=\"John\" where id=?", id)
	if err != nil { panic(err) }
	_, err = conn.Exec("DELETE FROM users where id=?", id)
	if err != nil { panic(err) }

}


func main() {
	key := os.Getenv("TELEGRAM_BOT_KEY")
	if key == "" { panic("No telegram key present")}
	bot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
