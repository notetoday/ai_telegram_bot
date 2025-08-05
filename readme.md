# ai-anti-bot

Anti-spam robot for [Telegram](https://telegram.org/) groups

### English | [简体中文](wiki/readme_zh.md)

<p align="center">
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-blue" alt="license MIT"></a>
<a href="https://golang.org"><img src="https://img.shields.io/badge/Golang-1.22.3-red" alt="Go version 1.22.3"></a>
<a href="https://github.com/tucnak/telebot"><img src="https://img.shields.io/badge/Telebot Framework-v3-lightgrey" alt="telebot v3"></a>
</p>

`Telegram` is a world-renowned, very convenient and elegant anonymous communication tool.        
However, due to the anonymity of the software, a lot of spam promotion information often appears in the group. We have no way to identify whether there is spam information in the group at all times.      
Fortunately, `Telegram` provides us with a very powerful `Api`. Now we can use `artificial intelligence` to automatically help us detect user behavior.     

If you are a group administrator of `Telegram`, you can directly deploy this project privately.     

If you are a `developer`, you can use this project to familiarize yourself with the interactive development of `Go language` and `Telegram`, so that you can use `Api` to develop your own robot later!


## References：
- Telegram Api: [https://core.telegram.org/bots/api](https://core.telegram.org/bots/api)      
- Telebot: [https://github.com/tucnak/telebot](https://github.com/tucnak/telebot)
- go-openai: [https://github.com/sashabaranov/go-openai](https://github.com/sashabaranov/go-openai)

## How to use

### Docker Compose

```shell
# Clone the project
git clone https://github.com/notetoday/ai-anti-bot.git && cd ai-anti-bot && mkdir data

# Set up your configuration
cp config.yml.example data/config.yml

# Fill in the configuration according to your needs
vi data/config.yml

# Start up the bot
docker compose up -d
```
It's very simple, right?😄

### Local Development

If you prefer to run the bot locally without Docker:

```shell
# Clone the project
git clone https://github.com/notetoday/ai-anti-bot.git && cd ai-anti-bot && mkdir data

# Set up your configuration
cp config.yml.example data/config.yml

# Fill in the configuration according to your needs
vi data/config.yml

# Build the project
go build -o ai-anti-bot cmd/main.go

# Run the bot
./ai-anti-bot
```

## How to configure
```yml
telegram:
  proxy: ""     # Optional: Proxy address for Telegram API, e.g., "http://localhost:8080"
  token: ""     # Required: Your Telegram Bot API Token
  groups: []    # Required: List of group IDs where the bot should be active (e.g., [-123456789, -987654321])
  owners: []    # Required: List of Telegram user IDs of super administrators (e.g., [12345, 67890])
  delete_prompt_messages: true # Optional: Whether to automatically delete bot's prompt messages (e.g., ban/unban hints) after a set time. Defaults to true.

identification_model: "chatgpt" # Required: The AI model to use for identification (e.g., "chatgpt")
clean_bot_message: false # Optional: Whether to periodically clean up bot messages. Defaults to false.

log:
  level: "info" # Optional: Logging level (e.g., "debug", "info", "warn", "error"). Defaults to "info".

retry:
  times: 3 # Optional: Number of retries for failed AI calls. Defaults to 3.
  delay: 5 # Optional: Delay in seconds between retries for AI calls. Defaults to 5.

# Strategy for user verification and threat assessment.
# Example: If a user joined more than 3 days ago AND has spoken more than 3 times OR has been verified once,
# they do not need re-verification. A ban is executed only if the threat score exceeds a certain limit.
# This helps save AI tokens by reducing unnecessary checks.
strategy:
  joined_time: 3        # Optional: Days since joining the group. If user joined longer than this, they are considered "old". Defaults to 3.
  number_of_speeches: 3 # Optional: Number of speeches. If user spoke more than this, they are considered "active". Defaults to 3.
  verification_times: 1 # Optional: Number of verifications. If user verified this many times, they are considered "trusted". Defaults to 1.
  score: 80             # Optional: Threat score threshold (0-100). If AI score exceeds this, action (e.g., ban) is taken. Defaults to 80.

chatgpt:
  proxy: ""   # Optional: Proxy address for OpenAI API, if necessary (e.g., "http://localhost:7890")
  apikey: ""  # Required: Your OpenAI API key
  model: "gpt-4o-mini"   # Required: The detection model to be used (e.g., "gpt-4o-mini"). Note: Models below GPT-4 may not support image and file interaction.

ai:
  timeout: 10 # Optional: AI call timeout in seconds. Defaults to 10 seconds.

# Message templates for bot responses.
message:
  block_hint: "🚫 **Ban Warning** 🚫\n\n[%s](%s) Please note, your username or message contains violations.\n⚠️ Identified as a high-risk user by AI, permanently banned.\n\nRisk Score: %d\n\nReason: ``` %s ```\n\nAI's Sarcastic Remark: ``` %s ```"

# Prompt templates for AI analysis.
prompt:
  user_info: "The user's name is:%s%s, this is the user's %dth message, this user joined the group on %s."
  text: |
    The following are the conditions for judging whether a message is spam or an advertisement:
    1. For new users joining the group (joined less than 1 day ago, spoke less than 3 times), if their message is short, we will judge carefully to reduce the probability of being identified as spam to avoid false bans.
    2. For new users joining the group, if their first few messages have obvious spam characteristics (e.g., using keywords like blockchain, gambling, pornography, finance, or using homophones, typos, homonyms, etc., to circumvent keyword detection, or inserting symbols, special characters, emojis, etc., in the chat content to confuse information), we should correctly identify and judge to avoid false bans.
    3. For existing users in the group (joined the group more than 1 day ago, spoke more than 3 times), if their message is short and has no obvious spam characteristics, we should forcibly determine that their message is not spam to avoid false bans.
    4. If the user's name also has obvious spam characteristics, we should also increase the probability of being judged as spam.

    Spam characteristics examples:
    - Contains fake payment institutions or bank card information, such as fake payment institutions, fake bank card purchases, etc.;
    - Induces users to join groups, click links, or participate in fake activities;
    - Involves illegal payments, gambling, selling prohibited items, and other illegal activities;
    - Provides illegal services, such as代开飞机会员 (opening flight member accounts), 代付 (proxy payment), 刷单 (brushing orders), 赌台 (gambling tables), 出U (selling USDT), 贷款 (loans), 色粉 (pornographic followers), 网赚 (online earning), 交友 (dating), etc.

    Please judge the user's message based on the above information and spam characteristics.

    This is the user's basic information:%s

    The content in double quotes is a message from the user: "%s"

    Based on the above information, is this message spam or promotional information?

    Please return the analysis result in the following JSON structure:
    {"state":<fill in 0 or 1, 1 means it is spam, 0 means it is not>,"spam_score":<fill in a number from 0-100, representing the probability of spam>,"spam_reason":"<judge whether it is spam and provide the reason>","spam_mock_text":"<If identified as spam, please make a sarcastic comment, but be careful not to use any information that might expose the user's identity in the comment. This includes but is not limited to usernames, @handles, and do not retain the information promoted by the advertisement. Also, remember to remind others not to easily believe such information. Comments are limited to 50 characters.>"}
    Please replace the content in angle brackets and directly answer the above JSON object in "plain text", without including any other text.
  image: |
    The following are the conditions for judging whether a message is spam or an advertisement:
    1. For new users joining the group (joined less than 1 day ago, spoke less than 3 times), if their message is short, we will judge carefully to reduce the probability of being identified as spam to avoid false bans.
    2. For new users joining the group, if their first few messages have obvious spam characteristics (e.g., using keywords like blockchain, gambling, pornography, finance, or using homophones, typos, homonyms, etc., to circumvent keyword detection, or inserting symbols, special characters, emojis, etc., in the chat content to confuse information), we should correctly identify and judge to avoid false bans.
    3. For existing users in the group (joined the group more than 1 day ago, spoke more than 3 times), if their message is short and has no obvious spam characteristics, we should forcibly determine that their message is not spam to avoid false bans.
    4. If the user's name also has obvious spam characteristics, we should also increase the probability of being judged as spam.

    Spam characteristics examples:
      - Contains fake payment institutions or bank card information, such as fake payment institutions, fake bank card purchases, etc.;
      - Induces users to join groups, click links, or participate in fake activities;
      - Involves illegal payments, gambling, selling prohibited items, and other illegal activities;
      - Provides illegal services, such as代开飞机会员 (opening flight member accounts), 代付 (proxy payment), 刷单 (brushing orders), 赌台 (gambling tables), 出U (selling USDT), 贷款 (loans), 色粉 (pornographic followers), 网赚 (online earning), 交友 (dating), etc.

    Please judge the user's message based on the above information and spam characteristics.

    This is the user's basic information:%s

    Based on the above information, is the information contained in the image in this message spam or promotional information?

    Please return the analysis result in the following JSON structure:
    {"state":<fill in 0 or 1, 1 means it is spam, 0 means it is not>,"spam_score":<fill in a number from 0-100, representing the probability of spam>,"spam_reason":"<judge whether it is spam and provide the reason>","spam_mock_text":"<If identified as spam, please make a sarcastic comment, but be careful not to use any information that might expose the user's identity in the comment. This includes but is not limited to usernames, @handles, and do not retain the information promoted by the advertisement. Also, remember to remind others not to easily believe such information. Comments are limited to 50 characters.>"}
    Please replace the content in angle brackets and directly answer the above JSON object in "plain text", without including any other text.
```
```

### Other Commands

Here are some other commands you can use with the bot:

```
/start       # Survival detection: The bot will respond if the service is running normally.

# Advertisement Management Commands (Requires telegram.owners to be configured)
# These commands allow super administrators to manage advertising buttons within the bot.

/add_ad      # Add a new advertisement button.
             # Format: /add_ad <title>|<jump_link>|<expiration_time>|<weight>
             # Example: /add_ad Hello|https://google.com|2099-01-01 00:00:00|100
             # - <title>: The text displayed on the button.
             # - <jump_link>: The URL the button links to.
             # - <expiration_time>: When the ad should expire (YYYY-MM-DD HH:MM:SS).
             # - <weight>: A numerical value for sorting (higher value = higher priority).

/all_ad      # View all currently configured advertisement buttons.

/del_ad      # Delete an advertisement button by its ID.
             # Example: /del_ad 1 (deletes the ad with ID 1)
```

## Preview
![preview.png](wiki/preview.png)
