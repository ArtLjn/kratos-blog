/**
 * @author ljn
 * @date 2024-03-16 21:32:50
 * @dev 一级评论
 */

package bean

type CommentParent struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	ArticleID   string `json:"article_id" gorm:"column:article_id"`
	Comment     string `json:"comment"`
	CommentTime string `json:"comment_time"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	CommentAddr string `json:"comment_addr" gorm:"column:comment_addr"`
}

func (c *CommentParent) TableName() string {
	return "comment_parent"
}
