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
vi data/config.yml

# 启动机器人
docker compose up -d
```
是不是很简单？😄

### 本地开发

如果您更喜欢在本地运行机器人而不使用 Docker：

```shell
# 克隆项目
git clone https://github.com/notetoday/ai-anti-bot.git && cd ai-anti-bot && mkdir data

# 设置您的配置
cp config.yml.example data/config.yml

# 根据您的需求填写配置
vi data/config.yml

# 构建项目
go build -o ai-anti-bot cmd/main.go

# 运行机器人
./ai-anti-bot
```

## 如何配置
```yml
telegram:
  proxy: ""     # 可选：Telegram API 的代理地址，例如 "http://localhost:8080"
  token: ""     # 必填：您的 Telegram 机器人 API Token
  groups: []    # 必填：机器人需要生效的群组 ID 列表（例如 [-123456789, -987654321]）
  owners: []    # 必填：超级管理员的 Telegram 用户 ID 列表（例如 [12345, 67890]）
  delete_prompt_messages: true # 可选：是否在设定的时间后自动删除机器人发出的提示消息（例如封禁提示、解封提示），默认为 true。

identification_model: "chatgpt" # 必填：用于识别的 AI 模型（例如 "chatgpt"）
clean_bot_message: false # 可选：是否定期清理机器人消息。默认为 false。

log:
  level: "info" # 可选：日志级别（例如 "debug", "info", "warn", "error"）。默认为 "info"。

retry:
  times: 3 # 可选：AI 调用失败的重试次数。默认为 3。
  delay: 5 # 可选：AI 调用重试之间的延迟（秒）。默认为 5。

# 用户验证和威胁评估策略。
# 示例：如果用户加入群组超过 3 天且发言次数超过 3 次，或者已验证一次，则无需再次验证。
# 只有当威胁分数超过一定限制时才执行封禁。
# 这有助于通过减少不必要的检查来节省 AI Token。
strategy:
  joined_time: 3        # 可选：加入群组的天数。如果用户加入时间超过此值，则被视为“老用户”。默认为 3。
  number_of_speeches: 3 # 可选：发言次数。如果用户发言次数超过此值，则被视为“活跃用户”。默认为 3。
  verification_times: 1 # 可选：验证次数。如果用户验证次数达到此值，则被视为“受信任用户”。默认为 1。
  score: 80             # 可选：威胁分数阈值（0-100）。如果 AI 分数超过此值，将采取行动（例如封禁）。默认为 80。

chatgpt:
  proxy: ""   # 可选：OpenAI API 的代理地址，如果需要（例如 "http://localhost:7890"）
  apikey: ""  # 必填：您的 OpenAI API 密钥
  model: "gpt-4o-mini"   # 必填：要使用的检测模型（例如 "gpt-4o-mini"）。注意：GPT-4 以下的模型可能不支持图像和文件交互。

ai:
  timeout: 10 # 可选：AI 调用超时时间（秒）。默认为 10 秒。

# 机器人回复的消息模板。
message:
  block_hint: "🚫 **封禁预警** 🚫\n\n[%s](%s) 请注意，你的用户名或发言存在违规。\n⚠️ 已被 AI 判断为高风险用户，永久封禁。\n\n风险分数：%d\n\n理由：``` %s ```\n\nAI 顺便嘲讽你：``` %s ```"

