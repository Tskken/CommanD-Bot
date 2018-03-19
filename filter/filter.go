package filter

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jbrukh/bayesian"
	"github.com/tsukinai/CommanD-Bot/botErrors"
	"github.com/tsukinai/CommanD-Bot/commands"
	"github.com/tsukinai/CommanD-Bot/servers"
	"github.com/tsukinai/CommanD-Bot/utility"
	"log"
	"path/filepath"
)

/*
WIP

TODO - Refine classifiers and threshold.  Possible float smoothing for possible underflow errors.
TODO - Comment
*/

const (
	Good   bayesian.Class = "Good"
	Bad    bayesian.Class = "Bad"
	Spam   bayesian.Class = "Spam"
	Thresh                = 0.8
)

func NewFilter() (*bayesian.Classifier, error) {
	c := bayesian.NewClassifier(Good, Bad, Spam)
	c.ConvertTermsFreqToTfIdf()
	if !c.DidConvertTfIdf {
		return nil, botErrors.NewError("Could not convert to TfIdf.", "filter.go")
	}
	return c, nil
}

/*
func CvTfIdf(c *bayesian.Classifier)error{
	c.ConvertTermsFreqToTfIdf()
	if !c.DidConvertTfIdf {
		return botErrors.NewError("Could not convert to TF-IDF.", "filter.go")
	}
	return nil
}*/

func Load(classes *bayesian.Classifier) error {
	log.Println("loading classifiers...")
	if path, err := filepath.Abs("../CommanD-Bot/source/data/"); err != nil {
		return err
	} else {
		if err := classes.ReadClassFromFile(Good, path); err != nil {
			return err
		}
	}

	if path, err := filepath.Abs("../CommanD-Bot/source/data/"); err != nil {
		return err
	} else {
		if err := classes.ReadClassFromFile(Bad, path); err != nil {
			return err
		}
	}

	if path, err := filepath.Abs("../CommanD-Bot/source/data/"); err != nil {
		return err
	} else {
		if err := classes.ReadClassFromFile(Spam, path); err != nil {
			return err
		}
	}

	return nil
}

func Save(classes *bayesian.Classifier) error {
	log.Println("Saving Classifiers...")
	if path, err := filepath.Abs("../CommanD-Bot/source/data/"); err != nil {
		return err
	} else {
		if err := classes.WriteClassesToFile(path); err != nil {
			return err
		}
		return nil
	}
}

func MScan(s *discordgo.Session, m *discordgo.Message, classes *bayesian.Classifier) error {
	if msg := utility.Parce(m.Content, "|"); len(msg) < 2 {
		return scan(s, m, classes)
	} else {
		if admin, err := servers.IsAdmin(s, m); err != nil {
			return err
		} else if !admin {
			_, err := s.ChannelMessageSend(m.ChannelID, "You do not have permission to teach the bot new things... Get good...")
			return err
		} else {
			switch msg[1] {
			case "Good":
				classes.Learn(utility.ToLower(utility.ParceInput(msg[0])), Good)
			case "Bad":
				classes.Learn(utility.ToLower(utility.ParceInput(msg[0])), Bad)
			case "Spam":
				classes.Learn(utility.ToLower(utility.ParceInput(msg[0])), Spam)
			default:
				return botErrors.NewError("Classifier not correct: "+msg[1], "filter.go")
			}

			return nil
		}
	}
}

func scan(s *discordgo.Session, m *discordgo.Message, classes *bayesian.Classifier) error {
	msg := utility.ToLower(utility.ParceInput(m.Content))

	if score, _, _, err := classes.SafeProbScores(msg); err != nil {
		return err
	} else {
		toDelete := false

		for i, s := range score {
			switch i {
			case 0:
				if s >= Thresh {
					break
				}
			case 1:
				if s >= Thresh {
					toDelete = true
				}
			case 2:
				if s >= Thresh {
					toDelete = true
				}
			}
		}

		if toDelete {
			if a, err := commands.ToDelete(s, m, "", 0, true); err != nil {
				return err
			} else {
				if err := s.ChannelMessageDelete(m.ChannelID, a[0]); err != nil {
					return err
				}
				if _, err := s.ChannelMessageSend(m.ChannelID, m.Author.Mention()+" don't say that... that's not nice... bad! :point_up: "); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
