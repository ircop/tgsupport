# tgsupport: Telegram bot for organizing support chat written with golang.

Usage: ./tgsupport -c /path/to/tgsupport.toml

## Config parameters:
token: telegram bot token
log: path to logfile. Should be writeable to bot user.
wh-url: web url, on which Telegram backend will send updates. Every time on startup soft will set this webhook url on telegram api.
listen: listen addr.
support-chat: Telegram chat ID where bot will forward user's messages. Bot should be Admin in this channel.

