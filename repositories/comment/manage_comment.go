package comment

import (
	"capstone/repositories/models"
	"fmt"
)

func (cr *CommentRepository) DeleteComments(commentIDs []int) error {
	// Hapus komentar berdasarkan IDs
	if err := cr.db.Where("id IN ?", commentIDs).Delete(&models.Comment{}).Error; err != nil {
		return fmt.Errorf("failed to delete comments: %w", err)
	}
	return nil
}
