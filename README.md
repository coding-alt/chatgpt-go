# Chat App

基于 Go 语言实现的简单聊天应用，可以与后端 API 进行一问一答式的对话。
支持配置代理模式访问。
支持配置 stream 模式，实现实时响应的逐字打印效果，提升体验。

## 配置

在 `config.json` 文件中，配置以下参数：

- `PROXY_OPTION`: 代理选项。例如，可以配置为 "DIRECT"（无代理）或 "SOCKS5 127.0.0.1:5000"（SOCKS5 代理）或 "HTTP 127.0.0.1:5000"（HTTP 代理）
- `APIKEY`: API 密钥
- `API_URL`: API URL，例如 "https://api.openai.com/v1/completions"
- `MODEL`: 模型名称，例如 "text-davinci-003"
- `API_TIMEOUT`: API 超时时间（以秒为单位）

## 编译

在项目根目录下运行以下命令以编译程序：

```sh
go build -o chat_app
```

## 运行

在项目根目录下运行以下命令以启动程序：

```
./chat_app
```

按照提示输入问题，程序将从 API 获取答案并实时显示。要退出程序，请输入 "quit"。

## 示例
```
$ ./chat_app
请输入问题，或输入 'quit' 退出：
What is Python?
答案： Python is a high-level, interpreted, general-purpose programming language. It is a powerful and versatile language that is used for a wide range of applications, from web development to data science. Python is known for its readability and ease of use, making it a great language for beginners. It also has a large and active community of developers who contribute to the language and its libraries.
请输入问题，或输入 'quit' 退出：
quit
```
