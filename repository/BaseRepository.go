package repository

import "github.com/Carrotxyy/syncSense/common/db"

type BaseRepository struct {
	DB *db.DB `inject:""`
}

/**
	根据反向条件，获取表数据

	@not   排除条件
	@out   装载结果
	@sel   查询的字段

	返回参数 (错误，数据数量)
*/
func (b *BaseRepository) Get_Not(not, out interface{}, sel string) (error, int) {
	var count int
	// 获取数据库对象,Not 反向获取
	db := b.DB.Conn.Not(not)
	if sel != "" {
		// 检索的字段
		db = db.Select(sel)
	}
	err := db.Find(out).Count(&count).Error
	return err, count
}

/**
	根据条件获取表数据

	@where 查询条件
	@out   装载条件
	@sel   查询字段

	返回参数 (错误，数据数量)
 */
func (b *BaseRepository) Get(where,out interface{},sel string,otherSql string)(error,int){
	var count int

	db := b.DB.Conn.Where(where)
	if otherSql != ""{
		// 其他条件
		db = db.Where(otherSql)
	}
	if sel != "" {
		// 检索的字段
		db = db.Select(sel)
	}
	err := db.Find(out).Count(&count).Error
	return err, count
}


//
/**
	更新数据

	@where 更新条件
	@newData 更新数据

	返回参数 (错误)
*/
func (b *BaseRepository) Update(model,where, newData interface{}) error {
	return b.DB.Conn.Model(model).Where(where).Updates(newData).Error
}

/**
	恢复标志位
	@model 表
	@field 根据当前字段，进行筛查
	@fieldValue 字段值的范围
	@value 更新的值

	返回参数 (错误)
 */
func (b *BaseRepository) UpdateMark(model interface{} , field string , fieldValue interface{} , value interface{})error{
	return b.DB.Conn.Model(model).Where(field + " IN (?)",fieldValue).Updates(value).Error
}