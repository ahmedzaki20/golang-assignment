package models

import "example.com/golang-assignment/db"

type Animal struct {
	ID    int64    `json:"id"`
	Name  string `json:"name" binding:"required"`
	Age   int64    `json:"age" binding:"required"`
	Breed string `json:"breed" binding:"required"`
}

func (a *Animal) Create() error {
	query := `INSERT INTO animals (name, age, breed) VALUES (?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result,err:= stmt.Exec(a.Name, a.Age, a.Breed)
	if err != nil {
		return err
	}
	a.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func GetAnimals() ([]Animal, error) {
	query := `SELECT id, name, age, breed FROM animals`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	animals := []Animal{}
	for rows.Next() {
		a := Animal{}
		err := rows.Scan(&a.ID, &a.Name, &a.Age, &a.Breed)
		if err != nil {
			return nil, err
		}
		animals = append(animals, a)
	}
	return animals, nil
}

func GetAnimal(id int64) (Animal, error) {
	query := `SELECT id, name, age, breed FROM animals WHERE id = ?`
	row := db.DB.QueryRow(query, id)
	a := Animal{}
	err := row.Scan(&a.ID, &a.Name, &a.Age, &a.Breed)
	if err != nil {
		return Animal{}, err
	}
	return a, nil
}

func (a *Animal) Update() error {
	query := `UPDATE animals SET name = ?, age = ?, breed = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(a.Name, a.Age, a.Breed, a.ID)
	if err != nil {
		return err
	}
	return nil
}

func (a *Animal) Delete() error {
	query := `DELETE FROM animals WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(a.ID)
	if err != nil {
		return err
	}
	return nil
}