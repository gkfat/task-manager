package api

type User struct {
	ID       uint
	UserName string
	Account  string
}

type Category struct {
	ID    uint
	Title string
}

type Task struct {
	ID          uint
	Title       string
	Description string
	Category    uint
	Month       uint
	Week        uint
	Weekday     uint
	WorkingHour float64
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
