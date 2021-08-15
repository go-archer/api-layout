package entity

type Area struct {
	AreaCode   int64   `json:"area_code"`   // 行政代码
	ParentCode int64   `json:"parent_code"` // 父级行政代码
	Level      int8    `json:"level"`       // 层级
	Name       string  `json:"name"`        // 名称
	ShortName  string  `json:"short_name"`  // 简称
	MergerName string  `json:"MergerName"`  // 组合名
	Pinyin     string  `json:"Pinyin"`      // 拼音
	ZipCode    int     `json:"zip_code"`    // 邮政编码
	CityCode   string  `json:"city_code"`   // 区号
	Lng        float64 `json:"lng"`         // 经度
	Lat        float64 `json:"lat"`         // 纬度
}
