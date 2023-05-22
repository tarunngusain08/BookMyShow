package repos

import (
	"database/sql"

	"github.com/dunzoit/BookMyShow/models"
)

const (
	getAuditoriums   = `SELECT a.id, a.name, t.id, t.name, t.address, c.id, c.name, c.state FROM auditorium AS a JOIN theatre AS t ON a.theatre_id = t.id JOIN city AS c ON t.city_id = c.id`
	addAuditorium    = `INSERT INTO auditorium (name, theatre_id) VALUES ($1, $2) RETURNING id`
	getAuditorium    = `SELECT a.id, a.name, t.id, t.name, t.address, c.id, c.name, c.state FROM auditorium AS a JOIN theatre AS t ON a.theatre_id = t.id JOIN city AS c ON t.city_id = c.id WHERE a.id = $1`
	updateAuditorium = `UPDATE auditorium SET name = $1 WHERE id = $2`
	deleteAuditorium = `DELETE FROM auditorium WHERE id = $1`
)

type Auditorium struct {
	db *sql.DB
}

func NewAuditoriumRepo(db *sql.DB) *Auditorium {
	return &Auditorium{
		db: db,
	}
}

func (a *Auditorium) GetAuditoriums() ([]*models.Auditorium, error) {
	rows, err := a.db.Query(getAuditoriums)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var auditoriums []*models.Auditorium
	for rows.Next() {
		var id int
		var name string
		var theatreID int
		var theatreName, theatreAddress string
		var cityID int
		var cityName, state string
		if err := rows.Scan(&id, &name, &theatreID, &theatreName, &theatreAddress, &cityID, &cityName,
			&state); err != nil {
			return nil, err
		}
		city := &models.City{
			Id:    cityID,
			Name:  cityName,
			State: state,
		}
		theatre := &models.Theatre{
			ID:      theatreID,
			Name:    theatreName,
			Address: theatreAddress,
			City:    city,
		}
		auditorium := &models.Auditorium{
			ID:      id,
			Name:    name,
			Theatre: theatre,
		}
		auditoriums = append(auditoriums, auditorium)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return auditoriums, nil
}

func (a *Auditorium) AddAuditorium(auditorium *models.Auditorium) error {
	var auditoriumID int
	err := a.db.QueryRow(addAuditorium, auditorium.Name, auditorium.Theatre.ID).Scan(&auditoriumID)
	if err != nil {
		return err
	}
	auditorium.ID = auditoriumID
	return nil
}

func (a *Auditorium) GetAuditorium(auditoriumID int) (*models.Auditorium, error) {
	row := a.db.QueryRow(getAuditorium, auditoriumID)
	var id int
	var name string
	var theatreID int
	var theatreName, theatreAddress string
	var cityID int
	var cityName, state string
	err := row.Scan(&id, &name, &theatreID, &theatreName, &theatreAddress, &cityID, &cityName, &state)
	if err != nil {
		return nil, err
	}
	city := &models.City{
		Id:    cityID,
		Name:  cityName,
		State: state,
	}
	theatre := &models.Theatre{
		ID:      theatreID,
		Name:    theatreName,
		Address: theatreAddress,
		City:    city,
	}
	auditorium := &models.Auditorium{
		ID:      id,
		Name:    name,
		Theatre: theatre,
	}
	return auditorium, nil
}

func (a *Auditorium) UpdateAuditorium(auditoriumID int, updatedValues *models.Auditorium) error {
	_, err := a.db.Exec(updateAuditorium, updatedValues.Name, auditoriumID)
	if err != nil {
		return err
	}
	return nil
}

func (a *Auditorium) DeleteAuditorium(auditoriumID int) error {
	_, err := a.db.Exec(deleteAuditorium, auditoriumID)
	if err != nil {
		return err
	}
	return nil
}
