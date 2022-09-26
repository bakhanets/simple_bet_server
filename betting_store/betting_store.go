package betting_store

import (
	"sync"
)

type Storage struct {
	sync.Mutex
	XBetLink      string
	MelBetLink    string
	WinLink       string
	DefaultData   string
	AppIsOnReview bool
}

func NewBettingStore() *Storage {
	return &Storage{
		XBetLink:   "\"https://refpa6627021.top/L?tag=d_871311m_1599c_&site=871311&ad=1599\"",
		MelBetLink: "\"https://refpa31055.top/L?tag=s_875293m_18637c_&site=875293&ad=18637\"",
		WinLink:    "\"https://1wlint.top/?open=register#x6ev\"",
		DefaultData: `{
  "predictions": [
    {
      "date": "04.10, 19:45 GMT +3",
      "league": "League of champions, Group stage",
      "group": "Group C",
      "playerOne": "Bavaria FC",
      "playerTwo": "FK Viktoria",
      "urlLogoOne": "https://upload.wikimedia.org/wikipedia/commons/thumb/1/1b/FC_Bayern_M%C3%BCnchen_logo_%282017%29.svg/1200px-FC_Bayern_M%C3%BCnchen_logo_%282017%29.svg.png",
      "urlLogoTwo": "https://pictures.sports.ru/s1NwMDA23GnDfr_p-QI5yBP9OMt40U119JsubkpwA_c/fill/400/400/no/1/czM6Ly9zdGF0X3BpY3R1cmUvVEVBTS9tYWluL2ZjX3Zpa3RvcmlhX3BsemVuLnBuZw.png",
      "predict": "Predict: Over 3.5 match goals"
    },
    {
      "date": "04.10, 22:00 GMT +3",
      "league": "League champion, Group stage",
      "group": "Group A",
      "playerOne": "Ajax",
      "playerTwo": "Napoli",
      "urlLogoOne": "https://upload.wikimedia.org/wikipedia/ru/thumb/7/79/Ajax_Amsterdam.svg/1200px-Ajax_Amsterdam.svg.png",
      "urlLogoTwo": "https://upload.wikimedia.org/wikipedia/commons/thumb/2/2d/SSC_Neapel.svg/1200px-SSC_Neapel.svg.png",
      "predict": "Predict: Under 2.5 match goals"
    }
  ]
}`,
		AppIsOnReview: true,
	}
}

func (s *Storage) ChangeReviewValue(newOnReviewValue bool) {
	s.Lock()
	defer s.Unlock()
	s.AppIsOnReview = newOnReviewValue
}

func (s *Storage) getReviewValue() bool {
	s.Lock()
	defer s.Unlock()
	return s.AppIsOnReview
}

func (s *Storage) GetMelBetLink() string {
	if s.getReviewValue() {
		return s.DefaultData
	}
	return s.MelBetLink
}

func (s *Storage) Get1XBetLink() string {
	if s.getReviewValue() {
		return s.DefaultData
	}
	return s.XBetLink
}

func (s *Storage) GetWinLink() string {
	if s.getReviewValue() {
		return s.DefaultData
	}
	return s.WinLink
}
