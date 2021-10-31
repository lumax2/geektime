package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"homework/model"
)

var Db *sql.DB

/**
数据库初始化
*/
func openJdbc() {
	var err error
	Db, err = sql.Open("mysql", "xyx:*****@tcp(10.135.106.***:3306)/sh_smcv_maxus_im_sitb")
	if err != nil {
		panic(err)
	}
}

/**
dao层
*/
func QueryUserByUuid(userUuid string) (model.ImUser, error) {
	// 测试连接数据库测试
	openJdbc()
	defer Db.Close()
	// 执行sql
	row := Db.QueryRow("select id,user_uuid,user_name,mobile,user_tag from tm_im_user where user_uuid = ?", userUuid)
	// 映射对象
	var user model.ImUser
	var err = row.Scan(&user.Id, &user.UserUuid, &user.UserName, &user.Mobile, &user.UserTag)
	// 1 将异常上抛 2 通过errors.Wrap方法，包装参数信息
	if err != nil {
		msg := fmt.Sprintf("dao.QueryUserByUuid:%s happened error", userUuid)
		return user, errors.Wrap(err, msg)
	}
	return user, nil
}
