package presenter

import (
	"time"

	"github.com/AVieraUY/go-forum/internal/domain/entity"
)

// Comment data
type Comment struct {
	ID        entity.ID `json:"id"`
	UserID    entity.ID `json:"userId"`
	PostID    entity.ID `json:"postId"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
