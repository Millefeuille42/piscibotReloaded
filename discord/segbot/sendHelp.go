package main

func sendHelp(agent discordAgent) {
	sendMessageWithMention("#### Admin\n\n"+
		"- `init` -> Register the server\n"+
		"- `admin @user (mention)` -> Give user(s) admin privileges, you can set multiple users at once\n"+
		"- `channel <command | leaderboard | success | started | location>` -> Set provided message stream(s) to current channel, you can set multiple streams at once\n"+
		"- `params` -> Get the server's settings\n"+
		"- `purge` -> Delete every message on all the bot's channels\n"+
		"- `lock` -> Lock the registrations, nobody can register anymore\n"+
		"- `unlock` -> Unlock the registrations\n\n", "", agent)
	_ = sendMessageWrapper(agent.session, agent.channel,
		"---\n#### User\n\n"+
			"- `start` -> Create your profile, you'll need to verify your account to complete the registration. \n"+
			"  The profile is independent of the servers\n"+
			"- `track <login>` -> Track provided student, one user per server\n"+
			"- `untrack` -> Untrack current server's student (Automatically set to spectate)\n"+
			"- `spectate` -> Get the spectator role, \n"+
			"  This role is intended to permit you to see dedicated channels without being tracking someone\n"+
			"- `ping <success | started | location>:<none | dm | mention | all>` -> Edit the notification system for provided message stream(s), you can set multiple streams at once\n"+
			"   - Example -> `ping started:dm` to get dm'd everytime your target starts a project\n"+
			"- `settings` -> Get your notification settings\n"+
			"- `help` -> Send the readme\n\n")
	_ = sendMessageWrapper(agent.session, agent.channel,
		"---\n##### Commands\n\n"+
			"- `list-<tracked | students | location | projects>` -> Send a list according to the provided parameter. (don't forget the dash!)\n"+
			"  - `tracked` -> Send a list of all the targets.\n"+
			"  - `students` -> Send a list of all the targets, indicating of the target is a student, or not.\n"+
			"  - `location` -> Send a list of all the targets, indicating their current location.\n"+
			"  - `projects` -> Send a list of all the currently available projects for the command `project` (see below).\n"+
			"- `profile <login>` -> Send the profile of the provided target(s). If none provided, sends the profile of your current target.\n"+
			"- `leaderboard <cursus>` -> Send a leaderboard for the provided cursus.\n"+
			"- `project <project>` -> Send the completion status for the given project(s) for all the targets on the server.\n"+
			"- `user-project <login>` -> Send the completion status of the projects for the provided target(s). If none provided, sends the profile of your current target.\n")
}
