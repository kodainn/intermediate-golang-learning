package services

import (
	"github.com/kodainn/repository/apperrors"
	"github.com/kodainn/repository/models"
	"github.com/kodainn/repository/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataField.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}