package filter

import (
	"bufio"
	"github.com/bwmarrin/discordgo"
	"github.com/jbrukh/bayesian"
	"github.com/tsukinai/CommanD-Bot/botErrors"
	"github.com/tsukinai/CommanD-Bot/commands"
	"github.com/tsukinai/CommanD-Bot/servers"
	"github.com/tsukinai/CommanD-Bot/utility"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/*
TODO - Add saving KeyWordMap to a file
*/

// Constants for algorithm //
const (
	// Bayesian Classifiers //
	Good bayesian.Class = "Good"
	Spam bayesian.Class = "Spam"

	// Algorithm thresholds //
	Thresh        = 0.75
	KeyWordThresh = 0.50

	// Data file root path //
	FilePath = "../CommanD-Bot/source/data/"
)

// Create Classifier filter //
// Returns a reference to a new Classifier.
// - nil Classifier if error
func NewFilter() (*bayesian.Classifier, error) {
	commands.KeyWordMap = make(map[string]bool)

	c := bayesian.NewClassifier(Good, Spam)
	c.ConvertTermsFreqToTfIdf()
	if !c.DidConvertTfIdf {
		return nil, botErrors.NewError("Could not convert to TfIdf.", "filter.go")
	}
	return c, nil
}

// Loads Classifier data from pri-trained data file //
func Load(classes *bayesian.Classifier) error {
	log.Println("loading classifiers...")
	// Get file from file path //
	// - Returns an error if err is not nil
	if path, err := filepath.Abs(FilePath); err != nil {
		return err
	} else {
		// Reads class data from file //
		// - Saves data to Good class
		// - Returns an error if err is not nil
		if err := classes.ReadClassFromFile(Good, path); err != nil {
			return err
		}
		// - Saves data to Spam class
		// - Returns an error if err is not nil
		if err := classes.ReadClassFromFile(Spam, path); err != nil {
			return err
		}
	}

	log.Println("loading key word filters...")
	// Gets list of file paths from path //
	// - Returns an error if err is not nil
	if paths, err := filepath.Glob(FilePath + "word_filter/*"); err != nil {
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
					commands.KeyWordMap[strings.TrimSuffix(v, "\n")] = true
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

// Update given words with in the KeyWordMap //
// - Word: The word to updated
// - val: The new value of the word
// -- True = Bad
// -- False = Good
func editKeyWordMap(val bool, words []string) {
	for _, word := range words {
		commands.KeyWordMap[word] = val
	}
}

// Save Classifier data to file //
// - Returns an error if err is not nil
func Save(classes *bayesian.Classifier) error {
	log.Println("Saving Classifiers...")
	// Get files to save to from path //
	// - Returns an error if err is not nil
	if path, err := filepath.Abs(FilePath); err != nil {
		return err
	} else {
		// Save class to file //
		// - Returns an error if err is not nil
		if err := classes.WriteClassesToFile(path); err != nil {
			return err
		}
		return nil
	}
}

// Scan a message to classify its content //
// - Returns an error if err is not nil
func MScan(s *discordgo.Session, m *discordgo.Message, classes *bayesian.Classifier) error {
	// Parce the messages on a | //
	// - If successful don't run scan on message.
	// -- Run learning on given message (Can only be used by users with the "Admin" tag)
	// - If not successful run scan on message
	if msg := utility.Parce(m.Content, "|"); len(msg) < 2 {
		// Parce returned only one item //
		// - Run scan on message
		// - Returns an error (nil if no error)
		return scan(s, m, classes)
	} else {
		// Parce returned more then one item //
		// Check for Admin privilege //
		// - Returns an error if err is not nil
		if admin, err := servers.IsAdmin(s, m); err != nil {
			return err
		} else if !admin {
			// If not admin return //
			_, err := s.ChannelMessageSend(m.ChannelID, "You do not have permission to teach the bot new things... Get good?!?!?")
			return err
		} else {
			// Had admin permissions //
			// Check second value for options //
			switch msg[1] {
			case "Good":
				// Add words to the key map that you want to be good //
				editKeyWordMap(false, utility.ToLower(utility.ParceInput(msg[0])))
			case "Bad":
				// Add word to the key map that you want to be bad //
				editKeyWordMap(true, utility.ToLower(utility.ParceInput(msg[0])))
			case "Spam":
				classes.Learn(utility.ToLower(utility.ParceInput(msg[0])), Spam)
			default:
				// Returns an error if the value after the | is not "Good", "Bad", or "Spam" //
				return botErrors.NewError("Classifier not correct: "+msg[1], "filter.go")
			}

			return nil
		}
	}
}

// Scan a message and to classify it //
// - Returns an error if err is not nil
func scan(s *discordgo.Session, m *discordgo.Message, classes *bayesian.Classifier) error {
	// Parce message on a space //
	msg := utility.ToLower(utility.ParceInput(m.Content))

	// Check for bad words //
	// Count of number of bad words with in sentence //
	bWordCount := 0
	// Check each word with in KeyWordMap //
	// - Increment BWordCount if word value is true
	for _, ms := range msg {
		if key, _ := commands.KeyWordMap[ms]; key {
			bWordCount++
		}
	}

	// Check average occurrence of bad words with in the message //
	// - Remove the message if average is over threshold
	if avr := float64(bWordCount) / float64(len(msg)); avr >= KeyWordThresh {
		// Get the message to delete //
		// - Returns an error if err is not nil
		if a, err := commands.ToDelete(s, m, "", 0, true); err != nil {
			return err
		} else {
			// Delete message //
			// - Return an error if err is not nil
			if err := s.ChannelMessageDelete(m.ChannelID, a[0]); err != nil {
				return err
			}

			// Notify user of there mistake //
			// - Returns an error if err is not nil
			if _, err := s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" don't say that... that's not nice... bad! :point_up: "); err != nil {
				return err
			}

			return nil
		}
	} else {
		// Find the score of the message with in the Good and Spam categories //
		// - Returns an error if err is not nil
		if scores, inx, strict, err := classes.SafeProbScores(msg); err != nil {
			return err
		} else {
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
						if score >= Thresh {
							// Get message to delete //
							// - Returns an error if err is not nil
							if a, err := commands.ToDelete(s, m, "", 0, true); err != nil {
								return err
							} else {
								// Delete message //
								// - Returns an error if err is not nil
								if err := s.ChannelMessageDelete(m.ChannelID, a[0]); err != nil {
									return err
								}
								// Notify member of there mistake //
								// - Returns an error if err is not nil
								if _, err := s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" don't say that... that's not nice... bad! :point_up: "); err != nil {
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
					if a, err := commands.ToDelete(s, m, "", 0, true); err != nil {
						return err
					} else {
						// Delete message //
						// - Return an error ir err is not nil
						if err := s.ChannelMessageDelete(m.ChannelID, a[0]); err != nil {
							return err
						}
						// Notify member of there mistake //
						// - Returns an error if err is not nil
						if _, err := s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" don't say that... that's not nice... bad! :point_up: "); err != nil {
							return err
						}
					}
				// Returns an error if inx was anything but 0 or 1 //
				default:
					return botErrors.NewError("inx was a value other then 0,1, or 2", "filter.go")
				}
			}
		}
	}

	return nil
}
