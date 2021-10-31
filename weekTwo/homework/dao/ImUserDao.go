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
	Db, err = sql.Open("mysql", "xyx:Pass1234$@tcp(10.135.106.240:3306)/sh_smcv_maxus_im_sitb")
	if err != nil {
		panic(err)
	}
}

/**
dao层
1 查询结果为空，不应该返回error
2 查询异常 err向上抛给业务方
3 查询结果正常，err=nil
*/
func QueryUserByUuid(userUuid string) (model.ImUser, error) {
	// 测试连接数据库测试
	openJdbc()
	defer Db.Close()
	// 执行sql
	row := Db.QueryRow("select id,user_uuid,user_name,from_base64(mobile),user_tag from tm_im_user where user_uuid = ?", userUuid)
	// 映射对象结果
	var user model.ImUser
	var err = row.Scan(&user.Id, &user.UserUuid, &user.UserName, &user.Mobile, &user.UserTag)
	// 结果查询为空也是正常的业务响应，在dao直接处理
	if errors.Is(err, sql.ErrNoRows) {
		return user, fmt.Errorf("dao.QueryUserByUuid:%s 查询为空", userUuid)
	}
	// 查询异常，异常封装信息向上传递
	if err != nil {
		msg := fmt.Sprintf("dao.QueryUserByUuid:%s 查询异常", userUuid)
		return user, errors.Wrapf(err, msg)
	}
	return user, nil
}
