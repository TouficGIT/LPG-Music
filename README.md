# LPG-Music
A music bot  for my Discord Server named : **LaPlanqueGaming**
<img src="/LogoBotZic_alpha.png" alt="lpgbot logo" align="right" width="300" height="244"/>
## Current features

 - `$add <ytb-link>`: to add a youtube song (video or playlist)
 - `$current`: display the current music played
 - `$join`: The bot join you VC channel
 - `$leave`: make the bot leave your channel
 - `$pause`: to pause the playlist
 - `$play` : start the playlist
 - `$play <ytb link>`: add the song and play it
 - `$queue`: display all the songs in the queue
 - `$resume`: stop the pause (resume the song)
 - `$stop`: stop the playlist
 - `$skip`: to skip to the next music
 - `$volume <0 to 100>`: to change the volume
 - `$youtube <search>`: search a music on Youtube

 - And you can type `$help` if you need some help

 ## Todo list / Roadmap
- [ ] Add Spotify search
- [ ] Fix the last bug (e.g buffer issue)

## Install
Requirements: Have Go installed (https://golang.org/)
1. Clone repository
2. Run `Go build` in the clone repository
3. Create a bot on your discord developper account
4. Create the `config.json` file and place it inside the repository, following the template below
5. Execute `LPGMusic`

## config.json
You'll need to create a new bot in your Discord developper account.
Then a bot token will be available. Add it to the config.json file
You need to create a '**config.json**' file where you'll save the Token of your bot.
Discord Developper: https://discordapp.com/developers/applications/

The syntax of your config file should be like :

    {
      "Token": "YOUR TOKEN",
      "BotPrefix": "$"
    }


The BotPrefix is the prefix before the bot command. Like "$command".

## Credits
To create this cool bot, I used the [discordGo](https://github.com/bwmarrin/discordgo) and [Airhorn revived](https://github.com/jbmagination/airhornrevived) librairies.

You should go check those amazing git repos ! :+1:

### @Toufic 
