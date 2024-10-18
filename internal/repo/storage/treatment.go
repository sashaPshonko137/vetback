package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
	"vetback/internal/model"
)

func (s *Storage) CreateTreatment(model model.TreatmentInfo) error {
	stmt, err := s.db.Prepare(`INSERT INTO treatments (doctor_id, name, animal_id, diagnosis_id, start, finish, price) VALUES ($1, $2, $3, $4, $5, $6, $7)`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(model.DoctorId, model.Name, model.AnimalId, model.DiagnosisId, time.Now().Format("2006-01-02"), model.Finish, model.Price)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetTreatment(id int) (*model.Treatment, error) {
	stmt, err := s.db.Prepare(`SELECT treatment_id, name, doctor_id, animal_id, diagnosis_id, start, finish, price FROM treatments WHERE treatment_id = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var res model.Treatment
	err = stmt.QueryRow(id).Scan(&res.TreatmentId, &res.Name, &res.DoctorId, &res.AnimalId, &res.DiagnosisId, &res.Start, &res.Finish, &res.Price)
	parsedData, _ := time.Parse("2006-01-02T15:04:05Z", res.Start)
	res.Start = parsedData.Format("2006-01-02")
	parsedData, _ = time.Parse("2006-01-02T15:04:05Z", res.Finish)
	res.Finish = parsedData.Format("2006-01-02")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &res, nil
}

func (s *Storage) GetManyTreatmentsByAnimal(id int) ([]model.Treatment, error) {
	stmt, err := s.db.Prepare(`SELECT treatment_id, name, doctor_id, animal_id, diagnosis_id, start, finish, price FROM treatments WHERE animal_id = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var res []model.Treatment
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var treatment model.Treatment

		err = rows.Scan(&treatment.TreatmentId, &treatment.Name, &treatment.DoctorId, &treatment.AnimalId, &treatment.DiagnosisId, &treatment.Start, &treatment.Finish, &treatment.Price)
		if err != nil {
			return nil, err
		}

		// Парсинг времени с учетом правильного формата
		parsedStart, err := time.Parse("2006-01-02T15:04:05Z", treatment.Start)
		if err != nil {
			return nil, fmt.Errorf("failed to parse start time: %w", err)
		}
		treatment.Start = parsedStart.Format("2006-01-02")

		parsedFinish, err := time.Parse("2006-01-02T15:04:05Z", treatment.Finish)
		if err != nil {
			return nil, fmt.Errorf("failed to parse finish time: %w", err)
		}
		treatment.Finish = parsedFinish.Format("2006-01-02")

		res = append(res, treatment)
	}

	return res, nil
}

func (s *Storage) GetManyTreatments() ([]model.Treatment, error) {
	stmt, err := s.db.Prepare(`SELECT treatment_id, name, doctor_id, animal_id, diagnosis_id, start, finish, price FROM treatments`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var res []model.Treatment
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var treatment model.Treatment

		err = rows.Scan(&treatment.TreatmentId, &treatment.Name, &treatment.DoctorId, &treatment.AnimalId, &treatment.DiagnosisId, &treatment.Start, &treatment.Finish, &treatment.Price)
		if err != nil {
			return nil, err
		}
		parsedData, _ := time.Parse("2006-01-02T15:04:05Z", treatment.Start)
		treatment.Start = parsedData.Format("2006-01-02")
		parsedData, _ = time.Parse("2006-01-02T15:04:05Z", treatment.Finish)
		treatment.Finish = parsedData.Format("2006-01-02")
		res = append(res, treatment)
	}
	return res, nil
}

func (s *Storage) DeleteTreatment(id int) error {
	stmt, err := s.db.Prepare(`DELETE FROM treatments WHERE treatment_id = $1`)
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

func (s *Storage) UpdateTreatment(id int, model model.TreatmentToUpdate) error {
	queryParts := []string{}
	args := []interface{}{}
	argId := 1

	if model.AnimalId > 0 {
		queryParts = append(queryParts, `animal_id = $`+fmt.Sprint(argId))
		args = append(args, model.AnimalId)
		argId++
	}

	if model.Name != "" {
		queryParts = append(queryParts, `name = $`+fmt.Sprint(argId))
		args = append(args, model.Name)
		argId++
	}

	if model.DiagnosisId > 0 {
		queryParts = append(queryParts, `diagnosis_id = $`+fmt.Sprint(argId))
		args = append(args, model.DiagnosisId)
		argId++
	}

	if model.Finish != "" {
		queryParts = append(queryParts, `finish = $`+fmt.Sprint(argId))
		args = append(args, model.Finish)
		argId++
	}

	if model.Price > 0 {
		queryParts = append(queryParts, `price = $`+fmt.Sprint(argId))
		args = append(args, model.Price)
		argId++
	}

	if len(queryParts) == 0 {
		return nil
	}

	query := `UPDATE treatments SET ` + strings.Join(queryParts, ", ") + ` WHERE treatment_id = $` + fmt.Sprint(argId)
	args = append(args, id)

	_, err := s.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
