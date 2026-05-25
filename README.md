<div align="center">
<img src="https://github.com/Samenal-sys/derail/blob/master/logo.png" width="200" />

  # derail

### Go Discord Bot
</div>
derail is a Golang tool integrated with Discord and Instagram that scrapes Instagram for videos and embeds them into Discord using vxinstagram.
# Libraries Used

| Library | Usage |
| :--- | :--- |
| chromedp | Used for scraping Instagram. |
| DiscordGo | Used to communicate with Discord. |
| godotenv | Used to store BOT_TOKEN and BOT_CHANNEL. |
| vxinstagram | Used to embed the videos in Discord. |

# Configuration
Create a .env file in the root directory to store your credentials:
```env
BOT_TOKEN=your_bot_token_here
BOT_CHANNEL=your_channel_id_here
```
