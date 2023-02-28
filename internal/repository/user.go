package repository

import (
	"context"
	"database/sql"
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

	_, err := r.db.ExecContext(ctxWithTimeout,
		"INSERT INTO Users (ID, UserName, Age, Gender, First_Name, Last_name, Mail, Password) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		user.ID, user.UserName, user.Age, user.Gender, user.FirstName, user.LastName, user.Email, user.Password)
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

	var id string
	var password string
	rows, err := r.db.QueryContext(ctxWithTimeout, "SELECT ID, Password from Users WHERE Users.Mail = ?", mail)
	if err != nil {
		return "", "", cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: getUser: query")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &password)
		if err != nil {
			return "", "", cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: getUser: scanning rows")
		}
	}
	if err = rows.Err(); err != nil {
		return "", "", cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: getUser: catching rows error")
	}
	return id, password, nil
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

	var id string
	rows, err := r.db.QueryContext(ctxWithTimeout,
		"SELECT User_ID from Session WHERE Expired_Date > DATETIME('now') AND Session.Token = ?", token)
	if err != nil {
		return "", cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: GetUserIDbyCookie: Query")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return "", cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: GetUserIDbyCookie: Scan rows")
		}
	}
	if err = rows.Err(); err != nil {
		return "", cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: getUserIDbyCookie: catching rows error")
	}
	return id, nil
}

func (r *UsersRepo) GetUserName(ctx context.Context, userID string) (string, error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxWithTimeout, "SELECT UserName from Users WHERE Users.ID = ?", userID)
	if err != nil {
		return "", cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: GetUserName: Query")
	}
	defer rows.Close()
	var username string
	for rows.Next() {
		err = rows.Scan(&username)
		if err != nil {
			return "", cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: GetUserName: Scan rows")
		}
	}
	if err = rows.Err(); err != nil {
		return "", cerror.WrapErrorf(err, cerror.ErrorCodeInternal, cerror.DefaultType, "repo: getUserName: catching rows error")
	}
	return username, nil
}
