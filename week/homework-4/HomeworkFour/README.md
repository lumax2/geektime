# 集成单元测试小案例

按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

## 通过swagger请求触发
```
1 调用查询评价模板接口
https://test102.***.cn/comment/base/questionnaire/model/query?MODELID=5
2 根据1响应调用提交评价接口
https://test102.***.cn/comment/base/questionnaire/submit
3 响应断言
```
## 项目采用kratos框架
```
https://go-kratos.dev/docs/
```
## 项目启动
```
kratos run
```
## 项目请求swagger地址
```
http://localhost:8000/q/swagger-ui
```