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

	@where 		查询条件
	@out   		装载条件
	@sel   		查询字段
	@otherSql 	其他查询条件 sql表达式

	返回参数 		(错误，数据数量)
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

/**
	关联查询

	@where 		查询条件
	@out		装载数据
	@refer		需要填充的字段切片
	@sel   		查询字段
	@otherSql 	其他查询条件 sql表达式

	返回参数 (错误，数据数量)
 */
func (b *BaseRepository) GetReferBelongsTo(where,out interface{},refer []string,sel,otherSql string)(error,int){
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

	// 填充字段
	for _, field := range refer {
		db = db.Preload(field)
	}

	err := db.Find(out).Count(&count).Error

	return err,count
}

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

/**
	获取需要上传的访客数据,独立方法,因为要求较多，无法集中优化

	SELECT * FROM
	(SELECT * FROM go_visitor WHERE Vis_SenseMark='1' AND Vis_State='1') AS vis
	INNER JOIN
	(SELECT Per_ID,Per_Name,Per_SensePerID FROM go_personinfo WHERE Per_Status="1" AND Per_SensePerID IS NOT NULL   ) AS per
	ON
	vis.Vis_PerID = per.Per_ID

 */
func (b *BaseRepository) GetUploadVisitor(where , out interface{})(error,int){
	var count int

	db := b.DB.Conn.Where(where)

	// 填充字段
	db = db.Preload("Vis_PersonRefer","Per_Status = ? AND Per_AllowVisit = ?","1","0")

	err := db.Find(out).Count(&count).Error

	return err,count
}