package storage

import (
	"database/sql"
	"errors"
	"time"
	"vetback/internal/model"
)

func (s *Storage) CreateDiagnosis(model model.DiagnosisInfo) error {
	stmt, err := s.db.Prepare(`INSERT INTO diagnosis (doctor_id, name, date) VALUES ($1, $2, $3)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(model.DoctorId, model.Name, time.Now().Format("2006-01-02"))
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetDiagnosis(id int) (*model.Diagnosis, error) {
	stmt, err := s.db.Prepare(`SELECT diagnosis_id, doctor_id, name, date FROM diagnosis WHERE diagnosis_id = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var res model.Diagnosis
	err = stmt.QueryRow(id).Scan(&res.DiagnosisId, &res.DoctorId, &res.Name, &res.Date)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	parsedData, _ := time.Parse("2006-01-02T15:04:05Z", res.Date)
	res.Date = parsedData.Format("2006-01-02")

	return &res, nil
}

func (s *Storage) GetManyDiagnosis() ([]model.Diagnosis, error) {
	stmt, err := s.db.Prepare(`SELECT * FROM diagnosis`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var res []model.Diagnosis
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var diagnosis model.Diagnosis

		err = rows.Scan(&diagnosis.DiagnosisId, &diagnosis.DoctorId, &diagnosis.Date, &diagnosis.Name)
		if err != nil {
			return nil, err
		}
		parsedData, _ := time.Parse("2006-01-02T15:04:05Z", diagnosis.Date)
		diagnosis.Date = parsedData.Format("2006-01-02")
		res = append(res, diagnosis)
	}
	return res, nil
}

func (s *Storage) UpdateDiagnosis(id int, model model.DiagnosisToUpdate) error {
	now := time.Now().Format("2006-01-02")
	stmt, err := s.db.Prepare(`UPDATE diagnosis SET name = $1, date = $2 WHERE diagnosis_id = $3`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(model.Name, now, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteDiagnosis(id int) error {
	stmt, err := s.db.Prepare(`DELETE FROM diagnosis WHERE diagnosis_id = $1`)
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
