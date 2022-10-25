
<h1 align="center">
	:books: ToDo-App
</h1>

<h2 align="center">
	📝 Description
</h2>

> _This is a simple RESTful API to help users manage their daily routine! Written in Golang using Postgres database for data storage._

<h2 align="center">
	🛠️ Usage
</h2>

1. First of all clone this project on your machine:
```shell
git clone git@github.com:MKKurbandibirov/todo-app.git
```
2. Create a `.env` file in which you must insert the password to log in to your database (the same as in docker-compose.yml) like:

```shell
DB_PASSWORD=<your_password>
```
3. Final, run it:
```shell
go run cmd/main.go
```
or
```shell
go build cmd/main.go && ./main
```
