package data

import "errors"

type CommentPower struct {
	Path  string
	Allow bool
}

func (c *CommentPower) TableName() string {
	return "comment_power"
}

func (r *blogRepo) GetCommentPower(path string) (bool, error) {
	if len(path) != 0 {
		var commentPower CommentPower
		if err := r.data.db.Where("path = ?", path).First(&commentPower).Error; err != nil {
			return false, err
		}
		return commentPower.Allow, nil
	}
	return false, errors.New("empty")
}
func (r *blogRepo) UpdateCommentPower(allow bool) bool {
	if err := r.data.pf.UpdateFunc(CommentPower{}, nil, map[string]interface{}{"allow": allow}, true); err != nil {
		return false
	}
	return true
}
