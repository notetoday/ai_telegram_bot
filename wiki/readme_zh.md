# ai-anti-bot

Telegram 群组反垃圾机器人

### English | [简体中文](wiki/readme_zh.md)

<p align="center">
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-blue" alt="license MIT"></a>
<a href="https://golang.org"><img src="https://img.shields.io/badge/Golang-1.22.3-red" alt="Go version 1.22.3"></a>
<a href="https://github.com/tucnak/telebot"><img src="https://img.shields.io/badge/Telebot Framework-v3-lightgrey" alt="telebot v3"></a>
</p>

`Telegram` 是一个享誉全球，非常便捷优雅的匿名通讯工具。
然而，由于软件的匿名性，群组中经常出现大量的垃圾推广信息。我们无法时刻识别群组中是否存在垃圾信息。
幸运的是，`Telegram` 为我们提供了非常强大的 `Api`。现在我们可以使用 `人工智能` 来自动帮助我们检测用户行为。

如果您是 `Telegram` 的群组管理员，您可以直接私有部署本项目。

如果您是 `开发者`，您可以使用本项目来熟悉 `Go语言` 和 `Telegram` 的交互式开发，以便您以后可以使用 `Api` 开发自己的机器人！

## 参考：
- Telegram Api: [https://core.telegram.org/bots/api](https://core.telegram.org/bots/api)
- Telebot: [https://github.com/tucnak/telebot](https://github.com/tucnak/telebot)
- go-openai: [https://github.com/sashabaranov/go-openai](https://github.com/sashabaranov/go-openai)

## 如何使用

### Docker Compose

```shell
# 克隆项目
git clone https://github.com/notetoday/ai-anti-bot.git && cd ai-anti-bot && mkdir data

# 设置您的配置
cp config.yml.example data/config.yml

# 根据您的需求填写配置
vi config.yml

# 启动
docker compose up -d
```
是不是很简单？😄

## 如何配置
```yml
telegram:
  proxy:
  token: ""     # 填写您的机器人Token
  groups: [""]  # 填写机器人需要生效的群组ID
  owners: [""]  # 填写超级管理员的telegram用户ID
identification_model: "chatgpt"
clean_bot_message: true # 定时清理机器人消息

# 以下是判断策略。例如，
# 以下配置项的含义是：如果加入群组超过3天且发言次数超过3次或已验证一次，则无需再次验证。
# 或者威胁分数必须超过一定限制才能执行封禁
# 这是为了节省您的Token😊
strategy:
  joined_time: 3
  number_of_speeches: 3
  verification_times: 1
  score: 80


chatgpt:
  proxy: ""   # OpenAI的代理地址，如果需要
  apikey: ""  # apikey
  model: "gpt-4o-mini"   # 要使用的检测模型。请注意，gpt4以下版本不支持图片和文件交互。

ai:
  timeout: 10 # AI调用超时，单位秒。默认为10秒。

# 如果您的母语不是中文而是其他语言，
# 请使用翻译替换以下提示为您想要的语言。
prompt:
   ...
```

### 其他命令
```
/start       # 生存检测：如果机器人服务正常运行，将给出反馈

# 我们还可以使用以下命令向机器人添加自己的广告按钮，但请务必配置telegram.owners选项

/add_ad     # 添加新广告，格式：广告标题|跳转链接|过期时间（以时，分，秒为单位）|权重（按降序排列，数值越大，权重越高），例如：/add_ad Hello|https://google.com|2099-01-01 00:00:00|100

/all_ad     # 查看所有广告按钮

/del_ad     # 删除广告按钮，例如：/del_ad 1（删除id为1的广告）
```

## 预览
![preview.png](wiki/preview.png)