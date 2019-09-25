package models

type Photo struct {
	Id        int64 `json:"id"`
	ImgSrc    string `json:"img_src"`
	EarthDate string `json:"earth_date"`
}
