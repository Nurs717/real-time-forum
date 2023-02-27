package repository

import (
	"context"
	"database/sql"
	"log"
	"rtforum/internal/cerror"
	"rtforum/internal/entity"
	"strings"
	"time"
)

type UsersRepo struct {
	db      *sql.DB
	timeout time.Duration
}

func NewUsersRepo(db *sql.DB, timeout time.Duration) *UsersRepo {
	return &UsersRepo{
		db:      db,
		timeout: timeout,
	}
}

func (r *UsersRepo) NewUser(ctx context.Context, user *entity.User) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxWithTimeout, "INSERT INTO Users (ID, UserName, Age, Gender, First_Name, Last_name, Mail, Password) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", user.ID, user.UserName, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		if strings.Contains(err.Error(), "UserName") {
			return cerror.WrapErrorf(err, cerror.ErrorCodeConflict, cerror.UserType, "user already exists")
		} else if strings.Contains(err.Error(), "Mail") {
			return cerror.WrapErrorf(err, cerror.ErrorCodeConflict, cerror.MailType, "email already exists")
		} else {
			return cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: NewUser: exec db")
		}
	}
	return nil
}

func (r *UsersRepo) GetUser(ctx context.Context, mail string) (string, string, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxWithTimeout, "SELECT ID, Mail, Password from Users")
	if err != nil {
		return "", "", cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: getUser: query db")
	}
	defer rows.Close()
	var id string
	var scanMail string
	var password string
	for rows.Next() {
		err := rows.Scan(&id, &scanMail, &password)
		if err != nil {
			return "", "", cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: getUser: scanning rows db")
		}
	}
	if err = rows.Err(); err != nil {
		return "", "", cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: getUser: catching rows error db")
	}
	if mail == scanMail {
		return id, password, nil
	}
	return "", "", cerror.NewErrorf(cerror.ErrorCodeUnauthorized, cerror.DefaultType, cerror.ErrMailNotExist)
}

func (r *UsersRepo) AddCookie(ctx context.Context, id string, cookieValue string, dt time.Time) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxWithTimeout, "INSERT INTO Session (Token, Expired_Date, User_ID) VALUES (?, ?, ?)", cookieValue, dt, id)
	if err != nil {
		return cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: AddCookie: add session to db")
	}
	return nil
}

func (r *UsersRepo) GetUserIDbyCookie(ctx context.Context, token string) (string, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxWithTimeout, "SELECT Token, User_ID from Session WHERE Expired_Date > DATETIME('now')")
	if err != nil {
		log.Printf("error occured getUser from db: %v\n", err)
		return "", err
	}
	defer rows.Close()
	for rows.Next() {
		var id string
		var tokenScan string
		err := rows.Scan(&tokenScan, &id)
		if err != nil {
			log.Printf("error occured getUser scanning rows from db: %v\n", err)
			continue
		}
		if token == tokenScan {
			return id, nil
		}
	}
	return "", cerror.ErrTokenInvalid
}
