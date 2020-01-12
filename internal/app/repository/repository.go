package repository

import "classifieds-api/internal/model"

type Repository interface {
	FindAd(id string, includeDescription, includePhotos bool) (*model.Ad, error)
	GetAds(limit, offset int, sort, desc string) (model.Ads, error)
	CreateAd(user *model.Ad) (*model.Ad, error)
}
