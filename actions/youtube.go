package actions

import "github.com/bwmarrin/discordgo"

var complete= make(chan bool)
type voice struct{
	voicechannels map[string]string
}

func (chosenarray voice) channelcollector(channel *discordgo.Channel) {
	chosenarray.voicechannels[channel.ID]=string(channel.Type)
}


func clownmusic(s *discordgo.Session,m *discordgo.MessageCreate){
channel,_:=s.Channel(m.ChannelID)
voicechannels:=new(voice)
PossibleVoiceChannel,_:= s.GuildChannels(channel.GuildID)
for _, channel := range PossibleVoiceChannel{
	voicechannels.channelcollector(channel)

}
}
