module github.com/oncehuman/tools

go 1.21

require (
	github.com/oncehuman/tools/backend v0.0.0
	github.com/oncehuman/tools/bot v0.0.0
)

replace (
	github.com/oncehuman/tools/backend => ./backend
	github.com/oncehuman/tools/bot => ./bot
)