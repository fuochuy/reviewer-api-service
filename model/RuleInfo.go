package model

type RuleInfo struct {
	ruleId            int64    `json:"rule_id"`
	ruleTypeId        int64    `json:"rule_type_id"`
	name              string   `json:"name"`
	ruleContent       string   `json:"rule_content"`
	isSystem          int64    `json:"is_system"`
	topicId           int64    `json:"topic_id"`
	ruleContentDetail []string `json:"rule_content_detail"`
}
