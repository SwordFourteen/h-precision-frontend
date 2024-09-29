package model

import (
	"github.com/moznion/go-optional"                  // 使用 go-optional 库提供 Optional 类型，避免空指针问题
	"github.com/ybkuroki/go-webapp-sample/repository" // 引入 repository 包，定义了数据库操作的接口
)

// Format defines struct of format data.
// Format 结构体定义了格式数据的结构，与数据库表格式相对应。
type Format struct {
	ID   uint   `gorm:"primary_key" json:"id"`    // ID 字段为主键，GORM 负责管理。JSON 序列化时键名为 "id"
	Name string `validate:"required" json:"name"` // Name 字段为必填项，使用 validate 标签进行验证，JSON 键名为 "name"
}

// TableName returns the table name of format struct and it is used by gorm.
// TableName 方法返回此结构体对应的数据库表名，GORM 根据此方法映射表名。
func (Format) TableName() string {
	return "format_master"
}

// NewFormat is constructor
// NewFormat 是一个构造函数，用于创建并初始化一个 Format 结构体实例。
func NewFormat(name string) *Format {
	return &Format{Name: name}
}

// FindByID returns a format full matched given format's ID.
// FindByID 方法通过给定的 ID 从数据库中查找完全匹配的 Format 记录。
// 如果找到记录则返回 Optional 包含的 Format 对象，否则返回空的 Optional。
func (f *Format) FindByID(rep repository.Repository, id uint) optional.Option[*Format] {
	var format Format
	// 通过 ID 在数据库中查找 Format 记录，如果发生错误返回 None
	if err := rep.Where("id = ?", id).First(&format).Error; err != nil {
		return optional.None[*Format]()
	}
	// 如果查找成功，返回 Some 包含的 Format 对象
	return optional.Some(&format)
}

// FindAll returns all formats of the format table.
// FindAll 方法返回 format 表中的所有 Format 记录。如果查询失败，返回错误信息。
func (f *Format) FindAll(rep repository.Repository) (*[]Format, error) {
	var formats []Format
	// 查找所有 Format 记录，存储在 formats 切片中
	if err := rep.Find(&formats).Error; err != nil {
		return nil, err
	}
	// 返回查找到的所有 Format 记录和 nil 错误
	return &formats, nil
}

// Create persists this format data.
// Create 方法将当前的 Format 实例持久化到数据库中，并返回持久化后的对象和可能发生的错误。
func (f *Format) Create(rep repository.Repository) (*Format, error) {
	// 使用 GORM 创建记录，如果发生错误返回 nil 和错误信息
	if err := rep.Create(f).Error; err != nil {
		return nil, err
	}
	// 返回持久化的 Format 对象和 nil 错误
	return f, nil
}

// ToString returns the string representation of the Format object.
// ToString 方法返回当前 Format 对象的字符串表示形式，用于调试或日志记录。
func (f *Format) ToString() string {
	return toString(f) // 调用辅助函数 toString，将结构体对象转为字符串
}
