module github.com/SpenserCai/OnceHumanTools

go 1.21

require (
	github.com/SpenserCai/OnceHumanTools/backend v0.0.0
	github.com/SpenserCai/OnceHumanTools/bot v0.0.0
)

replace (
	github.com/SpenserCai/OnceHumanTools/backend => ./backend
	github.com/SpenserCai/OnceHumanTools/bot => ./bot
)