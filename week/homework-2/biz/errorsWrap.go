package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"homework/dao"
	"homework/model"
)

/**
biz层
知识总结：
1 建议使用方式，1err层层上抛，2 err上抛前最好只包装一次 3 errors.is() errors.AS() 按照业务逻辑定义处理
*/
func main() {
	info, _ := queryUserInfo("b2cf40729c9f4e128cfd1172020e4fe71")
	fmt.Printf("用户信息：%v", info)

}

/**
业务层代码
修改用户姓名
 dao 层将查询结果为空的情况，按照error形式返回，业务层自行处理
*/
func queryUserInfo(userUuid string) (model.ImUser, error) {
	user, err := dao.QueryUserByUuid(userUuid)
	// 1 查询结果为空，正常响应
	if errors.Is(err, sql.ErrNoRows) {
		return user, err
	}
	// 2 异常向上抛出
	if err != nil {
		fmt.Printf("query customer err : %+v \n", err)
		return user, err
	}
	return user, err
}
