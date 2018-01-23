package CommanD_Bot

// TODO - Test and modify if necessary

import (
	"os"
	"encoding/gob"
)

var filePath string = "/data/maps.gob"

func SaveGob()error{
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil{
		return err
	}
	encoder := gob.NewEncoder(file)
	encoder.Encode(BotCommands)
	encoder.Encode(BanTime)
	return nil
}

func LoadGob()error{
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		file.Close()
		return err
	}
	decoder := gob.NewDecoder(file)
	decoder.Decode(BotCommands)
	decoder.Decode(BanTime)
	return nil

}