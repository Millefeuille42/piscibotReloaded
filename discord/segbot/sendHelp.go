package main

func sendHelp(agent discordAgent) {
	sendMessageWithMention("#### Admin\n\n"+
		"- `!init` -> Register the server\n"+
		"- `!admin @user (mention)` -> Give user(s) admin privileges, you can set multiple users at once\n"+
		"- `!channel <command | leaderboard | success | started | location>` -> Set provided message stream(s) to current channel, you can set multiple streams at once\n"+
		"- `!params` -> Get the server's settings\n\n"+
		"#### User\n\n"+
		"- `!start` -> Create your profile, you'll need to verify your account to complete the registration. \n"+
		"The profile is independent of the servers\n- `!track <login>` -> Track provided student, one user per server\n"+
		"- `!untrack` -> Untrack current server's student (Automatically set to spectate)\n"+
		"- `!spectate` -> Get the spectator role,\nThis role is intended to permit you to see dedicated channels without being tracking someone\n"+
		"- `!ping <success | started | location>:<none | dm | channel | all>` -> Edit the notification system for provided message stream(s), you can set multiple streams at once\n"+
		"- `!settings` -> Get your notification settings", "", agent)
}
