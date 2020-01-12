package repository

import (
	"classifieds-api/internal/model"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
)

type AdRepository struct {
	db *sql.DB
}

func (a AdRepository) FindAd(id string, includeDescription, includePhotos bool) (*model.Ad, error) {
	adObj := &model.Ad{}
	var description string
	var photos []string

	if err := a.db.QueryRow(
		"SELECT name, description, price, photos FROM ads WHERE id = $1",
		id,
	).Scan(
		&adObj.Name,
		&description,
		&adObj.Price,
		pq.Array(&photos),
	); err != nil {
		return nil, err
	}

	if len(photos) > 0 {
		adObj.PreviewImg = photos[0]
	}

	if includeDescription {
		adObj.Description = description
	}

	if includePhotos {
		adObj.Photos = photos
	}

	return adObj, nil
}

func (a AdRepository) GetAds(limit, offset int, sort, desc string) (model.Ads, error) {
	ads := make(model.Ads, 0)
	var photos []string

	query := "SELECT name, price, photos FROM ads "

	switch sort {
	case "date":
		query += fmt.Sprintf("ORDER BY created %s ", desc)
		break
	case "price":
		query += fmt.Sprintf("ORDER BY price %s ", desc)
		break
	}

	query += "LIMIT $1 OFFSET $2 "

	log.Println(query)

	rows, err := a.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		ad := model.Ad{}
		err := rows.Scan(
			&ad.Name,
			&ad.Price,
			pq.Array(&photos),
		)
		if err != nil {
			return nil, err
		}

		if len(photos) > 0 {
			ad.PreviewImg = photos[0]
		}

		ads = append(ads, &ad)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	return ads, nil
}

func (a AdRepository) CreateAd(ad *model.Ad) (*model.Ad, error) {
	query := "INSERT INTO ads (name, description, price, photos) VALUES ($1, $2, $3, $4) RETURNING id"

	if err := a.db.QueryRow(
		query,
		ad.Name,
		ad.Description,
		ad.Price,
		pq.Array(ad.Photos),
	).Scan(
		&ad.ID,
	); err != nil {
		return nil, err
	}

	return ad, nil
}

func NewAdRepository(db *sql.DB) Repository {
	return &AdRepository{db}
}
