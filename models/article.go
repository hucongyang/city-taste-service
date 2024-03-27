package models

type Article struct {
	Model

	UpId int `json:"up_id" gorm:"index"`
	Up Up `json:"up"`
	ShopId int `json:"shop_id" gorm:"index"`

	TagId int `json:"tag_id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Img string `json:"img"`
	BiliVideoUrlWeChat string `json:"bili_video_url_wechat"`
	Content string `json:"content"`
	BaiLng string `json:"bai_lng"`
	BaiLat string `json:"bai_lat"`
	// CreatedOn int `json:"createdOn"`
	CreatedBy string `json:"created_by"`
	DeletedOn int `json:"deleted_on"`
	State int `json:"state"`
	Views string `json:"views"`
	Comments string `json:"comments"`
}

// 获取单个文章
func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Up)
	return
}

// 获取文章列表
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	
	db.Preload("Up").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}
// 获取文章总数
func GetArticleTotal(maps interface{}) (count int) {
	
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

// 文章是否存在
func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	return article.ID > 0
}