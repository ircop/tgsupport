# tgsupport: Telegram bot for organizing support chat written with golang.

Usage: ./tgsupport -c /path/to/tgsupport.toml

## Description

This bot is forwarding private messages from users to your 'support chat', which id is defined in config.

To answer on user messages, just reply them in chat.

From users: any messages are forwarding.

From support: only text messages are copied to users.

## Config parameters:

token: telegram bot token

log: path to logfile. Should be writeable to bot user.

wh-url: web url, on which Telegram backend will send updates. Every time on startup soft will set this webhook url on telegram api.

listen: listen addr.

support-chat: Telegram chat ID where bot will forward user's messages. Bot should be Admin in this channel.

