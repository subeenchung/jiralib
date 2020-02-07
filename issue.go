package jiralib

type IssueStruct struct {
  Self string `json:"self"`
  Id string `json:"id"`
  Key string `json:"key"`
  Fields IssueFields `json:"fields"`
}

type IssueFields struct {
  IssueType IssueTypeStruct `json:"issuetype"`
  Timespent string `json:"timespent"`
  Project ProjectStruct `json:"project"`
  FixVersions []interface{} `json:"fixVersions"`
  AggregateTimespent string `json:"aggregatetimespent"`
  Resolution string `json:"resolution"`
  ResolutionDate string `json:"resolutiondate"`
  Created int `json:"created"`
  Priority PriorityStruct `json:"priority"`
  Creator JiraUser `json:"creator"`
  Reporter JiraUser `json:"reporter"`
  Duedate string `json:"duedate"`
  Comment CommentArrayStruct `json:"comment"`
}

type CommentArrayStruct struct {
  Comments []interface{} `json:"comments"`
}

type PriorityStruct struct {
  Self string `json:"self"`
  IconUrl string `json:"iconUrl"`
  Name string `json:name"`
  Id string `json:"id"`
}

type ProjectStruct struct {
  Self string `json:"self"`
  Id string `json:"id"`
  Key string `json:"key"`
  Name string `json:"name"`
  ProjectTypeKey string `json:"projectTypeKey"`
  AvatarUrls AvatarUrlsStruct `json:"avatarUrls"`
}

type IssueTypeStruct struct {
  Self string `json:"self"`
  Id string `json:"id"`
  Desc string `json:"Description"`
  IconUrl string `json:"iconUrl"`
  Name string `json:"name"`
  Subtask bool `json:"subtask"`
  AvatarId int `json:"avatarId"`
}
