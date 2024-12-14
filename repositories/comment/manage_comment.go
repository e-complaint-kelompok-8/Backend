package comment

import (
	"capstone/entities"
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

func (cr *CommentRepository) GetCommentsByNewsID(newsID, offset, limit int) ([]entities.Comment, int, error) {
	var comments []models.Comment

	// Query untuk mendapatkan komentar berdasarkan news ID dengan pagination
	err := cr.db.Preload("User").Preload("News").
		Where("news_id = ?", newsID).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&comments).Error
	if err != nil {
		return nil, 0, err
	}

	// Hitung total komentar untuk berita tertentu
	var total int64
	err = cr.db.Model(&models.Comment{}).
		Where("news_id = ?", newsID).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Konversi model ke entitas
	var result []entities.Comment
	for _, comment := range comments {
		result = append(result, comment.ToEntities())
	}

	return result, int(total), nil
}
