package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"title"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	Goal_amount      int    `json:"goal_amount"`
	Current_amount   int    `json:"current_amount"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.Goal_amount = campaign.Goal_amount
	campaignFormatter.Current_amount = campaign.Current_amount
	campaignFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaignMulti(campaigns []Campaign) []CampaignFormatter {
	var campaignMultiFormatter []CampaignFormatter

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignMultiFormatter = append(campaignMultiFormatter, campaignFormatter)
	}

	return campaignMultiFormatter
}
