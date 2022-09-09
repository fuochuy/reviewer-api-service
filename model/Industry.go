package model

import "google.golang.org/genproto/googleapis/type/date"

type Industry struct {
	id int64 `json:"id"`

	name string `json:"name"`

	nameEn string `json:"name_en"`

	created date.Date `json:"created"`

	updated date.Date `json:"updated"`

	description string `json:"description"`

	organizations string `json:"organizations"`

	show string `json:"show"`
}
