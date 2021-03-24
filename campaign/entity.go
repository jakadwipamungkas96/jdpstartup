package campaign

import "time"

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	Becker_count     int
	Goal_amount      int
	Current_amount   int
	Slug             string
	Created_at       time.Time
	Updated_at       time.Time
	CampaignImages   []CampaignImage
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  int
	Created_at time.Time
	Updated_at time.Time
}
