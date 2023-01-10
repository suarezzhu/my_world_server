package models

type Site struct {
	Model
	Name  string `json:"name"`
	Url   string `json:"url"`
	Notes string `json:"notes"`
	Tags  string `json:"tags"`
}

func AddSite(name string, url string, notes string, tags string) error {
	site := Site{
		Name:  name,
		Url:   url,
		Notes: notes,
		Tags:  tags,
	}
	if err := db.Create(&site).Error; err != nil {
		return err
	}
	return nil
}

func GetSites(pageNum int, pageSize int, maps interface{}) (site []Site) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&site)

	return
}
func GetSiteTotal(maps interface{}) (count int) {
	db.Model(&Site{}).Where(maps).Count(&count)

	return
}
