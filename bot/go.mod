module github.com/oncehuman/tools/bot

go 1.21

require (
	github.com/bwmarrin/discordgo v0.27.1
	github.com/joho/godotenv v1.5.1
	github.com/oncehuman/tools v0.0.0
)

replace github.com/oncehuman/tools => ../backend

require (
	github.com/gorilla/websocket v1.5.1 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
)