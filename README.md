# **CommanD-Bot**
A sudo command line for full control over your Discord server.  This bot is powered by [Go][1] and the [Discordgo][2] library.  Please support them, this bot would not be possible with out them.

[1]: https://golang.org/
[2]: https://github.com/bwmarrin/discordgo

### How to install/run
1. Download the zip file to your pc.
2. Go to [this link][3] and add the bot to your server
3. Go to the location of the zip file that you downloaded and extract it to a location on your pc.
4. Run the CommanD.exe to start the bot.
5. Fallow the instructions in the command line for initial startup. (these will only show up the first time you run the exe).

[3]: https://discordapp.com/oauth2/authorize?client_id=357950177945976839&scope=bot&permissions=0

---

# NOTE!!!: 
##This is not a correct and up to date README.  All example commands are subject to change along with any given library's and code design.

---

## Table Of Contents
+ [Introduction](#Introduction)
+ [Initial Setup](#Initial_Setup)
+ [Commands](#Commands)
  - [Admin Commands](#Admin_Commands)
  - [User Commands](#User_Commands)

### Introduction <a id="Introduction"></a>
After using a few other bots that give general control over a server.  I always found the bot ether had issues with speed, consistency, or would crash periodically.  I also found that most bots had really pore support when it comes to usability.  Most require a decent understanding of how to setup a bot with a server and give very little info on how to setup the bot for new people.  My goal is to make a bot that can both be setup and run quickly and efficiently.

A side effect of these ideas was the desire to work in Go.  Go is a relatively new programing language made by Google.  For more info on what the language can do please visit there website at [golang.org](https://golang.org/).  Go gives a lot of support similar to Python and Java, but is a compiled language like C.  It has decent speeds even with garbage collection and gives you the ability to have dynamic typing.  In general the language seemed like an interesting use for a Discord bot.  Python is one of the more common ways to program a Discord bot but it is notorious for being clunky when running.  Given my goal to have a fast and efficient bot, Python would not be my best option.

The main purpose of this bot is to give all the support you could want for servers of any size.  This comes in the way of commands which give you the ability to control the flow of your server.  Create new channels, delete messages, and permission control over the server.

### Initial Setup <a id="Initial_Setup"></a>
On first startup there will be some requirements that have to be filled when running the .exe file.  You will be able to either create your own announcement and bot channel or have it auto create them.  You will need to give the bot your username and # as it will give you a bot role which grants you permission to use the admin commands.  On first startup in a server the bot will create a role for itself.  This will grant it full admin roles so that it can do all of its functions.  It will also create a admin role so that who ever has this role will be able to use admin commands.  Give this role to who ever you want to have access to the admin commands.  **Note**:  These commands give full control over your server, don't give this role to some one you do not trust.  The bot will also create a text channel for the bot specifically.  This channel will only visible to admin users and is the location to be able to use admin commands.  You will only be able to use some admin commands outside of this text channel for security/privacy.  The channel is basically your own command console for the server, and will be treated as such by the bot.  The bot will ask if you wish to add an announcement and rules.  If you skip this step there will be no rules or announcement set but there will be a rules and an announcement channel created.  You will be able to set these two things later through bot commands.

### Commands <a id="Commands"></a>
Initial list of commands the bot will support

**NOTE: Not implemented yet.**  
**TODO**:  Finish initial command listing.  Implement basic functionality.

###### Admin Commands <a id="Admin_Commands"></a>

|   Command Name   |   Command Function   |
|---|---|
|   ~help   |   List of all commands an Admin of the server can use   |
|   ~help <*command name*>   |   Help for the given command   |
|   ~setRules <*Rule Name*> <*Rules Text*>   |   Set a rule for the server with the rule name and text   |
|   ~setAnnouncement <*Announcement Header*> <*Announcement Text*>   |   Set the current announcement of the server.  Will replace what ever announcement was prior   |
|   ~deleteMessage |   Delete the last message sent in the channel this is called   |
|   ~deleteMessage <*Number from 1 - N*>   |   Delete the last N messages sent in a channel   |
|   ~deleteMessage <*User Name*> <*Number fro 1 - N*>   |   Deletes the last N messages sent by a given user in the entered channel   |
|   ~clearChannel <*channel name*>   |   Will delete all messages from with in the given text channel. **WARNING** do not use this on a long existing channel, it may take a vary long time or even crash the bot   |
|  ~~~muteMessage <*Message ID*>~~  |   ~~Will change what ever the text of the message was to all *~~   |
|   ~muteMember <*UserID*>   |   Will delete the last 10 messages from a given member in all channels.  User will not be able to enter messages for a default minute (Length can be changed in settings). **Note**: My be slow depending on the size of the server|
|   ~muteMember <*UserID*> <*Time*>   |   Will mute the member for given time **Note**: Overrides default set with in server.  Only for this user once.  If called again with out given time, will go with default.   |
|   ~setMuteTimer <*Time*>   |   Will set the default duration of the ~muteMember command **Note**: If never set bot default is 1 minute   |
|   ~createChannel <*Channel name*>   |   Creates a new permanent text channel with the given name    |
|   ~deleteChannel <*Channel name*>   |   Delete the given channel   |
|   ~kickMember <*Member name*>   |   Kicks the given member from the server   |
|   ~banMember <*Member name*>   |   Bands the given member from the server   |

###### User Commands <a id="User_Commands"></a>

|   Command Name   |   Command Function   |
---|---
|   !help   |   List of all commands a User of the server can use   |
|   !help <*Command name*>   |   Help for the given command   |
|   !rules   |   Server rules **Note**: Default set "No rules set at this time."   |
|   !announcements   |   Server announcements **Note**: Default set "No announcements at this time."   |
|   !createChannel <*Channel name*>   |   Creates a new temporary text channel with the given name   |
|   !deleteMessage   |   Deletes the last message the user sent to the channel this was entered in |
|   !deleteMessage <*Number from 1 - 10*>   |   Deletes up to the last 10 messages that user sent in the entered channel   |
