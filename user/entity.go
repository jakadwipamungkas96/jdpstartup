package user

import "time"

type User struct {
	Id               int
	Name             string
	Occupation       string
	Email            string
	Password_hash    string
	Avatar_file_name string
	Role             string
	Created_at       time.Time
	Updated_at       time.Time
}
