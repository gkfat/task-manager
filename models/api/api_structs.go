package api

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
	Title       string  `binding:"required"`
	Description string  `binding:"required"`
	CategoryID  uint    `binding:"required"`
	Month       uint    `binding:"required"`
	Week        uint    `binding:"required"`
	Weekday     uint    `binding:"required"`
	WorkingHour float64 `binding:"required"`
}

type TaskQuery struct {
	ID          uint
	Title       string
	Description string
	Category    uint
	Month       uint
	Week        uint
	Weekday     uint
	WorkingHour float64
	Page        int
	Per         int
}
