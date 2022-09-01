package betting_store

import (
	"sync"
)

type BettingStore struct {
	sync.Mutex
	XBetLink         string
	MelBetLink       string
	WinLink          string
	XBetReviewLink   string
	MelBetReviewLink string
	WinReviewLink    string
	AppIsOnReview    bool
}

func NewBettingStore() *BettingStore {
	return &BettingStore{
		XBetLink:         "https://refpa6627021.top/L?tag=d_871311m_1599c_&site=871311&ad=1599",
		MelBetLink:       "https://refpa31055.top/L?tag=s_875293m_18637c_&site=875293&ad=18637",
		WinLink:          "https://1wlint.top/?open=register#x6ev",
		XBetReviewLink:   "https://google.com",
		MelBetReviewLink: "https://google.com",
		WinReviewLink:    "https://google.com",
		AppIsOnReview:    true,
	}
}

func (bs *BettingStore) ChangeReviewValue(newOnReviewValue bool) {
	bs.Lock()
	defer bs.Unlock()
	bs.AppIsOnReview = newOnReviewValue
}

func (bs *BettingStore) getReviewValue() bool {
	bs.Lock()
	defer bs.Unlock()
	return bs.AppIsOnReview
}

func (bs *BettingStore) GetMelBetLink() string {
	if bs.getReviewValue() {
		return bs.MelBetReviewLink
	}
	return bs.MelBetLink
}

func (bs *BettingStore) Get1XBetLink() string {
	if bs.getReviewValue() {
		return bs.XBetReviewLink
	}
	return bs.XBetLink
}

func (bs *BettingStore) GetWinLink() string {
	if bs.getReviewValue() {
		return bs.WinReviewLink
	}
	return bs.WinLink
}
