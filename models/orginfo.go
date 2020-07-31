package models


type Orginfo struct {
	Org_ID             int    `gorm:"primary_key;column:Org_ID"`                                                 //
	Org_Name           string `gorm:"column:Org_Name;not null" json:"Org_Name" validate:"required"`              // 机构名称 简称
	Org_CreateDate     string `gorm:"column:Org_CreateDate;not null"`                                            // 创建日期
	Org_LastModifyDate string `gorm:"column:Org_LastModifyDate;not null"`                                        //修改日期
	Org_CreateUser     string `gorm:"column:Org_CreateUser;not null"`                                            //创建人
	Org_LastModifyUser string `gorm:"column:Org_LastModifyUser;not null"`                                        //修改人名称
	Org_ParentID       int    `gorm:"column:Org_ParentID"`                                                       //父机构
	Org_FullName       string `gorm:"column:Org_FullName;not null" json:"Org_FullName" validate:"required"`      //机构全称
	Org_Status         string `gorm:"column:Org_Status;not null" json:"Org_Status"`                              //数据状态,1-有效0-无效（机构已删除）
	Org_ContactTel     string `gorm:"column:Org_ContactTel;not null" json:"Org_ContactTel"  validate:"required"` //联系电话
	Org_Leader         string `gorm:"column:Org_Leader;not null" json:"Org_Leader" validate:"required"`          //领导/法人
	Org_Contact        string `gorm:"column:Org_Contact;not null" json:"Org_Contact" validate:"required"`        //联系人
	Org_Email          string `gorm:"column:Org_Email;not null" json:"Org_Email" validate:"required" `           //机构邮箱
	Org_ExtentInfo     string `gorm:"column:Org_ExtentInfo" json:"Org_ExtentInfo"`                               //其他邮箱
	Org_Image          string `gorm:"column:Org_Image;not null"`                                                 //地图或形象图（图片的base64）
	Org_PeakID         string `gorm:"column:Org_PeakID" json:"Org_PeakID"`                                       //披克系统机构ID，0非披克用户，9999初创
	Org_SenseID        string `gorm:"column:Org_SenseID" json:"Org_SenseID"`                                     //商汤系统机构ID，0非商汤用户，9999初创
	Org_MegID          string `gorm:"column:Org_MegID" json:"Org_MegID"`                                         //旷视系统机构ID，0非旷视用户，9999初创
	Org_SenseMark      string `gorm:"column:Org_SenseMark" json:"Org_SenseMark"`                                 //同步商汤系统标志位 0：已同步 1：新增 2：修改 3：删除
}
