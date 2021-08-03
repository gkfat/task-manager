# Task Manager Backend
## A backend http server to manage your tasks

It's a small project for practicing Golang. Feel free to clone and give it a try, or report issues. Thanks for your comment.


## 1. Technology Stack
---
||Choice|
|-|-|
|Server|Golang|
|Database|Mariadb|
|API|RESTFUL|

## 2. Before Start
---
* Install Golang
* Install & Run Mariadb（default port 3306）
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
---
### /api/set/category `POST`
#### Request
##### Headers
Content-Type: `application/json`

##### Body
```javascript
{
  "Title": "Meeting"
}
```

#### Response
##### Body
```javascript
{
  "Code": 200,
  "Data": 1, // Category ID
  "Success": true
}
```

---
### /api/set/task `POST`
#### Request
##### Headers
Content-Type: `application/json`

##### Body
```javascript
{
  "Title":  "Task Manager 開發",    
  "Description": "API 撰寫",
  "CategoryID":  1,               // Use Category ID
  "LimitDate": "2021-05-23 04:12:29" // Time Format: YYYY-MM-dd HH:mm:ss
}
```

#### Response
##### Body
```javascript
{
    "Code": 200,
    "Data": 1, // Task ID
    "Success": true
}
```

---
### /api/fetch/category `GET`
#### Request
##### Headers
None
##### Body
None

#### Response
##### Body
```javascript
{
  "Code": 200,
  "Count": 1,
  "Data": [
    {
      "ID": 1,
      "Title": "Meeting"
    }
  ],
  "Success": true
}
```

---
### /api/fetch/task `POST`
#### Request
##### Headers
Content-Type: `application/json`

##### Body
```javascript
{
  "Page": 0,
  "Per": 1
}
```

#### Response
##### Body
```javascript
{
  "Code": 200,
  "Count": 1,
  "Data": [
    {
      "ID": 1,
      "Title": "Task Manager 開發",
      "Description": "API 撰寫",
      "Category": 1,
      "Month": 4,
      "Week": 2,
      "Weekday": 4,
      "WorkingHour": 1.5
    }
  ],
  "Page": 0,
  "Per": 1,
  "Success": true
}
```
