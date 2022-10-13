package postgres

import (
	"GraphNeo4jGO/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

var (
	ErrNoRowsAffected = errors.New("no rows affected")
	ErrNoRowFound     = errors.New("no rows found with")
)

// Create implements repository.User
// creates a new user in database and return its id
func (p *postgres) Create(ctx context.Context, u *model.User) (uint, error) {
	var id uint
	err := p.db.QueryRow(
		ctx,
		"INSERT INTO users (username, name, email, password, gender) VALUES ($1,$2,$3,$4,$5) RETURNING user_id",
		u.Username, u.Name, u.Email, u.Password, u.Gender,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// Delete implements repository.User
// it returns nil if delete operation was successful
func (p *postgres) Delete(ctx context.Context, id uint) error {
	res, err := p.db.Exec(
		ctx,
		"DELETE FROM users WHERE user_id=$1", id,
	)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return ErrNoRowsAffected
	}

	return nil
}

// Update implements repository.User
// update user in database from u
func (p *postgres) Update(ctx context.Context, u *model.User) error {
	_, err := p.db.Exec(
		ctx,
		"UPDATE users SET username=$1,name=$2,password=$3,email=$4,gender=$5 WHERE user_id=$6",
		u.Username, u.Name, u.Password, u.Email, u.Gender, u.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

// User implements repository.User
// this function get a user with given id from database, return ErrNoRowsFound if no user found
func (p *postgres) User(ctx context.Context, id uint) (*model.User, error) {
	u := model.User{}
	err := p.db.QueryRow(
		ctx,
		"SELECT user_id, username,name,email,password,gender FROM users WHERE user_id=$1", id,
	).Scan(&u.ID, &u.Username, &u.Name, &u.Email, &u.Password, &u.Gender)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNoRowFound
		}
		return nil, err
	}
	return &u, nil
}

// UserWithUsername implements repository.User
func (p *postgres) UserWithUsername(ctx context.Context, username string) (*model.User, error) {
	u := model.User{}
	err := p.db.QueryRow(
		ctx,
		"SELECT user_id, username,name,email,password,gender FROM users WHERE username=$1", username,
	).Scan(&u.ID, &u.Username, &u.Name, &u.Email, &u.Password, &u.Gender)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNoRowFound
		}
		return nil, err
	}
	return &u, nil
}
