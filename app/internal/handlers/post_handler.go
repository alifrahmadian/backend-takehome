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

type PostHandler struct {
	PostService services.PostService
}

func NewPostHandler(postService *services.PostService) *PostHandler {
	return &PostHandler{
		PostService: *postService,
	}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var req dtos.CreatePostRequest

	userID := c.GetInt64("user_id")

	err := c.ShouldBindJSON(&req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Title":
				responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrPostTitleRequired.Error())
				return
			case "Content":
				responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrPostContentRequired.Error())
				return
			}
		}
	}

	post := &models.Post{
		Title:     req.Title,
		Content:   req.Content,
		AuthorID:  userID,
		CreatedAt: time.Now(),
	}

	newPost, err := h.PostService.CreatePost(post)
	if err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := &dtos.CreatePostResponse{
		ID:        newPost.ID,
		Title:     newPost.Title,
		Content:   newPost.Content,
		AuthorID:  newPost.AuthorID,
		CreatedAt: newPost.CreatedAt,
	}

	responses.SuccessResponse(c, messages.MsgPostSuccessful, resp)
}

func (h *PostHandler) GetPostByID(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrInvalidPostID.Error())
		return
	}

	post, err := h.PostService.GetPostByID(int64(postID))
	if err != nil {
		if err == errors.ErrPostNotFound {
			responses.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := &dtos.GetPostResponse{
		ID:       post.ID,
		Title:    post.Title,
		Content:  post.Content,
		AuthorID: post.AuthorID,
		User: dtos.UserResponse{
			ID:    post.User.ID,
			Name:  post.User.Name,
			Email: post.User.Email,
		},
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	responses.SuccessResponse(c, messages.MsgGetPostSuccessful, resp)
}

func (h *PostHandler) GetAllPosts(c *gin.Context) {
	posts, err := h.PostService.GetAllPosts()
	if err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	responseData := make([]*dtos.GetPostResponse, len(posts))

	for i, post := range posts {
		responseData[i] = &dtos.GetPostResponse{
			ID:       post.ID,
			Title:    post.Title,
			Content:  post.Content,
			AuthorID: post.AuthorID,
			User: dtos.UserResponse{
				ID:    post.User.ID,
				Name:  post.User.Name,
				Email: post.User.Email,
			},
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}
	}

	responses.SuccessResponse(c, messages.MsgGetPostsSuccessful, responseData)
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrInvalidPostID.Error())
		return
	}

	var req dtos.UpdatePostRequest
	err = c.ShouldBindJSON(&req)
	if err != nil {
		responses.ErrorResponse(c, http.StatusBadRequest, errors.ErrInvalidPostRequestBody.Error())
		return
	}

	post, err := h.PostService.GetPostByID(int64(postID))
	if err != nil {
		if err == errors.ErrPostNotFound {
			responses.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userID := c.GetInt64("user_id")
	if post.AuthorID != userID {
		responses.ErrorResponse(c, http.StatusForbidden, errors.ErrUpdatePostForbidden.Error())
		return
	}

	if req.Title != "" {
		post.Title = req.Title
	}

	if req.Content != "" {
		post.Content = req.Content
	}

	post.UpdatedAt = time.Now()

	updatedPost, err := h.PostService.UpdatePost(post)
	if err != nil {
		responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	resp := &dtos.UpdatePostResponse{
		ID:        updatedPost.ID,
		Title:     updatedPost.Title,
		Content:   updatedPost.Content,
		AuthorID:  updatedPost.AuthorID,
		UpdatedAt: updatedPost.UpdatedAt,
	}

	responses.SuccessResponse(c, messages.MsgUpdatePostSuccessful, resp)
}
