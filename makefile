run:
	go run main.go

table: 
	sqlite3  ./database/user.sqlite < ./script/create-user-table.sql