package actions

import "github.com/bwmarrin/discordgo"

func Clown(s *discordgo.Session, m *discordgo.MessageCreate) {
	go s.UpdateGameStatus(1, "The circus has arrived")
	//find a way to disable max discord messages
	go func() {
		for i := 0; i < 20; i++ {
			s.ChannelMessageSend(m.ChannelID, "🎪Clowned🎪")
		}
	}()
	go clownmusic(s, m)

	//TODO: Use complete for new func
	<-complete

	//after that add circus noise
}
