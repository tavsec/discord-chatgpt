# ChatGPT Discord Bot
This Golang project is a Discord bot that interacts with ChatGPT, a conversational AI-powered chatbot platform. 

## Requirements
- You need to have access to Discord Development token. You can read more about Discord tokens and applications [here](https://discord.com/developers/docs/intro).
- You will also need OpenAI development key. More about [OpenAI API](https://openai.com/blog/openai-api).

### Discord Token permissions
Your Discord bot needs the following permissions in order to work:
- Send Message

This corresponds to the following [permission integer](https://discord.com/developers/docs/topics/permissions): **2048**

**Make sure that your Discord bot has MESSAGE CONTENT INTENT permission enabled**.



## Setup
First, build Golang program using
```bash
go build -o chatgpt-bot ./...
```

This command will generate executable. You can run the bot using following command:
```bash
DISCORD_AUTH_TOKEN=XXXX OPENAI_AUTH_TOKEN=sk-XXXX ./chatgpt-bot
```

## Usage
To use bot in Discord, simply tag the bot and interact.
