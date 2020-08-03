package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"messenger/db"
	"messenger/handlers"
	"messenger/service"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	mysqlDSN := os.Getenv("MYSQL_DSN")
	if mysqlDSN == "" {
		mysqlDSN = "root:123@/messenger"
	}

	dsn, err := mysql.ParseDSN(mysqlDSN)
	if err != nil {
		log.Fatalln(err)
	}

	// configure mysql
	dsn.ParseTime = true

	sqlDB, err := sqlx.Connect("mysql", dsn.FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	userDB := db.NewUser(sqlDB)
	chatDB := db.NewChat(sqlDB)
	userSrv := service.NewUserService(userDB)
	chatSrv := service.NewChatService(chatDB, userDB)
	messageSrv := service.NewMessageService(db.NewMessage(sqlDB), chatDB)

	http.Handle("/users/add", handlers.NewAddingUser(userSrv))

	http.Handle("/chats/add", handlers.NewAddingChat(chatSrv))

	http.Handle("/chats/get", handlers.NewGettingUserChats(chatSrv))

	http.Handle("/messages/add", handlers.NewAddingMessage(messageSrv))

	http.Handle("/messages/get", handlers.NewGettingChatMessages(messageSrv))

	log.Printf("starting listen on port: %s", port)
	fmt.Println(http.ListenAndServe(":"+port, nil))
}
