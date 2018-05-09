package CommanD_Bot

import (
	"encoding/gob"
	"log"
	"os"
	"path/filepath"
)

/*
WIP

TODO - Fix save and load
*/

const dataPath = "../CommanD-Bot/source/data"

func saveServer() error {
	log.Println("Saving server data...")
	path, err := filepath.Abs(dataPath)
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(path + "/server_data"))
	if err != nil {
		return err
	}
	enc := gob.NewEncoder(file)
	err = enc.Encode(serverList)
	if err != nil {
		return err
	}
	return nil
}

func loadServer() error {
	log.Println("Loading server data...")
	path, err := filepath.Abs(dataPath)
	if err != nil {
		return err
	}

	file, err := os.Open(filepath.Join(path + "/server_data"))
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(file)
	err = dec.Decode(&serverList)
	if err != nil {
		return err
	}
	return nil
}
