package links_storage

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

const defaultValueFileName = "data/DefaultValue"
const linksJsonFileName = "data/Links.json"

type Storage struct {
	m             sync.RWMutex
	links         map[string]string
	defaultData   string
	appIsOnReview bool
}

func (s *Storage) loadLinksMap() error {
	file, err := os.Open(linksJsonFileName)
	if err != nil {
		return errors.New("error opening file \"" + defaultValueFileName + "\" with links: " + err.Error())
	}

	if err = json.NewDecoder(file).Decode(&s.links); err != nil {
		return errors.New("error decoding json with links from file \"" + defaultValueFileName + "\": " + err.Error())
	}
	return nil
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
	if err := s.loadDefaultValue(); err != nil {
		return err
	}
	return s.loadLinksMap()
}

func (s *Storage) SetReviewValue(newValue bool) {
	s.m.Lock()
	defer s.m.Unlock()
	s.appIsOnReview = newValue
}

func (s *Storage) GetValueByKey(key string) (string, bool) {
	s.m.RLock()
	defer s.m.RUnlock()
	if s.appIsOnReview {
		return s.defaultData, true
	}
	value := s.links[key]
	return "\"" + s.links[key] + "\"", value != ""
}
