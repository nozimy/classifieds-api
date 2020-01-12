package delivery

import (
	"classifieds-api/internal/app/usecase"
	"classifieds-api/internal/model"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type AdHandler struct {
	use usecase.Usecase
}

func NewAdHandler(m *mux.Router, u usecase.Usecase) {
	handler := &AdHandler{u}

	m.HandleFunc("/item/{id}", handler.HandleGetAd).Methods(http.MethodGet)
	m.HandleFunc("/items", handler.HandleGetAdList).Methods(http.MethodGet)
	m.HandleFunc("/items", handler.HandleCreateAd).Methods(http.MethodPost)
}

func (h *AdHandler) HandleGetAd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	adObj, err := h.use.FindAd(id, r.URL.Query())

	if err != nil {
		Error(w, r, http.StatusNotFound, errors.New("Not Found"))
		return
	}

	Respond(w, r, http.StatusOK, adObj)
}

func (h *AdHandler) HandleGetAdList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ads, err := h.use.GetAds(r.URL.Query())

	if err != nil {
		Error(w, r, http.StatusNotFound, errors.New("Not Found"))
		return
	}

	Respond(w, r, http.StatusOK, ads)
}

func (h *AdHandler) HandleCreateAd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defer func() {
		if err := r.Body.Close(); err != nil {
			Error(w, r, http.StatusInternalServerError, err)
		}
	}()

	decoder := json.NewDecoder(r.Body)
	newAd := new(model.Ad)
	err := decoder.Decode(newAd)
	if err != nil {
		Error(w, r, http.StatusBadRequest, err)
		return
	}

	adObj, err := h.use.CreateAd(newAd)

	if err != nil {
		Error(w, r, http.StatusConflict, err)
		return
	}

	Respond(w, r, http.StatusCreated, adObj.ID)
}
