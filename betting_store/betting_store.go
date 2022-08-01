package betting_store

import (
	"sync"
)

type BettingData struct {
	XBetLink      string `json:"1xBetLink"`
	LeonLink      string `json:"LeonLink"`
	AppIsOnReview bool   `json:"onReview"`
}

type BettingStore struct {
	sync.Mutex

	bettingData BettingData
}

func New() *BettingStore {
	ts := &BettingStore{}
	ts.bettingData.XBetLink = "" // set 1xbet referral link later
	ts.bettingData.LeonLink = "" // set later
	ts.bettingData.AppIsOnReview = false
	return ts
}

func (ts *BettingStore) ChangeReviewValue(newOnReviewValue bool) bool {
	ts.Lock()
	defer ts.Unlock()

	ts.bettingData.AppIsOnReview = newOnReviewValue
	return ts.bettingData.AppIsOnReview
}

func (ts *BettingStore) GetLionLink() string {
	ts.Lock()
	defer ts.Unlock()
	return ts.bettingData.LeonLink
}

func (ts *BettingStore) Get1XbetLink() string {
	ts.Lock()
	defer ts.Unlock()
	return ts.bettingData.XBetLink
}
