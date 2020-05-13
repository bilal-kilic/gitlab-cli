package models

type CreateMergeRequest struct {
	SourceBranch string `json:"source_branch"`
	TargetBranch string `json:"target_branch"`
	Title string `json:"title"`
	Description string `json:"description"`
}
