# geektime
极客时间Go进阶训练营作业提交项目

1.第二周作业：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
以上作业，要求提交到自己的 GitHub 上面，然后把自己的 GitHub 地址填写到班班提供的表单中：
https://jinshuju.net/f/e3bFlA

知识点总结：
1 Go error设计：
1.1 error定义：代码可以在逻辑中处理的，可以用error，是程序员处理各种逻辑用的；
区别于panic，panic是导致程序卡死的，是bug,运行时的程序搞不定的，要bugfix
1.2 errors are value 
1.3 Go支持多参数返回，通过简单的约定，在方法透传error，实现业务异常的处理
1.4 出现error时，立即处理；判断方法：errors.IS AS 找到根因异常 （SentinalError ）
Error Type
2.1 sentinel error :预定义的特定错误（基础库中有大量定义）,用了interface就要对外暴露
2.2 fmt.Errorf() 会导致 == 判定失败
2.3 error types 自定义error：switch type断言 优点：可以获得更多上下文信息 看源码：
2.4 opaque errors 不透明error 
Handling Error
3.1 直线型代码结构
3.2 扫描器模式 sql.row ，注意return row.err
3.3 http writeResponse()方法 通过定义interface 实现if err!=nil 合并
3.4 you should handle error only once! 打日志，抛出来，降级return nil  选一个！


=》errors.Wrap + errors.IS errors.AS
=》fmt.Errorf("msg",err) + errors.IS

二 第三周作业

1. 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。




