package storage

// TODO - Test and modify if necessary

import (
	"os"
	"encoding/gob"
	"path/filepath"
)

var file = "../tsukinai/CommanD-Bot/source/data/data.gob"

type encdec struct {
	file *os.File
	enc *gob.Encoder
	dec *gob.Decoder
}

func NewEncDec()*encdec{
	return &encdec{}
}

func (ed *encdec)OpenFile()error{
	filePath, err := filepath.Abs(file)
	if err != nil {
		return err
	}
	ed.file, err = os.Create(filePath)
	return err
}
func (ed *encdec)CloseFile()error{
	err := ed.file.Close()
	return err
}
func (ed *encdec)getFile()*os.File{
	return ed.file
}

func (ed *encdec)NewEncGob(){
	ed.enc = gob.NewEncoder(ed.getFile())
}
func (ed *encdec)EncGob(val interface{})error{
	err := ed.enc.Encode(val)
	return err
}

func (ed *encdec)NewDecGob(){
	ed.dec = gob.NewDecoder(ed.getFile())
}
func (ed *encdec)DecGob(val interface{})error{
	err := ed.dec.Decode(val)
	return err
}

/*
func SaveGob()error{
	newGob := encdec{}
	err := newGob.openFile()
	if err != nil {
		return err
	}

	defer newGob.closeFile()
	newGob.newEncGob()
	err = newGob.encGob(BotCommands)
	if err != nil {
		return err
	}
	err = newGob.encGob(BanTime)
	return err
}

func LoadGob()error{
	newGob := encdec{}
	err := newGob.openFile()
	if err != nil {
		return err
	}
	newGob.newDecGob()
	newGob.decGob(BotCommands)
	newGob.decGob(BanTime)
	return nil

}*/