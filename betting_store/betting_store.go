package betting_store

import (
	"sync"
)

type BettingData struct {
	XBetLink   string `json:"1xBetLink"`
	MelbetLink string `json:"MelbetLink"`
	WinLink    string `json:"WinLink"`

	AppIsOnReview bool `json:"onReview"`
}

type BettingStore struct {
	sync.Mutex

	bettingData BettingData
}

func New() *BettingStore {
	ts := &BettingStore{}
	ts.bettingData.XBetLink = "https://refpa6627021.top/L?tag=d_871311m_1599c_&site=871311&ad=1599"
	ts.bettingData.MelbetLink = "https://refpa31055.top/L?tag=s_875293m_18637c_&site=875293&ad=18637"
	ts.bettingData.WinLink = "https://1wlint.top/?open=register#x6ev"
	ts.bettingData.AppIsOnReview = false
	return ts
}

func (ts *BettingStore) ChangeReviewValue(newOnReviewValue bool) bool {
	ts.Lock()
	defer ts.Unlock()

	ts.bettingData.AppIsOnReview = newOnReviewValue
	return ts.bettingData.AppIsOnReview
}

func (ts *BettingStore) GetReviewValue() bool {
	return ts.bettingData.AppIsOnReview
}

func (ts *BettingStore) GetMelbetLink() string {
	return ts.bettingData.MelbetLink
}

func (ts *BettingStore) Get1XbetLink() string {
	return ts.bettingData.XBetLink
}

func (ts *BettingStore) GetWinLink() string {
	return ts.bettingData.WinLink
}
