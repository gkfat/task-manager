# Task Manager Backend
## A backend http server to manage your tasks

It's a small project for practicing Golang. Feel free to clone and give it a try, or report issues. Thanks for your comment.


## 1. Technology Stack
||Choice|
|-|-|
|Server|Golang|
|Database|Mariadb|
|API|RESTFUL|

## 2. Before Start
* Install Golang
* Install Mariadb（default port 3306）
* Create root user by following command
```sql
MariaDB> CREATE USER 'root'@'localhost' IDENTIFIED BY '123456';
```
* Create a database by following command
```sql
MariaDB> CREATE DATABASE task_manager;
```
* And you're ready to start!
```go
$ go run main.go
$ Task Manager Server is running on http://localhost:7777
```

## 3. API Document

### Set Category

|Title|Description|
|-|-|
|API route|`/set/category`|
|Method|`POST`|
Request Header
Request Param
Response Http Code
Response Body





* `/set/task`


* `/fetch/category`


* `/set/task`