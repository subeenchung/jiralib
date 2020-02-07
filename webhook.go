package jiralib


type IssueWebhook struct { 
  Timestamp int `json:"timestamp"`
  WebhookEvent string `json:"webhookEvent"`
  IssueEventType string `json:"issue_event_type_name"`
  User JiraUser `json:"user"`
  Issue IssueStruct `json:"issue"`
}
