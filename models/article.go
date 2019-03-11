package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Article 模型
type Article struct {
	Model
	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	CreatedBy   string `json:"created_by" time_format:"20060102150405"`
	// UpdatedAt time.Time `json:"updated_at" time_format:time.RFC850`
	UpdatedAt time.Time `json:"updated_at" time_format:"20060102150405"`
	UpdatedBy string    `json:"updated_by"`
	State     int       `json:"state"`
}

// BeforeCreate 创建Article时设置更新时间，创建时间
func (Article *Article) BeforeCreate(scope *gorm.Scope) error {
	err:=scope.SetColumn("CreatedAt", time.Now().Unix())
	scope.SetColumn("UpdatedAt", time.Now().Unix())
	return err
}

// BeforeUpdate 更新Article时设置更新时间
func (Article *Article) BeforeUpdate(scope *gorm.Scope) error {
	err:=scope.SetColumn("UpdatedAt", time.Now().Unix())
	return err
}

// GetArticles 获取需要的文章
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

// GetArticle 获取一篇文章
func GetArticle(id int) (article Article) {
	db.Where("id=?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}

// GetArticleTotal 获取文章的总数
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

// CreateArticle 创建新文章
func CreateArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:       data["tag_id"].(int),
		Title:       data["title"].(string),
		Description: data["description"].(string),
		Content:     data["content"].(string),
		CreatedBy:   data["created_by"].(string),
		State:       data["state"].(int),
	})
	return true
}

// ExistArticleByID Article 是否存在
func ExistArticleByID(id int) bool {
	var Article Article
	db.Select("id").Where("id = ?", id).First(&Article)
	if Article.ID > 0 {
		return true
	}

	return false
}

// DeleteArticle 删除Article
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(&Article{})

	return true
}

// EditArticle 编辑Article
func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}
