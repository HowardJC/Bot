package actions

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strconv"
)

var complete= make(chan bool)
type voice struct{
	voicechannels map[string]discordgo.ChannelType
}

func (chosenarray voice) channelcollector(channel *discordgo.Channel) {

	chosenarray.voicechannels[channel.ID]=channel.Type
}

func NewCollector() voice{
	c:=new(voice)
	c.voicechannelSetup()
	return *c
}


func (chosenarray *voice) voicechannelSetup(){
	m:=make(map[string]discordgo.ChannelType)
	chosenarray.voicechannels=m
}


func clownmusic(s *discordgo.Session,m *discordgo.MessageCreate){
channel,_:=s.Channel(m.ChannelID)
voicechannels:=NewCollector()
PossibleVoiceChannel,_:= s.GuildChannels(channel.GuildID)
for _, channel := range PossibleVoiceChannel{
	voicechannels.channelcollector(channel)

}
	var chosenchannel int
for x,y := range voicechannels.voicechannels {

	if y == discordgo.ChannelTypeGuildVoice {
		chosenchannel, _ = strconv.Atoi(x)
		break
	}
}

	if chosenchannel==0{
		log.Println("Issues with channel ID or magically the general voicechat somehow disappeared")
		return
	}



}




