package models

type InLogin struct {
	Email    string `json:"email"      form:"email"		 binding:"required"`
	Password string `json:"password"   form:"password"	 binding:"required"`
}

type InUserInfo struct {
	Id int32 `uri:"id" binding:"required"`
}
