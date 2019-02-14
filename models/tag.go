package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Tag 模型
type Tag struct {
	Model

	Name      string `json:"name"`
	CreatedBy string `json:"created_by"`
	// UpdatedAt time.Time `json:"updated_at" time_format:time.RFC850`
	UpdatedAt time.Time `json:"updated_at" time_format:"20060102150405"`
	UpdatedBy string    `json:"updated_by"`
	State     int       `json:"state"`
}

// BeforeCreate 创建tag时设置更新时间，创建时间
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now().Unix())
	scope.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}

// BeforeUpdate 更新tag时设置更新时间
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}

// GetTags 获取需要的标签
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

// GetTagTotal 获取标签的总数
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

// ExistTagByName 是否有相同的标签
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

// CreateTag 创建新标签
func CreateTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}

// ExistTagByID tag 是否存在
func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

// DeleteTag 删除tag
func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

// EditTag 编辑tag
func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}
