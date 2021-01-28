package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/dlclark/regexp2"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type strawpollreq struct {
	Title string `json: "title"`
	options []string `json: "options"`
}

type strawpollres struct {
	Id int `json: "id"`
	Title string `json: "title"`
	Answers []string `json: "answers"`
	dupcheck string `json: "dupcheck"`
}



func Polling(s *discordgo.Session, m *discordgo.Message, prompt string) {
	Poll:=  fmt.Sprintf("Starting poll: %s",prompt)
	CurrentPoll,_:=s.ChannelMessageSend(m.ChannelID,Poll)

	go s.MessageReactionAdd(CurrentPoll.ChannelID,CurrentPoll.ID,"⬆️")
	go s.MessageReactionAdd(CurrentPoll.ChannelID,CurrentPoll.ID,"⬇️")
	go s.MessageReactionAdd(CurrentPoll.ChannelID,CurrentPoll.ID,"🧑‍🌾")

}

func StrawPoll(s *discordgo.Session,r *discordgo.MessageReactionAdd){
	m, err := s.ChannelMessage(r.ChannelID,r.MessageID)
	if err != nil{
		println("Message does not exist? Possibly clown bot user raider deleting too fast?")
		return
	}

	Identifier:=regexp2.MustCompile(`(?<=\Starting poll:).*`,0)
	PollPrompt,_:=Identifier.FindStringMatch(m.Content)
	println(PollPrompt.String())

	request:=&strawpollreq{
		PollPrompt.String(),
		[]string{"Up","Down"},
	}
		buffer,_:=json.Marshal(request)
		req, err := http.NewRequest("POST", "https://strawpoll.com/api/poll", bytes.NewBuffer(buffer))
				req.Header.Set(`Content-Type`,  `application/json`,)

		client:=&http.Client{Timeout:time.Second*50}
		resp,err:=client.Do(req)

		body,_:=ioutil.ReadAll(resp.Body)

		var strawres strawpollres
		err = json.Unmarshal(body,&strawres)
		if err!=nil{
			log.Print("Failed to recieve")
			return
		}
		s.ChannelMessageSend(m.ChannelID, "www.strawpoll.me/"+strconv.Itoa(strawres.Id))



}