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

func (cr *CommentRepository) ValidateCommentIDs(commentIDs []int) ([]int, error) {
	var existingIDs []int

	// Query untuk mendapatkan comment IDs yang valid di database
	if err := cr.db.Model(&models.Comment{}).Where("id IN ?", commentIDs).Pluck("id", &existingIDs).Error; err != nil {
		return nil, fmt.Errorf("failed to validate comment IDs: %w", err)
	}

	return existingIDs, nil
}
