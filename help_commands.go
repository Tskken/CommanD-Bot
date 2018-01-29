package CommanD_Bot

import "github.com/bwmarrin/discordgo"

// TODO - REDO !help with a database. (reserch best database for implementation)

type helpInfo struct {
	id string
	args []string
	info string
}

func (h *helpInfo) setID(id string){}
func (h *helpInfo) getId()string{return h.id}

func (h *helpInfo) setArgs(args []string){}
func (h *helpInfo) getArgs()[]string{return h.args}

func (h *helpInfo) setInfo(info string){}
func (h *helpInfo) getInfo()string{return h.info}


func Help(s *discordgo.Session, m *discordgo.Message)error {return nil}