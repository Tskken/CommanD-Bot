package storage

import (
	"encoding/gob"
	"github.com/tsukinai/CommanD-Bot/servers"
	"log"
	"os"
	"path/filepath"
)

/*
WIP

TODO - Fix save and load
*/

const rootPath = "../CommanD-Bot/source/data/"

func Save() error {
	return SaveData("data")
}

func Load() error {
	return LoadData("data")
}

func SaveData(fName string) error {
	log.Println("Saving data..")
	path := filepath.Join(rootPath, fName)
	if file, err := os.Create(path); err != nil {
		return err
	} else {
		enc := gob.NewEncoder(file)
		if err := enc.Encode(servers.BanTime); err != nil {
			return err
		}
		return nil
	}
}

func LoadData(fName string) error {
	log.Println("Loading data...")
	path := filepath.Join(rootPath, fName)
	if file, err := os.Open(path); err != nil {
		return err
	} else {
		dec := gob.NewDecoder(file)
		if err := dec.Decode(servers.BanTime); err != nil {
			return err
		}
		return nil
	}
}

/*

func OpenFile() error {
	filePath, err := filepath.Abs(file)
	if err != nil {
		return err
	}
	file, err = os.Create(filePath)
	return err
}
func (ed *encdec) CloseFile() error {
	err := ed.file.Close()
	return err
}
func (ed *encdec) getFile() *os.File {
	return ed.file
}

func (ed *encdec) NewEncGob() {
	ed.enc = gob.NewEncoder(ed.getFile())
}
func (ed *encdec) EncGob(val interface{}) error {
	err := ed.enc.Encode(val)
	return err
}

func (ed *encdec) NewDecGob() {
	ed.dec = gob.NewDecoder(ed.getFile())
}
func (ed *encdec) DecGob(val interface{}) error {
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
