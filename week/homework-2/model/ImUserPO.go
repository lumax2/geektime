package model

/**
 * po对象
 */
type ImUser struct {
	Id          int64
	UserUuid    string
	UserName    string
	Mobile      string
	UserTag     string
	userTagCode string
	isDel       int8
	createTime  string
	createBy    string
	updateTime  string
	updateBy    string
}
