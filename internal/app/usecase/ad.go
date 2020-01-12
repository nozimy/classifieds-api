package usecase

import (
	"classifieds-api/internal/app/repository"
	"classifieds-api/internal/model"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

type AdUsecase struct {
	rep repository.Repository
}

func (a AdUsecase) FindAd(id string, params map[string][]string) (*model.Ad, error) {
	fields := params["fields"]
	includeDescription := false
	includePhotos := false
	if len(fields) >= 1 {
		splitRelated := strings.Split(fields[0], ",")

		if contains(splitRelated, "description") {
			includeDescription = true
		}
		if contains(splitRelated, "photos") {
			includePhotos = true
		}
	}

	adObj, err := a.rep.FindAd(id, includeDescription, includePhotos)

	if err != nil {
		return nil, errors.Wrap(err, "rep.FindAd()")
	}

	return adObj, nil
}

func (a AdUsecase) GetAds(params map[string][]string) (model.Ads, error) {
	limit := 10
	offset := 0
	sort := "date"

	if len(params["page"]) >= 1 {
		page, _ := strconv.Atoi(params["page"][0])
		offset = (page - 1) * limit
	}

	if len(params["sort"]) >= 1 {
		sort = params["sort"][0]
	}

	desc := ""
	if len(params["desc"]) >= 1 && params["desc"][0] == "true" {
		desc = "desc"
	}

	ads, err := a.rep.GetAds(limit, offset, sort, desc)

	if err != nil {
		return nil, errors.Wrap(err, "rep.GetAds()")
	}

	return ads, nil
}

func (a AdUsecase) CreateAd(ad *model.Ad) (*model.Ad, error) {
	adObj, err := a.rep.CreateAd(ad)

	if err != nil {
		return nil, errors.Wrap(err, "rep.CreateAd()")
	}

	return adObj, nil
}

func NewAdUsecase(r repository.Repository) Usecase {
	return &AdUsecase{
		rep: r,
	}
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}

	return false
}
