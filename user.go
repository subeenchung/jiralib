package jiralib

type JiraUser struct {
  Self string `json:"self"`
  Name string `json:"name"`
  Key string `json:"key"`
  Email string `json:"emailAddress"`
  Avatar AvatarUrlsStruct `json:"avatarUrls"`
  DisplayName string `json:"displayName"`
  Active bool `json:"active"`
  TZraw string `json:"timeZone"`
}

//AvatarUrls struct is for user avatar url location per size
type AvatarUrlsStruct struct {
  FourEight string `json:"48x48"`
  TwoFour string `json:"24x24"`
  OneSix string `json:"16x16"`
  ThreeTwo string `json:"32x32"`
}
