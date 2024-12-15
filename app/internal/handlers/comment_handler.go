package handlers

import (
	"app/internal/handlers/dtos"
	"app/internal/handlers/responses"
	"app/internal/models"
	"app/internal/services"
	"app/pkg/errors"
	"app/pkg/messages"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CommentHandler struct {
	CommentService services.CommentService
}

func NewCommentHandler(commentService *services.CommentService) *CommentHandler {
	return &CommentHandler{
		CommentService: *commentService,
	}
}

func (h *CommentHandler) AddComment(c *gin.Context) {
	var req dtos.AddCommentRequest

	authorName := c.GetString("name")

	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrInvalidPostID.Error())
		return
	}

	err = c.ShouldBindJSON(&req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Content":
				responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrCommentContentRequired.Error())
				return
			}
		}
	}

	comment := &models.Comment{
		PostID:     int64(postID),
		AuthorName: authorName,
		Content:    req.Content,
		CreatedAt:  time.Now(),
	}

	newComment, err := h.CommentService.AddComment(comment)
	if err != nil {
		if err == errors.ErrPostNotFound {
			responses.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := &dtos.AddCommentResponse{
		ID:         newComment.ID,
		PostID:     newComment.PostID,
		AuthorName: newComment.AuthorName,
		Content:    newComment.Content,
		CreatedAt:  newComment.CreatedAt,
	}

	responses.SuccessResponse(c, messages.MsgAddCommentSuccessful, resp)
}

func (h *CommentHandler) GetCommentsByPostID(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrInvalidPostID.Error())
		return
	}

	comments, err := h.CommentService.GetCommentsByPostID(int64(postID))
	if err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	responseData := make([]*dtos.GetCommentResponse, len(comments))

	for i, comment := range comments {
		responseData[i] = &dtos.GetCommentResponse{
			ID:         comment.ID,
			AuthorName: comment.AuthorName,
			Content:    comment.Content,
			CreatedAt:  comment.CreatedAt,
		}
	}

	responses.SuccessResponse(c, messages.MsgGetCommentsSuccesful, responseData)
}
