package comment

func (cs *CommentService) DeleteComments(commentIDs []int) error {
	// Panggil repository untuk menghapus komentar
	err := cs.commentRepo.DeleteComments(commentIDs)
	if err != nil {
		return err
	}
	return nil
}
