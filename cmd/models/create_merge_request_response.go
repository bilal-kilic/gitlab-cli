package models

type CreateMergeRequestResponse struct {
	ApprovalsBeforeMerge interface{}   `json:"approvals_before_merge"`
	Assignee             interface{}   `json:"assignee"`
	Assignees            []interface{} `json:"assignees"`
	Author               struct {
		AvatarURL string `json:"avatar_url"`
		ID        int64  `json:"id"`
		Name      string `json:"name"`
		State     string `json:"state"`
		Username  string `json:"username"`
		WebURL    string `json:"web_url"`
	} `json:"author"`
	BlockingDiscussionsResolved bool        `json:"blocking_discussions_resolved"`
	ChangesCount                string      `json:"changes_count"`
	ClosedAt                    interface{} `json:"closed_at"`
	ClosedBy                    interface{} `json:"closed_by"`
	CreatedAt                   string      `json:"created_at"`
	Description                 string      `json:"description"`
	DiffRefs                    struct {
		BaseSha  string `json:"base_sha"`
		HeadSha  string `json:"head_sha"`
		StartSha string `json:"start_sha"`
	} `json:"diff_refs"`
	DiscussionLocked            interface{}   `json:"discussion_locked"`
	Downvotes                   int64         `json:"downvotes"`
	FirstDeployedToProductionAt interface{}   `json:"first_deployed_to_production_at"`
	ForceRemoveSourceBranch     interface{}   `json:"force_remove_source_branch"`
	HasConflicts                bool          `json:"has_conflicts"`
	HeadPipeline                interface{}   `json:"head_pipeline"`
	ID                          int64         `json:"id"`
	Iid                         int64         `json:"iid"`
	Labels                      []interface{} `json:"labels"`
	LatestBuildFinishedAt       interface{}   `json:"latest_build_finished_at"`
	LatestBuildStartedAt        interface{}   `json:"latest_build_started_at"`
	MergeCommitSha              interface{}   `json:"merge_commit_sha"`
	MergeError                  interface{}   `json:"merge_error"`
	MergeStatus                 string        `json:"merge_status"`
	MergeWhenPipelineSucceeds   bool          `json:"merge_when_pipeline_succeeds"`
	MergedAt                    interface{}   `json:"merged_at"`
	MergedBy                    interface{}   `json:"merged_by"`
	Milestone                   interface{}   `json:"milestone"`
	Pipeline                    interface{}   `json:"pipeline"`
	ProjectID                   int64         `json:"project_id"`
	Reference                   string        `json:"reference"`
	References                  struct {
		Full     string `json:"full"`
		Relative string `json:"relative"`
		Short    string `json:"short"`
	} `json:"references"`
	Sha                      string      `json:"sha"`
	ShouldRemoveSourceBranch interface{} `json:"should_remove_source_branch"`
	SourceBranch             string      `json:"source_branch"`
	SourceProjectID          int64       `json:"source_project_id"`
	Squash                   bool        `json:"squash"`
	SquashCommitSha          interface{} `json:"squash_commit_sha"`
	State                    string      `json:"state"`
	Subscribed               bool        `json:"subscribed"`
	TargetBranch             string      `json:"target_branch"`
	TargetProjectID          int64       `json:"target_project_id"`
	TaskCompletionStatus     struct {
		CompletedCount int64 `json:"completed_count"`
		Count          int64 `json:"count"`
	} `json:"task_completion_status"`
	TimeStats struct {
		HumanTimeEstimate   interface{} `json:"human_time_estimate"`
		HumanTotalTimeSpent interface{} `json:"human_total_time_spent"`
		TimeEstimate        int64       `json:"time_estimate"`
		TotalTimeSpent      int64       `json:"total_time_spent"`
	} `json:"time_stats"`
	Title     string `json:"title"`
	UpdatedAt string `json:"updated_at"`
	Upvotes   int64  `json:"upvotes"`
	User      struct {
		CanMerge bool `json:"can_merge"`
	} `json:"user"`
	UserNotesCount int64  `json:"user_notes_count"`
	WebURL         string `json:"web_url"`
	WorkInProgress bool   `json:"work_in_progress"`
}
