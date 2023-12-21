# Wabot - Go version
Developed by: [mendoza000](https://mendoza000.vercel.app)

## What is this?
This is a WhatsApp bot developed in Go, it is a simple bot that can send messages to a specific 
groups, this bot is not finished yet, but it is functional.

## Functions
- Send messages to a specific group list
- Send media
- Customizable messages, captions and media

## How to use
1. Clone this repository
2. Install the dependencies
3. Edit the ./utils/getGroupsToSend.go file and add the groups you want to send messages to
4. Build the project
    ````shell
    go build wabot
   ````
5. Install the binary
    ````shell
   go install wabot
   ````
   
6. Run the binary
    ````shell
    wabot
    ````
7. Select your option
8. Enjoy!