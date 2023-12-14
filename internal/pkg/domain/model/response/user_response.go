package response

import (
	"ahmadfarras/fiberframework/internal/pkg/domain/model"
	"time"

	"github.com/google/uuid"
)

type UserDetailResponse struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Password string    `json:"password"`
	Email    string    `json:"email"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func BuildUserDetailResponse(user model.User) UserDetailResponse {
	return UserDetailResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		Password:  user.Password,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func BuildUserListResponse(users []model.User) []UserDetailResponse {
	var responses []UserDetailResponse
	for _, user := range users {
		response := BuildUserDetailResponse(user)

		responses = append(responses, response)
	}

	return responses
}
