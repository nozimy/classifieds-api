package model

type Ad struct {
	ID          int64    `json:"id,omitempty"`
	Name        string   `json:"name"`
	Created     string   `json:"created,omitempty"`
	Price       float64  `json:"price"`
	Description string   `json:"description,omitempty"`
	Photos      []string `json:"photos,omitempty"`
	PreviewImg  string   `json:"previewImg"`
}

type Ads []*Ad
