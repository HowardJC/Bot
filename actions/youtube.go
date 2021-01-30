package actions

import (
	"fmt"
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"strconv"
)

var complete = make(chan bool)

type voice struct {
	voicechannels map[string]discordgo.ChannelType
	connection    *discordgo.VoiceConnection
}

func (chosenarray voice) channelcollector(channel *discordgo.Channel) {

	chosenarray.voicechannels[channel.ID] = channel.Type
}

func NewCollector() voice {
	c := new(voice)
	c.voicechannelSetup()
	return *c
}

func (chosenarray *voice) voicechannelSetup() {
	m := make(map[string]discordgo.ChannelType)
	chosenarray.voicechannels = m
}

func (Speaking voice) playmusic(connection *discordgo.VoiceConnection) {
	files, err := ioutil.ReadDir("./music")
	if err != nil {
		println("Error Opening file")
		return
	}
	for _, f := range files {
		dgvoice.PlayAudioFile(connection, fmt.Sprintf("./music/%s", f.Name()), make(chan bool))
	}

	//TODO:Finish this

	//	dgvoice.PlayAudioFile(connection,fmt.Sprintf("%s/%s",*Folder,f.Name()))
}

//Remind self to just create struct for server and channel id
func clownmusic(s *discordgo.Session, m *discordgo.MessageCreate) {
	channel, _ := s.Channel(m.ChannelID)
	voicechannels := NewCollector()
	PossibleVoiceChannel, _ := s.GuildChannels(channel.GuildID)
	for _, channel := range PossibleVoiceChannel {
		voicechannels.channelcollector(channel)

	}
	var chosenchannel int
	for x, y := range voicechannels.voicechannels {
		if y == discordgo.ChannelTypeGuildVoice {
			chosenchannel, _ = strconv.Atoi(x)
			break
		}
	}

	if chosenchannel == 0 {
		log.Println("Issues with channel ID or magically the general voicechat somehow disappeared")
		return
	}
	connection, _ := s.ChannelVoiceJoin(channel.GuildID, strconv.Itoa(chosenchannel), false, true)
	voicechannels.connection = connection

	go println("Finished joining")
	go voicechannels.playmusic(connection)
	//select{
	//case
	//}
}
