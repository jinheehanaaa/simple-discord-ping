# Intro
- When you send message "ping" to BOT, Bot will reply with message "pong"

[![Simple Discord Ping Bot](https://img.youtube.com/vi/WmlrXhSOmxg/0.jpg)](https://www.youtube.com/watch?v=WmlrXhSOmxg)

# Setup
- Create module <code>go mod init github.com/jinheehanaaa/simple-discord-ping</code>
- Match dependencies <code>go mod tidy</code>

# Requirement for Bot 
- Bot Token from [Discord Developers](https://discord.com/developers/docs/intro)
- Client ID for OAuth2 (You can get Client ID from [Discord Developers](https://discord.com/developers/docs/intro) as well)
- OAuth2 authentication @ <code >https://discordapp.com/oauth2/authorize?&client_id=CLIENTID&scope=bot </code>

# Main.go
- If configuration is correct, start the bot
- [Empty Struct Channel](https://github.com/jinheehanaaa/simple-discord-ping/blob/master/main.go#L22) for concurrency

# Config.go
- We need to convert from JSON to struct
- [Unmarshall(a,b)](https://github.com/jinheehanaaa/simple-discord-ping/blob/master/config/config.go#L39) => Unmarshall json file "a" & bring the json value into "b" struct type so it can be processed in Golang environment 

# Bot.go
- Create Bot Session 
- Send Message with handler

# 3rd Party Packages
- [DiscordGo](https://github.com/bwmarrin/discordgo) for bindings to the Discord API <code>go get github.com/bwmarrin/discordgo</code>

# Resource
- [Playlist](https://www.youtube.com/watch?v=r4uYpshNVyw&list=PL5dTjWUk_cPZwsRo2ZPtgp3KejezGQeae)
- [Part 1](https://youtu.be/r4uYpshNVyw)
- [Part 2](https://youtu.be/Zurdgn0P7jA)
- [Part 3](https://youtu.be/16nrPcZ9JlY)



