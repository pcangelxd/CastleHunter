package storage

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
)

type Storage interface {
	// Закрытие соединения с хранилищем
	Close()

	// Пользователь
	AuthorUser(email, password string) bool
	InsertUser(user User) bool
	DeleteUser(email string) bool
	GetUser(email string) (User, bool)
	UpdateUser(email, password string) bool

	// Заявки
	InsertRequest(req Request) (bool, int64)
	GetRequest(id int64) (Request, bool)
	GetRequests(email string) []Request
	GetAllRequests() []Request

	// Заявки для апи
	AcceptRequest(id int64)	bool
	CancelRequest(id int64) bool

	InsertFound(fnd Found) bool
	GetFound() (Found, bool)
}

// SQLite3
type SQLite struct {
	DB	*sql.DB
}

type User struct {
	ID						uint32
	Email					string
	Password				string
	Role					uint8
	DateOfRegistory			string
	DurationOfThePrivilege	string
}

type Request struct {
	ID						uint32
	Email					string
	DateOfRegistory			string
	Status					bool
	ValueDay				uint8
}

type Found struct {
	ID						uint32
	Email					string
	Kor						uint32
	Name					string
	Status					string
}

func NewStorageSQLite() Storage {
	var storage Storage

	path, _ := os.Getwd()

	db, err := sql.Open("sqlite3", strings.Replace(path, "\\", "/", -1) + "/application/storage/castleHunter.db")

	if err != nil {
		log.Fatal(err)
	}

	storage = &SQLite{
		DB: db,
	}

	return storage
}

// Строку в Md5
func GetMd5(text string) string {
	h := md5.New()
	h.Write([]byte(text))

	return hex.EncodeToString(h.Sum(nil))
}
