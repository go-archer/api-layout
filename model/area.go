package model

import "github.com/go-zelus/zelus/core/stores/mysql"

type Area struct {
	ID         int64   `db:"id"`
	Level      int8    `db:"level"`       // 层级
	ParentCode int64   `db:"parent_code"` // 父级行政代码
	AreaCode   int64   `db:"area_code"`   // 行政代码
	ZipCode    int     `db:"zip_code"`    // 邮政编码
	CityCode   string  `db:"city_code"`   // 区号
	Name       string  `db:"name"`        // 名称
	ShortName  string  `db:"short_name"`  // 简称
	MergerName string  `db:"merger_name"` // 组合名
	Pinyin     string  `db:"pinyin"`      // 拼音
	Lng        float64 `db:"lng"`         // 经度
	Lat        float64 `db:"lat"`         // 纬度
}

var (
	AreaFields = mysql.FieldNames(&Area{})
)
