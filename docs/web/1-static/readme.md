# 静态文件

import CodeBlock from '@theme/CodeBlock';

## 代码 - main.go

import maincode from '!!raw-loader!./main.go';// 下面必须空一行

<CodeBlock language="go">
{maincode}
</CodeBlock>

## 配置 - config.yaml

import config from '!!raw-loader!./config.yaml';// 下面必须空一行

<CodeBlock language="yaml">
{config}
</CodeBlock>

## 静态文件 - ./static/1.txt

import staticfile from '!!raw-loader!./static/1.txt';// 下面必须空一行

<CodeBlock language="text">
{staticfile}
</CodeBlock>

## 启动

```bash
go run main.go
```

## 访问

<http://127.0.0.1:8080/static/1.txt>
