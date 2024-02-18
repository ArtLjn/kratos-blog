package bean

type CommentBean struct {
	ArticleID     string         `json:"article_id"`
	CommentTime   string         `json:"comment_time"`
	ChildComments []ChildComment `json:"child_comments,omitempty"`
	CommentAddr   string         `json:"comment_addr"`
	Name          string         `json:"name"`
	Comment       string         `json:"comment"`
	OriginID      string         `json:"origin_id"`
	ID            string         `json:"id"`
	Email         string         `json:"email"`
}

// ChildComment represents a child comment structure
type ChildComment struct {
	RIpAddr       string `json:"reward_addr"`
	RewardTime    string `json:"reward_time"`
	RewardName    string `json:"reward_name"`
	RewardEmail   string `json:"reward_email"`
	RewardContent string `json:"reward_content"`
}