# AI 分析的提示模板。
prompt:
  user_info: "该用户的名称为:%s%s,这是该用户的第%d次发言，这个用户在%s加入了群组。"
  text: |
    以下是判断发言是否为垃圾广告的条件
    1. 对于新加入群组的用户（加入时间不到1天，发言次数少于3次），如果他的发言较短，我们会谨慎判断，降低识别为垃圾广告的概率，以免错误封禁。
    2. 对于新加入群组的用户，如果他的头几次发布的信息存在很明显的垃圾广告特征（如使用区块链、赌博、色情、金融等关键词，或使用谐音、错别字、同音字等变体来规避关键字检测，或在聊天内容中插入符号、特殊字符、emoji等来混淆信息），我们应当正确识别并判断，以免错误封禁。
    3. 对于群组中已存在的用户（加入群组时间超过1天，发言次数超过3次），如果他的发言字数较短且没有明显垃圾广告特征，我们应强制认定其发言不是垃圾广告，以免错误封禁。
    4. 如果用户的名称中也存在明显的垃圾广告特征，我们也应当提高判定为垃圾广告的概率。

    垃圾广告特征示例:
    - 包含虚假支付机构或银行卡信息，如冒牌支付机构、虚假银行卡购买等；
    - 诱导用户加入群组、点击链接或参与虚假活动;
    - 涉及非法支付、赌博、贩卖禁止物品等违法活动;
    - 提供非法服务，如代开飞机会员、代付、刷单、赌台、出U、贷款、色粉、网赚、交友等。

    请根据以上信息和垃圾广告特征，对用户发言进行判断。

    这是该用户的基本资料:%s

    双引号内的内容是一条来自该用户的发言:"%s"

    根据以上信息，这条发言是垃圾广告或推广信息吗?

    请以以下 JSON 结构返回分析结果:
    {"state":<填写0或1，1表示是垃圾广告，0表示不是>,"spam_score":<填写一个0-100的数字，表示垃圾广告的概率>,"spam_reason":"<判断是否为垃圾广告，并提供原因>","spam_mock_text":"<如果识别为垃圾广告，请进行反讽性的评论，但请注意，在评论中避免使用任何可能暴露用户身份的信息。包括但不限于用户名称、@特号，也不要保留广告所推广的信息。另外，记得提醒其他人不要轻易相信此类信息。评论限制在50字以内>"}
    请替换尖括号中的内容，并以"纯文本"形式直接回答上述的JSON对象，不要包含任何其他的文本。
  image: |
    以下是判断发言是否为垃圾广告的条件
    1. 对于新加入群组的用户（加入时间不到1天，发言次数少于3次），如果他的发言较短，我们会谨慎判断，降低识别为垃圾广告的概率，以免错误封禁。
    2. 对于新加入群组的用户，如果他的头几次发布的信息存在很明显的垃圾广告特征（如使用区块链、赌博、色情、金融等关键词，或使用谐音、错别字、同音字等变体来规避关键字检测，或在聊天内容中插入符号、特殊字符、emoji等来混淆信息），我们应当正确识别并判断，以免错误封禁。
    3. 对于群组中已存在的用户（加入群组时间超过1天，发言次数超过3次），如果他的发言字数较短且没有明显垃圾广告特征，我们应强制认定其发言不是垃圾广告，以免错误封禁。
    4. 如果用户的名称中也存在明显的垃圾广告特征，我们也应当提高判定为垃圾广告的概率。

    垃圾广告特征示例:
      - 包含虚假支付机构或银行卡信息，如冒牌支付机构、虚假银行卡购买等；
      - 诱导用户加入群组、点击链接或参与虚假活动;
      - 涉及非法支付、赌博、贩卖禁止物品等违法活动;
      - 提供非法服务，如代开飞机会员、代付、刷单、赌台、出U、贷款、色粉、网赚、交友等。

    请根据以上信息和垃圾广告特征，对用户发言的图片内容进行判断。

    这是该用户的基本资料:%s

    根据以上信息，这条发言里面图片包含的信息是垃圾广告或推广信息吗?

    请以以下 JSON 结构返回分析结果:
    {"state":<填写0或1，1表示是垃圾广告，0表示不是>,"spam_score":<填写一个0-100的数字，表示垃圾广告的概率>,"spam_reason":"<判断是否为垃圾广告，并提供原因>","spam_mock_text":"<如果识别为垃圾广告，请进行反讽性的评论，但请注意，在评论中避免使用任何可能暴露用户身份的信息。包括但不限于用户名称、@特号，也不要保留广告所推广的信息。另外，记得提醒其他人不要轻易相信此类信息。评论限制在50字以内>"}
    请替换尖括号中的内容，并以"纯文本"形式直接回答上述的JSON对象，不要包含任何其他的文本。
```
```

### 其他命令

以下是您可以与机器人一起使用的其他命令：

```
/start       # 生存检测：如果机器人服务正常运行，将给出反馈。

# 广告管理命令（需要配置 telegram.owners）
# 这些命令允许超级管理员在机器人中管理广告按钮。

/add_ad      # 添加新的广告按钮。
             # 格式：/add_ad <标题>|<跳转链接>|<过期时间>|<权重>
             # 示例：/add_ad Hello|https://google.com|2099-01-01 00:00:00|100
             # - <标题>：按钮上显示的文本。
             # - <跳转链接>：按钮链接到的 URL。
             # - <过期时间>：广告应何时过期（YYYY-MM-DD HH:MM:SS）。
             # - <权重>：用于排序的数值（值越大，优先级越高）。

/all_ad      # 查看所有当前配置的广告按钮。

/del_ad      # 根据 ID 删除广告按钮。
             # 示例：/del_ad 1（删除 ID 为 1 的广告）
```

## 预览
![preview.png](wiki/preview.png)