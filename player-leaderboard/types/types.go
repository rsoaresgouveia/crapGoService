package types

import "time"

type User struct {
	Username string
	Password string `json:"-"`
	IsAdmin bool
	CreatedAt time.Time
}

type Truck struct{
	Melancia string
	Gatinha string
	Picole string
}

type leaderboard struct{
	id int
	player string
	leaderboard int64
}
