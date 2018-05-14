package CommanD_Bot

import (
	"bufio"
	"errors"
	"github.com/bwmarrin/discordgo"
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
func newFilter() error {
	filterClassifier = bayesian.NewClassifier(good, spam)
	filterClassifier.ConvertTermsFreqToTfIdf()
	if !filterClassifier.DidConvertTfIdf {
		return errors.New("could not convert to TfIdf")
	}
	return nil
}

// Loads Classifier data from pri-trained data file //
func loadFilter() error {
	log.Println("loading classifiers...")

	if err := newFilter(); err != nil {
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

func (s *server) loadWordsFromFile() error {
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
func saveFilter() error {
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

// Scan a message and to classify it //
// - Returns an error if err is not nil
func scan(session *discordgo.Session, message *discordgo.Message) error {
	// Parce message on a space //
	msg := toLower(strings.Fields(message.Content))

	g, err := getGuild(session, message)
	if err != nil {
		return err
	}
	server := serverList[g.ID]

	// Check for bad words //
	// Count of number of bad words with in sentence //
	bWordCount := 0
	// Check each word with in KeyWordMap //
	// - Increment BWordCount if word value is true
	for _, ms := range msg {
		if _, ok := server.WordFilter[ms]; ok {
			bWordCount++
		}
	}

	//log.Println(float64(bWordCount) / float64(len(msg)))

	// Check average occurrence of bad words with in the message //
	// - Remove the message if average is over threshold
	if avr := float64(bWordCount) / float64(len(msg)); avr >= serverList[g.ID].wordFilterThresh {
		// Get the message to delete //
		// - Returns an error if err is not nil
		if a, err := toDelete(session, message, "", 0, true); err != nil {
			return err
		} else {
			// Delete message //
			// - Return an error if err is not nil
			if err := session.ChannelMessageDelete(message.ChannelID, a[0]); err != nil {
				return err
			}

			// Notify user of there mistake //
			// - Returns an error if err is not nil
			if _, err := session.ChannelMessageSend(message.ChannelID, message.Author.Mention()+" don't say that... that's not nice... bad! :point_up:"); err != nil {
				return err
			}

			return nil
		}
	} else {
		// Find the score of the message with in the Good and Spam categories //
		// - Returns an error if err is not nil
		if scores, inx, strict, err := filterClassifier.SafeProbScores(msg); err != nil {
			return err
		} else {
			//log.Println(scores)

			// If score of both classes are the same check with threshold //
			if !strict {
				// Goes through scores of each class //
				for i, score := range scores {
					switch i {
					// Skip good class value //
					case 0:
						break
					// Check Spam class value //
					case 1:
						// Check if value is over threshold //
						// - Delete if true
						if score >= serverList[g.ID].spamFilterThresh {
							// Get message to delete //
							// - Returns an error if err is not nil
							if a, err := toDelete(session, message, "", 0, true); err != nil {
								return err
							} else {
								// Delete message //
								// - Returns an error if err is not nil
								if err := session.ChannelMessageDelete(message.ChannelID, a[0]); err != nil {
									return err
								}
								// Notify member of there mistake //
								// - Returns an error if err is not nil
								if _, err := session.ChannelMessageSend(message.ChannelID, message.Author.Mention()+" Why must you spam... No spamming... bad! :point_up:"); err != nil {
									return err
								}
							}
						}
					}
				}
			} else {
				// One of the values was larger then the other //
				// - Good = 0
				// - Spam = 1
				switch inx {
				// Return if Good is the largest value //
				case 0:
					return nil
				// Delete messages if Spam is larger //
				case 1:
					// Get message to delete //
					// - Return an error if err is not nil
					if a, err := toDelete(session, message, "", 0, true); err != nil {
						return err
					} else {
						// Delete message //
						// - Return an error ir err is not nil
						if err := session.ChannelMessageDelete(message.ChannelID, a[0]); err != nil {
							return err
						}
						// Notify member of there mistake //
						// - Returns an error if err is not nil
						if _, err := session.ChannelMessageSend(message.ChannelID, message.Author.Mention()+" Why must you spam... No spamming... bad! :point_up:"); err != nil {
							return err
						}
					}
				// Returns an error if inx was anything but 0 or 1 //
				default:
					return errors.New("inx was a value other then 0,1, or 2")
				}
			}
		}
	}

	return nil
}
