package comment

import (
	"capstone/entities"
	"fmt"
)

func (cs *CommentService) DeleteComments(commentIDs []int) error {
	// Validasi comment IDs
	existingIDs, err := cs.commentRepo.ValidateCommentIDs(commentIDs)
	if err != nil {
		return err
	}

	// Jika ada ID yang tidak ditemukan, return error
	if len(existingIDs) != len(commentIDs) {
		return fmt.Errorf("some comment IDs do not exist")
	}

	// Hapus komentar
	err = cs.commentRepo.DeleteComments(existingIDs)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CommentService) GetCommentsByNewsID(newsID, page, limit int) ([]entities.Comment, int, error) {
	// Hitung offset untuk pagination
	offset := (page - 1) * limit

	// Panggil repository untuk mendapatkan komentar berdasarkan news ID
	comments, total, err := cs.commentRepo.GetCommentsByNewsID(newsID, offset, limit)
	if err != nil {
		return nil, 0, err
	}

	return comments, total, nil
}
