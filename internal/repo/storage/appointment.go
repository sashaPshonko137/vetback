package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
	"vetback/internal/model"
)

func (s *Storage) CreateAppointment(model model.AppointmentInfo) error {
	stmt, err := s.db.Prepare(`INSERT INTO appointments (doctor_id, owner_id, animal_id, diagnosis_id, date, status, reason) VALUES ($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(model.DoctorId, model.OwnerId, model.AnimalId, model.DiagnosisId, model.Date, model.Status, model.Reason)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetAppointment(id int) (*model.Appointment, error) {
	stmt, err := s.db.Prepare(`SELECT appointment_id, doctor_id, owner_id, animal_id, diagnosis_id, date, status, reason FROM appointments WHERE appointment_id = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var res model.Appointment
	err = stmt.QueryRow(id).Scan(&res.AppointmentId, &res.DoctorId, &res.OwnerId, &res.AnimalId, &res.DiagnosisId, &res.Date, &res.Status, &res.Reason)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	parsedData, _ := time.Parse("2006-01-02T15:04:05Z", res.Date)
	res.Date = parsedData.Format("2006-01-02 15:04")
	return &res, nil
}

func (s *Storage) GetAppointmentsByAnimal(id int) ([]model.Appointment, error) {
	stmt, err := s.db.Prepare(`SELECT * FROM appointments WHERE animal_id = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appointments []model.Appointment

	for rows.Next() {
		var appointment model.Appointment // Изменено на значение, а не указатель

		err = rows.Scan(&appointment.AppointmentId, &appointment.DoctorId, &appointment.OwnerId, &appointment.AnimalId, &appointment.DiagnosisId, &appointment.Date, &appointment.Status, &appointment.Reason)
		if err != nil {
			return nil, err
		}

		// Парсинг даты с учетом формата
		parsedData, err := time.Parse("2006-01-02T15:04:05Z", appointment.Date) // Изменен формат
		if err != nil {
			return nil, err
		}
		appointment.Date = parsedData.Format("2006-01-02 15:04")

		appointments = append(appointments, appointment) // Сохраняем значение, а не указатель
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return appointments, nil
}

func (s *Storage) GetManyAppointments() ([]model.Appointment, error) {
	stmt, err := s.db.Prepare(`SELECT * FROM appointments`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var res []model.Appointment
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var appointment model.Appointment

		err = rows.Scan(&appointment.AppointmentId, &appointment.DoctorId, &appointment.OwnerId, &appointment.AnimalId, &appointment.DiagnosisId, &appointment.Date, &appointment.Status, &appointment.Reason)
		if err != nil {
			return nil, err
		}

		// Парсинг даты с учетом правильного формата
		parsedData, err := time.Parse("2006-01-02T15:04:05Z", appointment.Date)
		if err != nil {
			return nil, fmt.Errorf("failed to parse appointment date: %w", err)
		}
		appointment.Date = parsedData.Format("2006-01-02 15:04")

		res = append(res, appointment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Storage) DeleteAppointment(id int) error {
	stmt, err := s.db.Prepare(`DELETE FROM appointments WHERE appointments.appointment_id = $1`)
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

func (s *Storage) UpdateAppointment(id int, model model.AppointmentToUpdate) error {
	queryParts := []string{}
	args := []interface{}{}
	argId := 1

	if model.OwnerId > 0 {
		queryParts = append(queryParts, `owner_id = $`+fmt.Sprint(argId))
		args = append(args, model.OwnerId)
		argId++
	}

	if model.AnimalId > 0 {
		queryParts = append(queryParts, `animal_id = $`+fmt.Sprint(argId))
		args = append(args, model.AnimalId)
		argId++
	}

	if model.DiagnosisId > 0 {
		queryParts = append(queryParts, `diagnosis_id = $`+fmt.Sprint(argId))
		args = append(args, model.DiagnosisId)
		argId++
	}

	if model.Date != "" {
		queryParts = append(queryParts, `date = $`+fmt.Sprint(argId))
		args = append(args, model.Date)
		argId++
	}

	if model.Status != "" {
		queryParts = append(queryParts, `status = $`+fmt.Sprint(argId))
		args = append(args, model.Status)
		argId++
	}

	if model.Reason != "" {
		queryParts = append(queryParts, `reason = $`+fmt.Sprint(argId))
		args = append(args, model.Reason)
		argId++
	}

	if len(queryParts) == 0 {
		return nil
	}

	query := `UPDATE appointments SET ` + strings.Join(queryParts, ", ") + ` WHERE appointment_id = $` + fmt.Sprint(argId)
	args = append(args, id)

	_, err := s.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
