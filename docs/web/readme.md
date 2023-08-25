# web 服务相关

## http 服务端启动流程

版本 v2.5.2

1. 创建 Server 对象（[GetServer](https://github.com/gogf/gf/blob/d72997da0482c9ea39470fbaae03c47132f11037/net/ghttp/ghttp_server.go#L90)）
2. 路由或中间件绑定
3. [s.Start()](https://github.com/gogf/gf/blob/d72997da0482c9ea39470fbaae03c47132f11037/net/ghttp/ghttp_server.go#L122)（非阻塞）或 [s.Run()](https://github.com/gogf/gf/blob/d72997da0482c9ea39470fbaae03c47132f11037/net/ghttp/ghttp_server.go#L420)（阻塞）
4. 接下来是 s.Start()的流程
   1. 判断 s.config.SwaggerPath 是否为空，不为空则注册 swagger 路由
   2. 判断 s.config.OpenApiPath 是否为空，不为空则注册 openapi 路由
   3. s.handlePreBindItems(ctx)注册分组路由
   4. 服务进程配置初始化
   5. 判断服务是否已经被启动过，如果是则返回错误
   6. 判断 s.config.LogPath 是否为空，不为空则初始化日志配置
   7. 判断 s.config.SessionStorage 是否为空，不为空则初始化默认的 session 存储配置
   8. 初始化 Session 管理器
   9. 判断 s.config.PProfEnabled 是否为 true，如果是则开启 pprof 功能以及路由
   10. 判断 s.config.Handler 是否为空，不为空则使用默认的 [Handler](https://github.com/gogf/gf/blob/ea6a773d60fdf964ed3dde6730935167ebbf4fd8/net/ghttp/ghttp_server_handler.go#L31)
   11. 安装扩展插件
   12. s.handlePreBindItems(ctx)注册分组路由，再次注册分组路由，因为上面的步骤可能会注册分组路由
   13. 检查路由表是否为空，如果为空则报错返回
   14. [s.startServer](https://github.com/gogf/gf/blob/d72997da0482c9ea39470fbaae03c47132f11037/net/ghttp/ghttp_server.go#L213) 启动监听
       1. 处理 https ，此时的 s.newGracefulServer 会把 s.config.Handler 写入到服务中
       2. 处理 http ，此时的 s.newGracefulServer 会把 s.config.Handler 写入到服务中
       3. 遍历上面的服务列表，[启动服务](https://github.com/gogf/gf/blob/d72997da0482c9ea39470fbaae03c47132f11037/net/ghttp/ghttp_server.go#L572)
       4. 后续收到的所有请求将会调用 s.config.Handler，默认的 [Handler](https://github.com/gogf/gf/blob/ea6a773d60fdf964ed3dde6730935167ebbf4fd8/net/ghttp/ghttp_server_handler.go#L31)
   15. 日志输出 Swagger 和 OpenApi 信息
   16. 初始化生成路由 OpenApi 对象
   17. 注册服务，用于服务发现
   18. 输出路由表

## ServeHTTP 服务请求处理流程

默认的 [Handler](https://github.com/gogf/gf/blob/ea6a773d60fdf964ed3dde6730935167ebbf4fd8/net/ghttp/ghttp_server_handler.go#L31)

1. 判断 s.config.ClientMaxBodySize 是否大于 0，如果是则设置 r.Body 的最大读取长度
2. 处理或重写请求 URL
3. 创建 ghttp.Request 对象
4. 定义一些延迟调用（收尾处理），例如错误处理、panic 兜底、日志处理、body 关闭等
5. 路由匹配，优先级：静态文件 > 动态路由 > 静态目录
6. HookBeforeServe
7. 动态路由由 [request.Middleware.Next()](https://github.com/gogf/gf/blob/ea6a773d60fdf964ed3dde6730935167ebbf4fd8/net/ghttp/ghttp_server_handler.go#L132) 处理
   1. 循环执行匹配到的所有[handlers](https://github.com/gogf/gf/blob/6eb0de42f817f3433e63bc3d8b30fbdca42bc60a/net/ghttp/ghttp_request_middleware.go#L37)
   2. 循环中包含了 TryCatch 逻辑，如果遇到 panic 则结束循环
      1. 执行中间件
      2. 根据 handler 类型执行
         1. HandlerTypeObject
         2. HandlerTypeHandler
         3. HandlerTypeMiddleware（全局中间件）
      3. 规范路由和普通 handler 会执行 [callHandlerFunc](https://github.com/gogf/gf/blob/6eb0de42f817f3433e63bc3d8b30fbdca42bc60a/net/ghttp/ghttp_request_middleware.go#L129)
         1. m.request.[Parse](https://github.com/gogf/gf/blob/53e5a040734ea6c79085c233ce4ac25880f66c4c/net/ghttp/ghttp_request_param.go#L54)将标准路由的参数传入
         2. r.doGetRequestStruct(pointer)合并了标准路由参数字段的 tag 定义值，如：默认值定义、参数值来源等
   3. 循环结束后根据状态设置 response 的状态码
8. HookAfterServe
9. HookBeforeOutput
10. 判断状态码是否为 0，如果是，有数据返回则设置状态码为 200，有错误则设置状态码为 500，其它情况则设置状态码为 404
11. 自定义状态码处理，如果注册了自定义状态码处理且状态码不为 200，则调用对应的自定义状态码处理
12. 设置 Session
13. 输出 cookie
14. 输出返回数据
15. HookAfterOutput
