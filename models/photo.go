package models

type Photo struct {
	id        int64
	ImgSrc    string `json:"img_src"`
	EarthDate string `json:"earth_date"`
}
