# Segbot 

![Website](https://img.shields.io/website?label=bot&url=http%3A%2F%2Fmlabouri.tech%3A8000%2Fdiscord)
![GitHub last commit](https://img.shields.io/github/last-commit/Millefeuille42/piscibotReloaded)
[![DeepSource](https://deepsource.io/gh/Millefeuille42/piscibotReloaded.svg/?label=active+issues&show_trend=true)](https://deepsource.io/gh/Millefeuille42/piscibotReloaded/?ref=repository-badge)

The perfect tool to track your progress and compare with your friends or to organize contests for the piscines!

### Features

- Admin: 
    * Auto-Role management
    * Customisable channels
    * Admin list system

- User:
    * Customizable dms and mentions notification settings
    * Track students or pisciners
    * On demand profiles, leaderboards, etc...
    * Real-Time alerts about your target
    * Multi-server support

## How to use

### Getting started

First join a server containing the bot, or 
[add the bot](https://discord.com/api/oauth2/authorize?client_id=816962157841350657&permissions=268823664&scope=bot)
to your own server
<br/>

### Commands

#### Admin

- `init` -> Register the server
- `admin @user (mention)` -> Give user(s) admin privileges, you can set multiple users at once
- `channel <command | leaderboard | success | started | location>` -> Set provided message stream(s) to current channel, you can set multiple streams at once
- `params` -> Get the server's settings
- `purge` -> Delete every message on all the bot's channels
- `lock` -> Lock the registrations, nobody can register anymore
- `unlock` -> Unlock the registrations

#### User

- `start` -> Create your profile, you'll need to verify your account to complete the registration. 
  The profile is independent of the servers
- `track <login>` -> Track provided student, one user per server
- `untrack` -> Untrack current server's student (Automatically set to spectate)
- `spectate` -> Get the spectator role, 
  This role is intended to permit you to see dedicated channels without being tracking someone
- `ping <success | started | location>:<none | dm | mention | all>` -> Edit the notification system for provided message stream(s), you can set multiple streams at once
   - Example -> `ping started:dm` to get dm'd everytime your target starts a project
- `settings` -> Get your notification settings
- `help` -> Send the readme

##### Commands

- `list-<tracked | students | location | projects>` -> Send a list according to the provided parameter. (don't forget the dash!)
  - `tracked` -> Send a list of all the targets.
  - `students` -> Send a list of all the targets, indicating of the target is a student, or not.
  - `location` -> Send a list of all the targets, indicating their current location.
  - `projects` -> Send a list of all the currently available projects for the command `project` (see below).
- `profile <login>` -> Send the profile of the provided target(s). If none provided, sends the profile of your current target.
- `leaderboard <cursus>` -> Send a leaderboard for the provided cursus.
- `project <project>` -> Send the completion status for the given project(s) for all the targets on the server.
- `user-project <login>` -> Send the completion status of the projects for the provided target(s). If none provided, sends the profile of your current target.


## Host your own

You can host your own version of the bot, 
you'll need to set up a discord bot, and a 42 bot.

You'll also need to define the following environment variables : 
- `DB_URL` -> The database url
- `DB_USERNAME` -> The database username of the bot
- `DB_PASSWORD` -> The database password of the bot
- `BOT_TOKEN` -> The discord bot token
- `BOT_DEV_CHANNEL` -> The discord bot default error channel
- `API_PORT` -> The internal 42api port
- `SEGBOT_PORT` -> The internal discord messaging port, should not be conflicting with the previous one
- `APP_HOST` -> The hostname / IP of the server hosting the app
- `APPUID` -> The 42 app UID
- `SECRET` -> The 42 app secret
- `SEGBOT_PREFIX` -> Choose a custom prefix (Defaults to `!` if none provided)

Then, run the app with `docker-compose up --env-file <.env file> -d` and shut it down with `docker-compose down`

## Code overview

![GitHub repo size](https://img.shields.io/github/repo-size/Millefeuille42/piscibotReloaded)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/Millefeuille42/piscibotReloaded)
![Lines of code](https://img.shields.io/tokei/lines/github/Millefeuille42/piscibotReloaded)
![GitHub](https://img.shields.io/github/license/Millefeuille42/piscibotReloaded)

| Segbot                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    	| 42API                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              	| Checker                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       	|
|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------	|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------	|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------	|
| The Discord Bot section, handle commands and manages users.                                                                                                                                                                                                                                                                                                                                                                                                                                               	| Fetches and store students data from the 42 Api.                                                                                                                                                                                                                                                                                                                                                                                                                                                   	| Cycles on tracked students, check differences and report updates to discord.                                                                                                                                                                                                                                                                                                                                                                                                                                  	|
| ![ GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/Millefeuille42/piscibotReloaded?filename=discord%2Fsegbot%2Fgo.mod&label=go%20version)    ![GitHub Workflow Status](https://img.shields.io/github/workflow/status/Millefeuille42/piscibotReloaded/Go%20Build%20Segbot?label=go%20build)  ![GitHub Workflow Status](https://img.shields.io/github/workflow/status/Millefeuille42/piscibotReloaded/Docker%20Build%20Segbot?label=docker%20build)    	| ![ GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/Millefeuille42/piscibotReloaded?filename=42API%2Fapi%2Fgo.mod&label=go%20version)    ![GitHub Workflow Status](https://img.shields.io/github/workflow/status/Millefeuille42/piscibotReloaded/Go%20Build%2042API?label=go%20build)  ![GitHub Workflow Status](https://img.shields.io/github/workflow/status/Millefeuille42/piscibotReloaded/Docker%20Build%2042API?label=docker%20build)    	| ![ GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/Millefeuille42/piscibotReloaded?filename=APICheck%2Fchecker%2Fgo.mod&label=go%20version)    ![GitHub Workflow Status](https://img.shields.io/github/workflow/status/Millefeuille42/piscibotReloaded/Go%20Build%20Checker?label=go%20build)  ![GitHub Workflow Status](https://img.shields.io/github/workflow/status/Millefeuille42/piscibotReloaded/Docker%20Build%20Checker?label=docker%20build)    	|
| ![GitHub last commit](https://img.shields.io/github/last-commit/Millefeuille42/piscibotReloaded/discord)  ![GitHub branch checks state](https://img.shields.io/github/checks-status/Millefeuille42/piscibotReloaded/discord?label=checks)                                                                                                                                                                                                                                                                 	| ![GitHub last commit](https://img.shields.io/github/last-commit/Millefeuille42/piscibotReloaded/42api)  ![GitHub branch checks state](https://img.shields.io/github/checks-status/Millefeuille42/piscibotReloaded/42api?label=checks)                                                                                                                                                                                                                                                              	| ![GitHub last commit](https://img.shields.io/github/last-commit/Millefeuille42/piscibotReloaded/checker)  ![GitHub branch checks state](https://img.shields.io/github/checks-status/Millefeuille42/piscibotReloaded/checker?label=checks)                                                                                                                                                                                                                                                                     	|
|                                                                                                                                                                                                                             [Millefeuille](https://github.com/Millefeuille42)                                                                                                                                                                                                                             	|                                                                                                                                                                                                                            [BoyerDamien](https://github.com/BoyerDamien) and [Millefeuille](https://github.com/Millefeuille42)                                                                                                                                                                        |                                                                                                                                                                                                                                 [BoyerDamien](https://github.com/BoyerDamien) and [Millefeuille](https://github.com/Millefeuille42)                                                                                                                                                                               |
