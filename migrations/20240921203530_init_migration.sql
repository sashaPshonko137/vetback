-- +goose Up
CREATE TABLE IF NOT EXISTS users (
                                     user_id SERIAL PRIMARY KEY,
                                     hash text NOT NULL,
                                     full_name varchar(50) NOT NULL,
                                     address varchar(50) NOT NULL,
                                     phone_number varchar(50) NOT NULL UNIQUE,
                                     role varchar(50) NOT NULL,
                                     specialization varchar(50)
);

CREATE INDEX IF NOT EXISTS idx_user_id ON users(user_id);

CREATE TABLE IF NOT EXISTS diagnosis (
                                         diagnosis_id SERIAL PRIMARY KEY,
                                         doctor_id int REFERENCES users (user_id),
                                         date date NOT NULL,
                                         name varchar(50) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_diagnosis_id ON diagnosis(diagnosis_id);

CREATE TABLE IF NOT EXISTS animals (
    animal_id SERIAL PRIMARY KEY,
    doctor_id int REFERENCES users (user_id),
    owner_id int REFERENCES users (user_id),
    diagnosis_id int REFERENCES diagnosis (diagnosis_id),
    name varchar(50) NOT NULL,
    birthdate date NOT NULL,
    breed varchar(50) NOT NULL,
    sex varchar(10) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_animals_id ON animals(animal_id);

CREATE TABLE IF NOT EXISTS treatments (
    name varchar(50) NOT NULL,
    treatment_id SERIAL PRIMARY KEY,
    doctor_id int REFERENCES users (user_id),
    animal_id int REFERENCES animals (animal_id),
    diagnosis_id int REFERENCES diagnosis (diagnosis_id),
    start date NOT NULL,
    finish date NOT NULL,
    price float8 NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_treatment_id ON treatments(treatment_id);

CREATE TABLE IF NOT EXISTS appointments (
    appointment_id SERIAL PRIMARY KEY,
    doctor_id int REFERENCES users (user_id),
    owner_id int REFERENCES users (user_id),
    animal_id int REFERENCES animals (animal_id),
    diagnosis_id int REFERENCES diagnosis (diagnosis_id),
    date date NOT NULL,
    status varchar(50) NOT NULL,
    reason varchar(200)
);

CREATE INDEX IF NOT EXISTS idx_appointment_id ON appointments(appointment_id);

-- +goose Down
DROP TABLE animals;
DROP TABLE users;
DROP TABLE diagnosis;
DROP TABLE treatments;
DROP TABLE appointments;