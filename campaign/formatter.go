package campaign

import "strings"

type CampaignFormatter struct {
	ID               string `json:"id"`
	UserID           string `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}

	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.ImageUrl = ""
	campaignFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	CampaignFormatter
	Perks  []string                  `json:"perks"`
	User   CampaignUserFormatter     `json:"user"`
	Images []CampaignImagesFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type CampaignImagesFormatter struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{}

	campaignDetailFormatter.CampaignFormatter = FormatCampaign(campaign)
	campaignDetailFormatter.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	campaignDetailFormatter.Perks = perks

	user := campaign.User
	campaignUser := CampaignUserFormatter{}
	campaignUser.Name = user.Name
	campaignUser.ImageUrl = user.AvatarFileName

	campaignDetailFormatter.User = campaignUser

	campaignImages := []CampaignImagesFormatter{}
	for _, image := range campaign.CampaignImages {
		campaignImage := CampaignImagesFormatter{}
		campaignImage.ImageUrl = image.FileName
		campaignImage.IsPrimary = image.IsPrimary == 1

		campaignImages = append(campaignImages, campaignImage)
	}
	campaignDetailFormatter.Images = campaignImages

	return campaignDetailFormatter
}
