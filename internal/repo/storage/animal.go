package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
	"vetback/internal/model"
)

func (s *Storage) CreateAnimal(model model.AnimalInfo) error {
	stmt, err := s.db.Prepare(`INSERT INTO animals (doctor_id, owner_id, diagnosis_id, name, birthdate, breed, sex) VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(model.DoctorId, model.OwnerId, model.DiagnosisId, model.Name, model.Birthdate, model.Breed, model.Sex)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetAnimal(id int) (*model.Animal, error) {
	stmt, err := s.db.Prepare(`SELECT * FROM animals WHERE animal_id = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res := model.Animal{}

	err = stmt.QueryRow(id).Scan(&res.AnimalId, &res.DoctorId, &res.OwnerId, &res.DiagnosisId, &res.Name, &res.Birthdate, &res.Breed, &res.Sex)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	parsedDate, _ := time.Parse("2006-01-02T15:04:05Z", res.Birthdate)
	res.Birthdate = parsedDate.Format("2006-01-02")

	return &res, nil
}

func (s *Storage) GetManyAnimals() ([]model.Animal, error) {
	stmt, err := s.db.Prepare(`SELECT * FROM animals`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var res []model.Animal

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var animal model.Animal
		if err := rows.Scan(&animal.AnimalId, &animal.DoctorId, &animal.OwnerId, &animal.DiagnosisId, &animal.Name, &animal.Birthdate, &animal.Breed, &animal.Sex); err != nil {
			return nil, err
		}
		parsedDate, _ := time.Parse("2006-01-02T15:04:05Z", animal.Birthdate)
		animal.Birthdate = parsedDate.Format("2006-01-02")
		res = append(res, animal)

	}

	return res, nil
}

func (s *Storage) DeleteAnimal(id int) error {
	stmt, err := s.db.Prepare(`DELETE FROM animals WHERE animal_id = $1`)
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

func (s *Storage) UpdateAnimal(id int, model model.AnimalToUpdate) error {
	queryParts := []string{}
	args := []interface{}{}
	argId := 1

	// Добавляем поля в запрос только если они заданы
	if model.DoctorId > 0 {
		queryParts = append(queryParts, `doctor_id = $`+fmt.Sprint(argId))
		args = append(args, model.DoctorId)
		argId++
	}

	if model.DiagnosisId > 0 {
		queryParts = append(queryParts, `diagnosis_id = $`+fmt.Sprint(argId))
		args = append(args, model.DiagnosisId)
		argId++
	}

	if model.Name != "" {
		queryParts = append(queryParts, `name = $`+fmt.Sprint(argId))
		args = append(args, model.Name)
		argId++
	}

	if model.Birthdate != "" {
		queryParts = append(queryParts, `birthdate = $`+fmt.Sprint(argId))
		args = append(args, model.Birthdate)
		argId++
	}

	if model.Breed != "" {
		queryParts = append(queryParts, `breed = $`+fmt.Sprint(argId))
		args = append(args, model.Breed)
		argId++
	}

	if model.Sex != "" {
		queryParts = append(queryParts, `sex = $`+fmt.Sprint(argId))
		args = append(args, model.Sex)
		argId++
	}

	// Если нет полей для обновления, ничего не делаем
	if len(queryParts) == 0 {
		return nil
	}

	// Финальный запрос
	query := `UPDATE animals SET ` + strings.Join(queryParts, ", ") + ` WHERE animal_id = $` + fmt.Sprint(argId)
	args = append(args, id)

	// Выполняем запрос
	_, err := s.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
