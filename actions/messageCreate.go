package actions

import (
	"github.com/bwmarrin/discordgo"
	"github.com/dlclark/regexp2"
	"regexp"
	"strings"
)

func MessageHandler(s* discordgo.Session, m *discordgo.MessageCreate){
	user,_:=s.User("@me")
	if m.Content[:1]=="?" && m.Author.ID!=user.ID {

		switch strings.ToLower(m.Content[1:]) {
		//stuff to do, actually format my strings using a variable instead of just typing
		case "ping":
			println("Asking for ping")
			s.ChannelMessageSend(m.ChannelID, "I'm still live")

		case "clown":
			println("Asking for a clowning")
			Clown(s,m)
		}

		if match,_:=regexp.MatchString(`^?poll`,m.Content);match {
		println("Asking for a poll")
		Identifier:=regexp2.MustCompile(`(?<=\?poll).*`,0)
		PollPrompt,_:=Identifier.FindStringMatch(m.Content)
		Polling(s,m.Message,PollPrompt.String())
		}
	}}