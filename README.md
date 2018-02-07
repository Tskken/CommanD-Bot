 # **CommanD-Bot**
A sudo command line for full control over your Discord server.  This bot is powered by [Go][1] and the [Discordgo][2] library.  Please support them, this bot would not be possible without them.

[1]: https://golang.org/
[2]: https://github.com/bwmarrin/discordgo

### How to install/run
1. Go to [this link][3] and add the bot to your server
2. That's it :)

[3]: https://discordapp.com/oauth2/authorize?client_id=357950177945976839&scope=bot&permissions=1

---

# NOTE!!!
THIS README IS SUBJECT TO CHANGE.  ALL GIVEN FUNCTIONALITY COULD BE CHANGED OR REMOVED IF IT IS DEEMED APPROPRIATE

---

## Table Of Contents
+ [Introduction](#Introduction)
+ [<s>Initial Setup</s>](#Initial_Setup) Out of date
+ [Commands](#Commands)
+ [Sub Commands](#sub_commands)
   - [!message](#message)
   - [!player](#player)
+ [Planned Commands](#planned_commands)

### Introduction <a id="Introduction"></a>
After using a few other bots that give general control over a server.  I always found the bot ether had issues with speed, consistency, or would crash periodically.  I also found that most bot’s had really pore support when it comes to usability.  Most require a decent understanding of how to setup a bot with a server and give very little info on how to setup the bot for new people.  My goal is to make a bot that can both be setup and run quickly and efficiently.

A side effect of these ideas was the desire to work in Go.  Go is a relatively new programing language made by Google.  For more info on what the language can do please visit there website at [golang.org](https://golang.org/).  Go gives a lot of support similar to Python and Java, but is a compiled language like C.  It has decent speeds even with garbage collection and gives you the ability to have dynamic typing.  In general the language seemed like an interesting use for a Discord bot.  Python is one of the more common ways to program a Discord bot but it is notorious for being clunky when running.  Given my goal to have a fast and efficient bot, Python would not be my best option.

The main purpose of this bot is to give all the support you could want for servers of any size.  This comes in the way of commands which give you the ability to control the flow of your server.  Create new channels, delete messages, and permission control over the server.

### Initial Setup <a id="Initial_Setup"></a>
WIP! May be redacted later

### Commands <a id="Commands"></a>
Current supported commands

|   Command Name |  Command Function |
| --- | --- |
|   !help or !h | Gives info on the supported wrapper commands.  Add one of the rappers to !help as an argument and it will give you info on the supported sub commands that command supports. |
|   [!message](#message) or [!ms](#message) |  !message is the wrapper command for all messages with in the server. |
|   [!player](#player) or [!pl](#player) |   !player is the wrapper function for all player based commands with in the server. |

### Sub commands <a id="sub_commands></a>
Currently supported sub commands that fall under there parent wrapper.

#### !message <a id="message"></a>

|   Command Name |  Arguments |  Command Function |
| --- | --- |
|   -delete or -del |   <number of messages> <player name> |    Deleted messages from with in the channel the command is sent.  You can pass it any combinations of arguments.  No arguments passed with simply delete the last message sent in the channel.  **Note** Only admin users can delete messages created by other people.  Non admins can only delete messages that they sent. |
|   -clear or -cl | no arguments    | This deletes all messages within a channel that can be deleted.  All messages older then two weeks will not be deleted as its is a limitation presented by the Discord API.  A work around may be possible but for the time being not planed. |

#### !player <a id="player"></a>

|   Command Name |  Arguments |  Command Function |
| --- | --- |
|   -kick or -k |   <username> |    Kicks a user from the guild. |
|   -ban or -b |    <username> |    Bands a user from the guild.  Ban time of the user is 30 days by default.  This can be changed with other commands. |
|   -bantime or -bt |   <number of days> |   Sets the number of days for a user to be band.  If this is not called or changed after startup of the bot the default will be 30 days.  **Note** The time currently set to 30 days for every server the bot is in when the bot starts.  This will be changed in later versions to save the priset times for each server so that the information does not need to be re-entered every time the bot needs to be restarted. |

### planned commands <a id="planned_commands"></a>
All planned commands are subject to change.  There arguments, naming, and functionality could change based on feedback, limitations of the discord api and discordgo library, limitations of Golang, and any other variables that could change planned development.

|   Command Name |  Arguments |  Command Function |
| --- | --- |
|   !channel -create | <name> <type> | Created a new channel.  The name is mandatory but the type is optional.  If type is not given a text channel will be created by default. |
|   !channel -delete | <name> | Deletes the given channel. **Note** This will only be able to be used by admins unless there is a way to determine the creator of the channel.  If this is possible then non admin users will be able to delete only channels they created. |
|   !utility -dice | <lower bound> <upper bound> | Will role a dice.  The upper bound is mandatory.  The lower bound can be omitted and will default to 0.  If the lower bound is omitted then the dice will role from 0 to the upper bound that is given. |
|   !utility -trinity | <username> <trinity name> | Trinity lets you give "roles" to guild members.  The name is mandatory but the trinity name can be omitted.  If the trinity name is omitted it will by default return the current trinity values given to that user.  If the trinity name is given it will set that trinity name to the given user. |
|   !utility -ign   | <username> <ign> | Ign lets you give a new name to a user outside of there normal discord names.  The user name is mandatory but the ign can be omitted.  If the ign is omitted then it will by default return the current ign's for that user.  If the ign is given it will add that ign to the users list of ign's. |