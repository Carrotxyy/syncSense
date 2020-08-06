package models

type Personinfo struct {
	Per_ID             int     `gorm:"column:Per_ID;primary_key" `
	Per_Name           string  `gorm:"column:Per_Name;not null" json:"Per_Name"  validate:"required"`             //人员名称 ·
	Per_OrgID          int     `gorm:"column:Per_OrgID;not null" json:"Per_OrgID"  validate:"required"`           //所属机构
	Per_Sex            string  `gorm:"column:Per_Sex;not null" json:"Per_Sex"  validate:"required"`               //性别,1-男  0-女 ·
	Per_Birth          string  `gorm:"column:Per_Birth" json:"Per_Birth"`                                         //出生日期 ·
	Per_CreateDate     string  `gorm:"column:Per_CreateDate;not null"`                                            //创建日期
	Per_CreateUser     string  `gorm:"column:Per_CreateUser;not null"`                                            //创建用户名
	Per_LastModifyDate string  `gorm:"column:Per_LastModifyDate;not null"`                                        //最后修改日期
	Per_LastModifyUser string  `gorm:"column:Per_LastModifyUser;not null"`                                        //最后修改人姓名
	Per_Status         string  `gorm:"column:Per_Status;not null" json:"Per_Status"  validate:"required"`         //用户状态,'1'：启用 '0'：禁用 ·
	Per_ContactTel     string  `gorm:"column:Per_ContactTel;not null" json:"Per_ContactTel"  validate:"required"` //联系电话 ·
	Per_Mobile         string  `gorm:"column:Per_Mobile;not null" json:"Per_Mobile"  validate:"required"`         //移动电话
	Per_Image          string  `gorm:"column:Per_Image;not null" json:"Per_Image"  validate:"required"`           //照片
	Per_Email          string  `gorm:"column:Per_Email;not null" json:"Per_Email"  validate:"required"`           //电子邮箱
	Per_PeakPerID      string  `gorm:"column:Per_PeakPerID"`                                                      //披克系统ID
	Per_SensePerID     string  `gorm:"column:Per_SensePerID"`                                                     //商汤系统ID
	Per_MegPerID       string  `gorm:"column:Per_MegPerID"`                                                       //旷视系统ID
	Per_AllowVisit     string  `gorm:"column:Per_AllowVisit;not null" json:"Per_AllowVisit"  validate:"required"` //允许接收访问 '0':启用 '1':禁用
	Per_Position       string  `gorm:"column:Per_Position" json:"Per_Position"`                                   //职位
	Per_EmpDate        string  `gorm:"column:Per_EmpDate" json:"Per_EmpDate"`                                     //进入公司日期
	Per_WorkID         string  `gorm:"column:Per_WorkID" json:"Per_WorkID"`                                       //公司工号
	Per_CardID1        string  `gorm:"column:Per_CardID1"`                                                        //卡号1
	Per_CardID2        string  `gorm:"column:Per_CardID2"`                                                        //卡号2
	Per_CardID3        string  `gorm:"column:Per_CardID3"`                                                        //卡号3
	Per_SenseMark      string  `gorm:"column:Per_SenseMark"`                                                      // 商汤系统同步标志位 0：已同步 1：新增 2：修改 3：删除
	Per_OrgRefer       Orginfo `gorm:"foreignkey:Per_OrgID;association_foreignkey:Org_ID"`
}
