# slack-2fa-check
Command line utility to make sure all members of a Slack team have 2FA enabled. You will need a token from https://api.slack.com/web#authentication for your team to make this work. The tool is appropriate for use as a nagios plugin right out of the box (exit code 2 when it finds any non-2fa users). The tool ignores deleted users and bot users since niether are allowed to log into your team.

# obtaining
go get github.com/apokalyptik/slack-2fa-check

# running
path/to/slack-2fa-check -token=xxx
