package dto

import (
	"base-gin/domain"
	"base-gin/domain/dao"
	"time"
)

type AuthorFilter struct {
	Keyword string `form:"q" binding:"omitempty"`
	Start   int    `form:"s" binding:"omitempty,min=0"`
	Limit   int    `form:"l" binding:"omitempty,min=1"`
}

type AuthorDetailResp struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Gender   string `json:"gender"`
}

func (o *AuthorDetailResp) FromEntity(item *dao.Author) {
	var gender string
	if item.Gender == nil {
		gender = "-"
	} else if *item.Gender == domain.GenderFemale {
		gender = "wanita"
	} else {
		gender = "pria"
	}

	o.Fullname = item.Fullname
	o.Gender = gender
	o.ID = int(item.ID)
}

type AuthorUpdateReq struct {
	ID           uint      `json:"-"`
	Fullname     string    `json:"fullname" binding:"required,min=4,max=56"`
	Gender       string    `json:"gender" binding:"required,oneof=m f"`
	BirthDateStr string    `json:"birth_date" binding:"required,datetime=2006-01-02"`
	BirthDate    time.Time `json:"-"`
}	

func (o *AuthorUpdateReq) GetGender() domain.TypeGender {
	if o.Gender == "f" {
		return domain.GenderFemale
	}

	return domain.GenderMale
}

func (o *AuthorUpdateReq) GetBirthDate() (time.Time, error) {
	return time.Parse("2006-01-02", o.BirthDateStr)
}
