package users

import (
	"github.com/gin-gonic/gin"

	"realworld-backend/common"
)

type ProfileSerializer struct {
	C *gin.Context
	UserModel
}

// Declare your response schema here
type ProfileResponse struct {
	ID        uint    `json:"-"`
	Username  string  `json:"username"`
	Bio       string  `json:"bio"`
	Image     *string `json:"image"`
	Following bool    `json:"following"`
}

// Put your response logic including wrap the userModel here.
func (p *ProfileSerializer) Response() ProfileResponse {
	myUserModel := p.C.MustGet("my_user_model").(UserModel)
	profile := ProfileResponse{
		ID:        p.ID,
		Username:  p.Username,
		Bio:       p.Bio,
		Image:     p.Image,
		Following: myUserModel.isFollowing(p.UserModel),
	}
	return profile
}

type UserSerializer struct {
	c *gin.Context
}

type UserResponse struct {
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Bio      string  `json:"bio"`
	Image    *string `json:"image"`
	Token    string  `json:"token"`
}

func (u *UserSerializer) Response() UserResponse {
	myUserModel := u.c.MustGet("my_user_model").(UserModel)
	user := UserResponse{
		Username: myUserModel.Username,
		Email:    myUserModel.Email,
		Bio:      myUserModel.Bio,
		Image:    myUserModel.Image,
		Token:    common.GenToken(myUserModel.ID),
	}
	return user
}
