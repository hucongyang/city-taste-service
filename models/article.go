package models

type Article struct {
	Model

	TagId int `json:"tagId"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Img string `json:"img"`
	BiliVideoUrlWeChat string `json:"biliVideoUrlWeChat"`
	Content string `json:"content"`
	ShopId int `json:"shopId"`
	UpId int `json:"upId"`
	Lng string `json:"lng"`
	Lat string `json:"lat"`
	CreatedOn int `json:"createdOn"`
	CreatedBy string `json:"createdBy"`
	DeletedOn int `json:"deletedOn"`
	State int `json:"state"`
	Views string `json:"views"`
	Comments string `json:"comments"`
}

// 获取文章列表
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}
// 获取文章总数
func GetArticleTotal(maps interface{}) (count int) {
	
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}