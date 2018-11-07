package CommanD_Bot

import (
	"bufio"
	"errors"
	"github.com/jbrukh/bayesian"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/*
TODO - Fix filter
*/

// Constants for algorithm //
const (
	// Bayesian Classifiers //
	good bayesian.Class = "Good"
	spam bayesian.Class = "Spam"
)

var filterClassifier *bayesian.Classifier // Bayesian Classifiers

// Create Classifier filter //
// Returns a reference to a new Classifier.
// - nil Classifier if error
func NewFilter() error {
	filterClassifier = bayesian.NewClassifier(good, spam)
	filterClassifier.ConvertTermsFreqToTfIdf()
	if !filterClassifier.DidConvertTfIdf {
		return errors.New("could not convert to TfIdf")
	}
	return nil
}

// Loads Classifier data from pri-trained data file //
func LoadFilter() error {
	log.Println("loading classifiers...")

	if err := NewFilter(); err != nil {
		return err
	}

	// Get file from file path //
	// - Returns an error if err is not nil
	if path, err := filepath.Abs(dataPath); err != nil {
		return err
	} else {
		// Reads class data from file //
		// - Saves data to Good class
		// - Returns an error if err is not nil
		if err := filterClassifier.ReadClassFromFile(good, path); err != nil {
			return err
		}
		// - Saves data to Spam class
		// - Returns an error if err is not nil
		if err := filterClassifier.ReadClassFromFile(spam, path); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) LoadWordsFromFile() error {
	log.Println("loading key word filter default for server...")
	// Gets list of file paths from path //
	// - Returns an error if err is not nil
	if paths, err := filepath.Glob(dataPath + "word_filter/*"); err != nil {
		return err
	} else {
		// Get each file for each path with in returned list of paths //
		for _, path := range paths {
			// Opens file from path //
			// - Returns error if err is not nil
			if f, err := os.Open(path); err != nil {
				return err
			} else {
				// Create file reader //
				r := bufio.NewReader(f)
				// Read first line of file //
				v, err := r.ReadString('\n')

				// Return error if error is not EOF //
				if err != nil && err != io.EOF {
					return err
				}
				// While there is no error or EOF read opened file //
				for err == nil {
					// Add word to map //
					// Default is all words are bad //
					// - True = bad
					// - False = good
					s.WordFilter[strings.TrimSuffix(v, "\n")] = true
					// Read next line //
					v, err = r.ReadString('\n')

					// Return error if error is not EOF //
					if err != nil && err != io.EOF {
						return err
					}
				}
			}
		}
	}
	return nil
}

// Save Classifier data to file //
// - Returns an error if err is not nil
func SaveFilter() error {
	log.Println("Saving Classifiers...")
	// Get files to save to from path //
	// - Returns an error if err is not nil
	if path, err := filepath.Abs(dataPath); err != nil {
		return err
	} else {
		// Save class to file //
		// - Returns an error if err is not nil
		if err := filterClassifier.WriteClassesToFile(path); err != nil {
			return err
		}
		return nil
	}
}