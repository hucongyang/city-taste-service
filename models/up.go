package models

type Up struct {
	Model

	Type int `json:"type"`
	Name string `json:"name"`
	Fans string `json:"fans"`
	Desc string `json:"desc"`
	Img string `json:"string"`
	Track string `json:"track"`
	CreatedBy string `json:"created_by"`
	DeletedOn int `json:"deleted_on"`

}

// 获取单个文章
func GetUp(id int) (up Up) {
	db.Where("id = ?", id).First(&up)
	db.Model(&up)
	return
}


// 获取up列表
func GetUps(pageNum int, pageSize int, maps interface{}) (ups []Up) {

	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&ups)
	return
}

// 获取up总数
func GetUpTotal(maps interface{}) (count int) {

	db.Model(&Up{}).Where(maps).Count(&count)
	return
}

// up是否存在
func ExistUpByID(id int) bool {
	var up Up
	db.Select("id").Where("id = ?", id).First(&up)
	return up.ID > 0
}