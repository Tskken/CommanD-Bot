 # **CommanD-Bot**
A sudo command line for full control over your Discord server.  This bot is powered by [Go][1], [Discordgo][2], and [jbrukh/bayesian][3] library's.  Please support them, this bot would not be possible without them.

[1]: https://golang.org/
[2]: https://github.com/bwmarrin/discordgo
[3]: https://github.com/jbrukh/bayesian

### How to install/run
#### Note: Current version must be compiled as no exe is given.  Install Go and compile to run on your own computer. Later verisons will be easier to run.  This code is only currently in a beta state for testing.

[3]: https://discordapp.com/oauth2/authorize?client_id=357950177945976839&scope=bot&permissions=8

---

## Table Of Contents
+ [What is CommanD-Bot?](#Introduction)
+ [How to use](#HowTo)
+ [Primary Commands](#Commands)
    - [!help](#help)
    - [!message](#message)
    - [!player](#player)
    - [!channel](#channel)
+ [Sub Commands](#sub_commands)
    - [message commands](#subMessage)
    - [player commands](#subPlayer)
    - [channel commands](#subChannel)
    - [utility commands](#subUtility)
+ [Whats Planned?](#planned)

### Introduction <a id="Introduction"></a>
CommanD-Bot was created with a specific goal. Give Discord server moderators the tools they need to make there job easier. Commands to help moderate chat and servers along with a spam and bad word filter to keep the chat civil and salt free.

### How to use <a id="HowTo"></a>
All commands require a `!`.  The bot relays on a main command and sub command argument pair. This takes the shape of something like `!message -delete`. This was to give the ability to have multiple command functions which can differ based on the first command given.
The spam filter automatically scans every message sent with in any channel. It deletes messages based on excessive bad words or spam based on a bayesian learning algorithm. 

### Commands <a id="Commands"></a>
List of all currently supported primary commands

|   Command Name |  Command Function |
| --- | --- |
|   [!help](#help) or [!h](#help) | !help will give info on each command. Called on its own it list each primary command. Called with a primary command as an arugment and it will list that primary commands sub-commands. You can also just enter a primary command with out any arguments to get help on that command. |
|   [!message](#message) or [!ms](#message) |  !message is the wrapper command for all messages with in the server. |
|   [!player](#player) or [!pl](#player) |   !player is the wrapper function for all player based commands with in the server. |
|   [!channel](#channel) or [!ch](#channel) |   !channel gives commands for all channel functions. |
|   [!utility](#utility) or [!util](#utility) |  !utility commands for extra fun stuff with in a server |

### Sub commands <a id="sub_commands"></a>
Commands that can be called by each primary command.

#### !message <a id="subMessage"></a>

|   Command Name |  Arguments |  Command Function |
| --- | --- | --- |
|   -delete or -del | "number of messages" "player name" | Deleted messages from with in the channel the command is sent.  You can pass it any combinations of arguments.  No arguments passed with simply delete the last message sent in the channel.  **Note** Only admin users can delete messages created by other people.  Non admins can only delete messages that they sent. |
|   -clear or -cl |  | This deletes all messages within a channel that can be deleted.  All messages older then two weeks will not be deleted as its is a limitation presented by the Discord API.  A work around may be possible but for the time being not planed. |

#### !player <a id="subPlayer"></a>

|   Command Name |  Arguments |  Command Function |
| --- | --- | --- |
|   -kick or -k |   "username" |    Kicks a user from the guild. |
|   -ban or -b |    "username" |    Bands a user from the guild.  Ban time of the user is 30 days by default.  This can be changed with other commands. |
|   -bantime or -bt |   "number of days" |   Sets the number of days for a user to be band.  If this is not called or changed after startup of the bot the default will be 30 days.  **Note** The time currently set to 30 days for every server the bot is in when the bot starts.  This will be changed in later versions to save the priset times for each server so that the information does not need to be re-entered every time the bot needs to be restarted. |

#### !channel <a id=subChannel></a>

|   Command Name | Arguments | Command Function |
| --- | --- | --- |
| -create or -c |   "name" "type" | Creates a channel with a given name. You can give it a type as `text` or `voice`. If no type is given it is default a text channel. |
| -delete or -del | "name" | Deletes the given channel. |

#### !utility <a id=subUtility><a/>

|   Command Name | Arguments | Command Function |
| --- | --- | --- |
| -dice or -d | "number greater the zero" | Roles a dice between 1 and the entered number. |

### Whats Planned? <a id="planned"><a/>
Check out the CommanD-Bot [Trello][4] page. It gives a decent look at what is planned with in the bot.

[4]:https://trello.com/b/UB3kUIlz/command-bot