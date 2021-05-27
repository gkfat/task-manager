package api

import "time"

type User struct {
	ID       uint
	UserName string
	Account  string
}

type Category struct {
	ID    uint
	Title string `binding:"required"`
}

type Task struct {
	ID          uint
	CreatedAt   time.Time
	Title       string `binding:"required"`
	Description string `binding:"required"`
	CategoryID  uint   `binding:"required"`
	LimitDate   string `binding:"required"`
}

type TaskQuery struct {
	ID          uint
	Title       string
	Description string
	Category    uint
	LimitDate   string
	Page        int
	Per         int
}
