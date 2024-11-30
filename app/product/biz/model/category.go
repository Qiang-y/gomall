package model

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
)

type Category struct {
	Base
	Name        string `json:"name"`
	Description string `json:"description"`

	Products []Product `json:"product" gorm:"many2many:product_category"`
}

func (Category) TableName() string {
	return "category"
}

type CategoryQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func (c CategoryQuery) GetProductByCategoryName(name string) (categories []Category, err error) {
	err = c.db.WithContext(c.ctx).Model(&Category{}).Where(&Category{Name: name}).Preload("Products").Find(&categories).Error
	klog.Infof("categories: %v have %v", categories[0].Name, len(categories[0].Products))
	return
}

func NewCategoryQuery(ctx context.Context, db *gorm.DB) *CategoryQuery {
	return &CategoryQuery{
		ctx: ctx,
		db:  db,
	}
}
