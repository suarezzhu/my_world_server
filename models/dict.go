package models

import "github.com/jinzhu/gorm"

type Dict_data struct {
	Model
	Key       string `json:"key"`
	Value     string `json:"value"`
	Type_name string `json:"type_name"`
	Type_id   string `json:"type_id"`
}

func DictDataByTypeName(dictTypeName string) ([]*Dict_data, error) {
	var dicts []*Dict_data
	err := db.Raw("select * from m_dict_data where type_code = ? ", dictTypeName).Scan(&dicts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return dicts, nil
}
