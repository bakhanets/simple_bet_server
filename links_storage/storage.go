package links_storage

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"
)

const OnReviewValue = "review"
const NotOnReviewValue = "not review"

const reviewValueDataFileName = "data/ReviewSavedData"
const defaultValueFileName = "data/DefaultValue"
const linksJsonFileName = "data/Links.json"

type Storage struct {
	m             sync.RWMutex
	links         map[string]linkRecord
	defaultData   string
	appIsOnReview bool
}

type linkRecord struct {
	link      string
	countries []string
}

func (s *Storage) loadLinksMap() error {
	file, err := os.Open(linksJsonFileName)
	if err != nil {
		return errors.New("error opening file \"" + linksJsonFileName + "\" with links: " + err.Error())
	}
	var data map[string]map[string][]string
	if err = json.NewDecoder(file).Decode(&data); err != nil {
		return errors.New("error decoding json with links from file \"" + linksJsonFileName + "\": " + err.Error())
	}
	s.links = make(map[string]linkRecord)
	for key, val := range data {
		for link, countries := range val {
			s.links[key] = linkRecord{
				link:      link,
				countries: countries,
			}
		}
	}
	return nil
}

func (s *Storage) loadReviewValue() {
	s.appIsOnReview = true
	if d, err := os.ReadFile(reviewValueDataFileName); err == nil {
		str := string(d)
		if str == OnReviewValue {
			s.appIsOnReview = true
		} else if str == NotOnReviewValue {
			s.appIsOnReview = false
		}
	} else {
		s.SetReviewValue(true)
	}
}

func (s *Storage) loadDefaultValue() error {
	d, err := os.ReadFile(defaultValueFileName)
	if err != nil {
		return errors.New("error loading default value from file \"" + defaultValueFileName + "\": " + err.Error())
	}
	s.defaultData = string(d)
	return nil
}

func (s *Storage) LoadValues() error {
	s.loadReviewValue()
	if err := s.loadDefaultValue(); err != nil {
		return err
	}
	return s.loadLinksMap()
}

func (s *Storage) SetReviewValue(newValue bool) {
	s.m.Lock()
	defer s.m.Unlock()
	s.appIsOnReview = newValue
	var err error
	if newValue {
		err = os.WriteFile(reviewValueDataFileName, []byte(OnReviewValue), 0666)
	} else {
		err = os.WriteFile(reviewValueDataFileName, []byte(NotOnReviewValue), 0666)
	}
	if err != nil {
		log.Println("error saving review data: " + err.Error())
	}
}

func (s *Storage) GetValueByKeyForCountry(key string, isoCountryCode string) (string, bool) {
	if isoCountryCode == "" {
		return "", false
	}
	s.m.RLock()
	defer s.m.RUnlock()
	value, ok := s.links[key]
	if ok {
		if s.appIsOnReview {
			return s.defaultData, true
		} else {
			for _, code := range value.countries {
				if code == isoCountryCode {
					return "\"" + value.link + "\"", true
				}
			}
			return s.defaultData, true
		}
	} else {
		return "", false
	}
}
