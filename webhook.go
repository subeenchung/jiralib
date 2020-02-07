package jiralib

import (
  "time"
)

type IssueWebhook struct { 
  Timestamp time.Time `json:"timestamp"`
  WebhookEvent string `json:"webhookEvent"`
  IssueEventType string `json:"issue_event_type_name"`
  User JiraUser `json:"user"`
  Issue IssueStruct `json:"issue"`
}
