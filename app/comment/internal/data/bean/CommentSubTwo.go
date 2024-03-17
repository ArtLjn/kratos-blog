/**
 * @author ljn
 * @date 2024-03-16 21:35:44
 * @dev 二级评论结构
 */

package bean

/**
 * CommentSubTwo
 */

type CommentSubTwo struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	ParentID    int    `json:"parent_id" gorm:"parent_id"`
	Comment     string `json:"comment"`
	CommentTime string `json:"comment_time"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	CommentAddr string `json:"comment_addr" gorm:"comment_addr"`
	ArticleID   int    `json:"article_id"`
}

func (c *CommentSubTwo) TableName() string {
	return "comment_sub_two"
}
