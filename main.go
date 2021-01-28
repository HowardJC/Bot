package main

import (
	"Bot/actions"
	"fmt"
	_ "fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	discord, err := discordgo.New("Bot " + config.token)
	discord.AddHandler(actions.MessageHandler)
	discord.AddHandler(actions.ReactionHandlers)
	if err!=nil{
		panic("Panic")
	}
	if err:=discord.Open(); err!=nil{
		panic("Error connecting")
	} else{
		fmt.Print("Connected")
	}


	lock:=make(chan os.Signal)
	signal.Notify(lock, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-lock
	fmt.Print("Closing")


}
