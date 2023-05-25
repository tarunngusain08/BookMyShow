package repos

import (
	"database/sql"

	"BookMyShow/models"
)

const (
	getTheatres = `SELECT t.id, t.name, t.address, c.id, c.name, c.state
	                    FROM theatre AS t
	                    JOIN city AS c ON t.city_id = c.id`
	addTheatre = `INSERT INTO theatre (name, address, city_id) VALUES ($1, $2, $3) RETURNING id`
	getTheatre = `SELECT t.id, t.name, t.address, c.id, c.name, c.state
	                    FROM theatre AS t
	                    JOIN city AS c ON t.city_id = c.id
	                    WHERE t.id = $1`
	updateTheatre = `UPDATE theatre SET name = $1, address = $2, city_id = $3 WHERE id = $4`
	deleteTheatre = `DELETE FROM theatre WHERE id = $1`
)

type Theatre struct {
	db *sql.DB
}

func NewTheatreRepo(db *sql.DB) *Theatre {
	return &Theatre{
		db: db,
	}
}

func (t *Theatre) GetTheatres(cityID int) ([]*models.Theatre, error) {
	rows, err := t.db.Query(getTheatres)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var theatres []*models.Theatre
	for rows.Next() {
		var id int
		var name string
		var address string
		var cityID int
		var cityName string
		var state string
		if err := rows.Scan(&id, &name, &address, &cityID, &cityName, &state); err != nil {
			return nil, err
		}
		city := &models.City{
			Id:    cityID,
			Name:  cityName,
			State: state,
		}
		theatre := &models.Theatre{
			ID:      id,
			Name:    name,
			Address: address,
			City:    city,
		}
		theatres = append(theatres, theatre)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return theatres, nil
}

func (t *Theatre) AddTheatre(theatre *models.Theatre) error {
	var theatreID int
	err := t.db.QueryRow(addTheatre, theatre.Name, theatre.Address, theatre.City.Id).Scan(&theatreID)
	if err != nil {
		return err
	}
	theatre.ID = theatreID
	return nil
}

func (t *Theatre) GetTheatre(theatreID int) (*models.Theatre, error) {
	row := t.db.QueryRow(getTheatre, theatreID)
	var id int
	var name string
	var address string
	var cityID int
	var cityName string
	var state string
	err := row.Scan(&id, &name, &address, &cityID, &cityName, &state)
	if err != nil {
		return nil, err
	}
	city := &models.City{
		Id:    cityID,
		Name:  cityName,
		State: state,
	}
	theatre := &models.Theatre{
		ID:      id,
		Name:    name,
		Address: address,
		City:    city,
	}
	return theatre, nil
}

func (t *Theatre) UpdateTheatre(theatreID int, updatedValues *models.Theatre) error {
	_, err := t.db.Exec(updateTheatre, updatedValues.Name, updatedValues.Address, updatedValues.City.Id, theatreID)
	if err != nil {
		return err
	}
	return nil
}

func (t *Theatre) DeleteTheatre(theatreID int) error {
	_, err := t.db.Exec(deleteTheatre, theatreID)
	if err != nil {
		return err
	}
	return nil
}
