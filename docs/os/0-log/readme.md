# 日志

## 自定义json格式输出

import CodeBlock from '@theme/CodeBlock';

import maincode from '!!raw-loader!./0-jsonlog/main.go';// 下面必须空一行

<CodeBlock language="go">
{maincode}
</CodeBlock>

## 修改glog默认日志为配置文件日志

```go
// 初始化日志
glog.SetDefaultLogger(g.Log())
glog.SetFlags(glog.GetFlags() | glog.F_FILE_SHORT)
```