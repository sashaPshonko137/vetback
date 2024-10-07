package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"vetback/internal/model"
)

func (s *Storage) CreateUser(model model.UserToStorage) (int, error) {
	stmt, err := s.db.Prepare(`INSERT INTO users (hash, full_name, address, phone_number, role, specialization) VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var userId int

	err = stmt.QueryRow(model.Hash, model.FullName, model.Address, model.PhoneNumber, model.Role, model.Specialization).Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (s *Storage) GetUser(id int) (*model.User, error) {
	row := s.db.QueryRow(`SELECT * FROM users WHERE user_id = $1`, id)

	res := model.User{}

	err := row.Scan(&res.UserId, &res.Hash, &res.FullName, &res.Address, &res.PhoneNumber, &res.Role, &res.Specialization)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &res, nil
}
func (s *Storage) GetUserByPhone(phoneNumber string) (*model.User, error) {
	row := s.db.QueryRow(`SELECT * FROM users WHERE phone_number = $1`, phoneNumber)

	res := model.User{}

	err := row.Scan(&res.UserId, &res.Hash, &res.FullName, &res.Address, &res.PhoneNumber, &res.Role, &res.Specialization)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &res, nil
}

func (s *Storage) GetManyUsers() ([]model.SafeUser, error) {
	rows, err := s.db.Query(`SELECT user_id, full_name, phone_number, address, role, specialization FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []model.SafeUser

	for rows.Next() {
		var user model.SafeUser
		err = rows.Scan(&user.UserId, &user.FullName, &user.PhoneNumber, &user.Address, &user.Role, &user.Specialization)
		if err != nil {
			return nil, err
		}
		res = append(res, user)
	}
	return res, nil
}

func (s *Storage) DeleteUser(id int) error {
	stmt, err := s.db.Prepare(`DELETE FROM users WHERE user_id = $1`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdateUser(id int, model model.UpdatedUserToStorage) error {
	queryParts := []string{}
	args := []interface{}{}
	argId := 1

	if model.FullName != "" {
		queryParts = append(queryParts, `full_name = $`+fmt.Sprint(argId))
		args = append(args, model.FullName)
		argId++
	}

	if model.Address != "" {
		queryParts = append(queryParts, `address = $`+fmt.Sprint(argId))
		args = append(args, model.Address)
		argId++
	}

	if model.PhoneNumber != "" {
		queryParts = append(queryParts, `phone_number = $`+fmt.Sprint(argId))
		args = append(args, model.PhoneNumber)
		argId++
	}

	if model.Specialization != "" {
		queryParts = append(queryParts, `specialization = $`+fmt.Sprint(argId))
		args = append(args, model.Specialization)
		argId++
	}

	if model.Hash != "" {
		queryParts = append(queryParts, `hash = $`+fmt.Sprint(argId))
		args = append(args, model.Specialization)
		argId++
	}

	if len(queryParts) == 0 {
		return nil
	}

	query := `UPDATE users SET ` + strings.Join(queryParts, ", ") + ` WHERE user_id = $` + fmt.Sprint(argId)
	args = append(args, id)

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}

	return nil
}
