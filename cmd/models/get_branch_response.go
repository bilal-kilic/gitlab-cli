package models

type GetBranchResponse struct {
	CanPush bool `json:"can_push"`
	Commit  struct {
		AuthorEmail    string      `json:"author_email"`
		AuthorName     string      `json:"author_name"`
		AuthoredDate   string      `json:"authored_date"`
		CommittedDate  string      `json:"committed_date"`
		CommitterEmail string      `json:"committer_email"`
		CommitterName  string      `json:"committer_name"`
		CreatedAt      string      `json:"created_at"`
		ID             string      `json:"id"`
		Message        string      `json:"message"`
		ParentIds      interface{} `json:"parent_ids"`
		ShortID        string      `json:"short_id"`
		Title          string      `json:"title"`
		WebURL         string      `json:"web_url"`
	} `json:"commit"`
	Default            bool   `json:"default"`
	DevelopersCanMerge bool   `json:"developers_can_merge"`
	DevelopersCanPush  bool   `json:"developers_can_push"`
	Merged             bool   `json:"merged"`
	Name               string `json:"name"`
	Protected          bool   `json:"protected"`
}

