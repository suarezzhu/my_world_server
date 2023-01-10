package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Site struct {
	Model
	Name   string `json:"name"`
	Url    string `json:"url"`
	Note   string `json:"note"`
	Tag    string `json:"tag"`
	TypeId string `json:"type_id"`
}

func AddSite(name string, url string, note string, tag string, typeId string) error {
	site := Site{
		Name:   name,
		Url:    url,
		Note:   note,
		Tag:    tag,
		TypeId: typeId,
	}
	if err := db.Create(&site).Error; err != nil {
		return err
	}
	return nil
}

func GetAllSites() []*Site {
	//err := db.Raw("select * from m_user where id=?", id).First(&user).Error
	var sites []*Site
	err := db.Raw("select * from m_site").Scan(&sites).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return sites
}

func GetAllSites2(name string, url string, note string, tag string, pageNum int, pageSize int) []*Site {
	//err := db.Raw("select * from m_user where id=?", id).First(&user).Error
	var sites []*Site
	err :=
		//db.Raw("select * from m_dict_data a inner join m_site b on a.value= b.type_id").Where("name like ? AND url like  ? note like ? AND tag like ?", name, url, note, tag).Scan(&sites).Error
		db.Raw("select * from m_dict_data a inner join m_site b on a.value= b.type_id where name like ? AND url like  ? AND note like ? AND tag like ?", name, url, note, tag).Scan(&sites).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return sites
	fmt.Println(sites, "我是网站全部")
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
