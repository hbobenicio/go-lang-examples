package util

// CreatedBody is a dummy struct intended to represent a body of a 201 Created HTTP response
type CreatedBody struct {
	ID int64 `json:"id"`
}

// NewCreatedBody creates a new Created object
func NewCreatedBody(id int64) CreatedBody {
	return CreatedBody{
		ID: id,
	}
}
