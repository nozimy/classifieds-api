package usecase

import "classifieds-api/internal/model"

type Usecase interface {
	FindAd(id string, params map[string][]string) (*model.Ad, error)
	GetAds(params map[string][]string) (model.Ads, error)
	CreateAd(ad *model.Ad) (*model.Ad, error)
}
