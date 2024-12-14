package comment

import "fmt"

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
