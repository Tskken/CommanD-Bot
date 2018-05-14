package CommanD_Bot

import (
	"encoding/gob"
	"log"
	"os"
	"path/filepath"
)

// Default path for all locally saved data //
const dataPath = "../CommanD-Bot/source/data"

// Save server data //
// - returns an error (nil if non)
func saveServer() error {
	log.Println("Saving server data...")
	// Get path to data folder //
	// - returns an error if err is not nil
	path, err := filepath.Abs(dataPath)
	if err != nil {
		return err
	}

	// Open file at given path //
	// - creates file if it does not already exist
	// - returns an error if err is not nil
	file, err := os.Create(filepath.Join(path + "/server_data"))
	if err != nil {
		return err
	}

	// Create gob encoder for file //
	enc := gob.NewEncoder(file)

	// Encode server list data to file //
	// - returns an error if err is not nil
	if err := enc.Encode(serverList); err != nil {
		return err
	}

	return nil
}

// Load data from server data file //
// - returns an error (nil if non)
func loadServer() error {
	log.Println("Loading server data...")
	// Get path to data folder //
	// - returns an error if err is not nil
	path, err := filepath.Abs(dataPath)
	if err != nil {
		return err
	}

	// Open server data file to read data from //
	// - returns an error if err is not nil
	file, err := os.Open(filepath.Join(path + "/server_data"))
	if err != nil {
		return err
	}

	// Create gob decoder for file //
	dec := gob.NewDecoder(file)

	// Decode server list data from file //
	// - returns an error if err is not nil
	if err := dec.Decode(&serverList); err != nil {
		return err
	}

	return nil
}
