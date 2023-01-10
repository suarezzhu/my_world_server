package models

import (
	"github.com/jinzhu/gorm"
)

type Site struct {
	Model
	Name   string `json:"name"`
	Url    string `json:"url"`
	Notes  string `json:"notes"`
	Tags   string `json:"tags"`
	TypeId string `json:"type_id"`
}

func AddSite(name string, url string, notes string, tags string, typeId string) error {
	site := Site{
		Name:   name,
		Url:    url,
		Notes:  notes,
		Tags:   tags,
		TypeId: typeId,
	}
	if err := db.Create(&site).Error; err != nil {
		return err
	}
	return nil
}

func GetAllSites() ([]*Site) {
	//err := db.Raw("select * from m_user where id=?", id).First(&user).Error
	var sites []*Site
	err := db.Raw("select * from m_site").Scan(&sites).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return sites
}

func GetSites(pageNum int, pageSize int, maps interface{}) (site []Site) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&site)

	return
}
func GetSiteTotal(maps interface{}) (count int) {
	db.Model(&Site{}).Where(maps).Count(&count)

	return
}
