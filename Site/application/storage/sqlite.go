package storage

import (
	"fmt"
	"strings"
	"time"
)

func (storage *SQLite) Close()  {
	storage.DB.Close()
}

// Пользователь
func (storage *SQLite) AuthorUser(email, password string) bool {
	u := User{}
	row := storage.DB.QueryRow("select * from Users where email = $1", email)
	row.Scan(&u.ID, &u.Email, &u.Password, &u.DateOfRegistory, &u.DurationOfThePrivilege, &u.Role)

	if u.Password == GetMd5(password) {
		return true
	} else {
		return false
	}
}

func (storage *SQLite) UpdateUser(email, password string) bool {
	pass := GetMd5(password)

	_, err := storage.DB.Exec("update Users set password = $1 where email = $2",
		pass, email)
	if err != nil {
		return false
	}

	return true
}

func (storage *SQLite) InsertUser(user User) bool {
	pass := GetMd5(user.Password)
	_, err := storage.DB.Exec("insert into Users (email, password, date_of_registory, duration_of_the_privilege, role) values ($1, $2, $3, $4, $5)",
		user.Email, pass, user.DateOfRegistory, user.DurationOfThePrivilege, user.Role)
	if err != nil {
		return false
	}

	return true
}

func (storage *SQLite) DeleteUser(email string) bool {
	_, err := storage.DB.Exec("delete from user where email = $1", email)
	if err != nil {
		return false
	}

	return true
}

func (storage *SQLite) GetUser(email string) (User, bool) {
	u := User{}
	row := storage.DB.QueryRow("select * from Users where email = $1", email)

	row.Scan(&u.ID, &u.Email, &u.Password, &u.DateOfRegistory, &u.DurationOfThePrivilege, &u.Role)

	if u.Email == "" {
		return u, false
	} else {
		date := strings.Split(u.DurationOfThePrivilege, "/")
		resDate := ""

		if len(date) == 3 {
			if len(date[2]) == 1 {
				resDate = resDate + "0" + date[2] + "."
			} else {
				resDate = resDate + date[2] + "."
			}

			if len(date[1]) == 1 {
				resDate = resDate + "0" + date[1] + "."
			} else {
				resDate = resDate + date[1] + "."
			}
		}

		u.DurationOfThePrivilege =  resDate + date[0]

		return u, true
	}
}

// Заявки
func (storage *SQLite) InsertRequest(req Request) (bool, int64) {
	pos, err := storage.DB.Exec("insert into Requests (email, date_of_registory, value_day) values ($1, $2, $3)",
		req.Email, req.DateOfRegistory, req.ValueDay)
	if err != nil {
		return false, -1
	}
	id, _ := pos.LastInsertId()

	return true, id
}

func (storage *SQLite) GetRequest(id int64) (Request, bool) {
	req := Request{}
	row := storage.DB.QueryRow("select * from Requests where id = $1", id)

	row.Scan(&req.ID, &req.Email, &req.DateOfRegistory, &req.Status, &req.ValueDay)

	if req.Email == "" {
		return req, false
	} else {
		return req, true
	}
}

func (storage *SQLite) GetRequests(email string) []Request {
	rows, err := storage.DB.Query("select * from Requests where email = $1", email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	requests := []Request{}

	for rows.Next(){
		p := Request{}
		err := rows.Scan(&p.ID, &p.Email, &p.DateOfRegistory, &p.Status, &p.ValueDay)
		if err != nil{
			continue
		}

		requests = append(requests, p)
	}

	return requests
}

func (storage *SQLite) GetAllRequests() []Request {
	type req struct{
		ID						uint32
		Email					string
		DateOfRegistory			string
		Status					int
		ValueDay				uint8
	}
	rows, err := storage.DB.Query("select * from Requests where status = $1", 0)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	requests := []Request{}

	for rows.Next(){
		p := req{}
		err := rows.Scan(&p.ID, &p.Email, &p.DateOfRegistory, &p.Status, &p.ValueDay)
		if err != nil{
			continue
		}

		r := Request{
			ID: p.ID,
			Email: p.Email,
			DateOfRegistory: p.DateOfRegistory,
			Status: p.Status == 1,
			ValueDay: p.ValueDay,
		}
		requests = append(requests, r)
	}

	return requests
}

// Заявки для апи
func (storage *SQLite) AcceptRequest(id int64) bool {
	req := Request{}
	u := User{}
	var durationOfThePrivilege string
	status := 1

	_, err := storage.DB.Exec("update Requests set status = $1 where id = $2",
		status, id)
	if err != nil {
		return false
	}
	row := storage.DB.QueryRow("select * from Requests where id = $1", id)

	row.Scan(&req.ID, &req.Email, &req.DateOfRegistory, &req.Status, &req.ValueDay)

	row = storage.DB.QueryRow("select * from Users where email = $1", req.Email)

	row.Scan(&u.ID, &u.Email, &u.Password, &u.DateOfRegistory, &u.DurationOfThePrivilege, &u.Role)
	if u.DurationOfThePrivilege != "0000/0/0" {
		date, _ := time.Parse("2006/1/2", u.DurationOfThePrivilege)
		date = date.AddDate(0, 0, int(req.ValueDay))
		durationOfThePrivilege = date.Format("2006/1/2")
	} else {
		date := time.Now()
		date = date.AddDate(0, 0, int(req.ValueDay))
		durationOfThePrivilege = date.Format("2006/1/2")
	}
	_, err = storage.DB.Exec("update Users set duration_of_the_privilege = $1 where email = $2",
		durationOfThePrivilege, req.Email)
	if err != nil {
		return false
	}

	return true
}

func (storage *SQLite) CancelRequest(id int64) bool {
	status := "Отклонён"
	_, err := storage.DB.Exec("update Requests set status = $1 where id = $2",
		status, id)
	if err != nil {
		return false
	}

	return true
}

func (storage *SQLite) InsertFound(fnd Found) bool {
	_, err := storage.DB.Exec("insert into found (email, kor, name, status) values ($1, $2, $3, $4)",
		fnd.Email, fnd.Kor, fnd.Name, fnd.Status)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func (storage *SQLite) GetFound() (Found, bool) {
	fnd := Found{}
	row := storage.DB.QueryRow("select * from found where status = $1", "В очереди")

	row.Scan(&fnd.ID, &fnd.Email, &fnd.Kor, &fnd.Name, &fnd.Status)

	if fnd.Email == "" {
		return fnd, false
	} else {
		_, err := storage.DB.Exec("update found set status = $1 where id = $2",
			"Обработан", fnd.ID)
		if err != nil {
			return fnd, false
		}
		return fnd, true
	}
}

