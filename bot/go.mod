module github.com/SpenserCai/OnceHumanTools/bot

go 1.21

require (
	github.com/SpenserCai/OnceHumanTools/backend v0.0.0
	github.com/bwmarrin/discordgo v0.27.1
	github.com/joho/godotenv v1.5.1
)

replace github.com/SpenserCai/OnceHumanTools/backend => ../backend

require (
	github.com/gorilla/websocket v1.5.1 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
)
