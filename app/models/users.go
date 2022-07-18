package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
	Todos     []Todo
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func (u *User) CreateUser() (err error) {
	cmd := `INSERT INTO users (
		uuid,
		name,
		email,
		password,
		created_at) VALUES (? ,? ,? ,? ,?)`

	_, err = Db.Exec(cmd, createUUID(), u.Name, u.Email, Encrypt(u.PassWord), time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetUser(id int) (User, error) {
	user := User{}
	cmd := `SELECT id, uuid, name, email, password, created_at FROM users WHERE id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)

	return user, err
}

func (u *User) UpdateUser() error {
	cmd := `UPDATE users SET name = ?, email = ? WHERE id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (u *User) DeleteUser() error {
	cmd := `DELETE FROM users WHERE id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetUserByEmail(email string) (User, error) {
	user := User{}
	cmd := `SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)

	return user, err
}

func (u *User) CreateSession() (Session, error) {
	session := Session{}
	cmd1 := `INSERT INTO sessions (
		uuid,
		email,
		user_id,
		created_at) VALUES (?, ?, ?, ?)`

	_, err := Db.Exec(cmd1, createUUID(), u.Email, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}

	cmd2 := `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = ? and email = ?`
	err = Db.QueryRow(cmd2, u.ID, u.Email).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserId,
		&session.CreatedAt,
	)
	return session, err

}

func (session *Session) CheckSession() (bool, error) {
	cmd := `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?`
	err := Db.QueryRow(cmd, session.UUID).Scan(
		&session.ID,
		&session.UUID,
		&session.Email,
		&session.UserId,
		&session.CreatedAt,
	)
	valid := false
	if err != nil {
		return valid, err
	}
	if session.ID != 0 {
		valid = true
	}
	return valid, err
}

func (session *Session) DeleteSessionByUUID() error {
	cmd := `DELETE FROM sessions WHERE uuid = ?`
	_, err := Db.Exec(cmd, session.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (session *Session) GetUserBySession() (User, error) {
	user := User{}
	cmd := `SELECT id, uuid, name, email, created_at FROM users WHERE id = ?`
	err := Db.QueryRow(cmd, session.UserId).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
	)

	return user, err
}
