package actions

import (
	"github.com/bwmarrin/discordgo"
)

func ReactionHandlers(s *discordgo.Session, r *discordgo.MessageReactionAdd){
	if s.State.User.ID==r.UserID {
		return
	}
	switch r.Emoji.Name {
	case "🧑‍🌾":
		go StrawPoll(s,r)
	}


}