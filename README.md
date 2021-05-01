# SlackBot
This is a slackbot application with default responses. This application is developed using Golang, RTM and Docker.

Steps to run:

1) [Create Slackbot](https://$YOUR_ID.slack.com/apps/A0F7YS25R-bots) and copy the token.
2) In application.properties file, add copied token against **SLACK_TOKEN**
3) Run __main.go__ file.

To run as docker image:
1) Run `docker build -t imageName .`
2) Then `docker run imageName`

Now, you can chat with you slackbot.
