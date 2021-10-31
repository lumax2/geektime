package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"homework/dao"
)

/**
biz层
1 b2cf40729c9f4e128cfd1172020e4fe7111 ：query result is null
2 正常响应 b2cf40729c9f4e128cfd1172020e4fe7 ：query customer :  {1902 b2cf40729c9f4e128cfd1172020e4fe7 程浩 MTU5MDE5NDUyNTM= 大通客户  0    }
3 异常上抛情况-> 在 if err！=nil 中可以看到相关信息
query customer err : sql: no rows in result set
dao.QueryUserByUuid:b2cf40729c9f4e128cfd1172020e4fe71 happened error
homework/dao.QueryUserByUuid
        /Users/chenghao/GolandProjects/geekweek-go/geektime/weekTwo/homework/dao/ImUserDao.go:39

=》》建议使用方式，1err层层上抛，2 err上抛前最好只包装一次 3 errors.is() errors.AS() 按照业务逻辑定义处理
*/
func main() {
	user, err := dao.QueryUserByUuid("b2cf40729c9f4e128cfd1172020e4fe71")
	// 查询结果为空，正常响应
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("query result is null")
		return
	}
	// 异常情况查询为空 panic
	if err != nil {
		fmt.Printf("query customer err : %+v", err)
		return
	}

	fmt.Println("query customer : ", user)

}
