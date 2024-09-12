package route

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wesleyfebarretos/challenge-bravo/job/internal/infra/db"
	"github.com/wesleyfebarretos/challenge-bravo/job/internal/infra/service"
	"github.com/wesleyfebarretos/challenge-bravo/pkg/utils"
)

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       int
	Email    string
	Role     string
	Password string
}

func handleSignIn(router *gin.RouterGroup) {
	router.POST("auth", func(c *gin.Context) {
		body := SignInRequest{}

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}

		user := User{}
		fmt.Println(body)
		err := db.Conn.
			QueryRow(c, "SELECT id, email, role, password FROM users WHERE email = $1", body.Email).
			Scan(&user.ID, &user.Email, &user.Role, &user.Password)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
			return

		}

		if !utils.IsValidPassword(user.Password, body.Password) {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": "email or password invalid",
			})
			return
		}

		jwtUser := service.JwtUser{
			Email: user.Email,
			Role:  user.Role,
			ID:    user.ID,
		}

		token, err := service.NewJwtService().CreateToken(jwtUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "email or password invalid",
			})
			return
		}

		res := map[string]any{
			"user": map[string]any{
				"id":    user.ID,
				"email": user.Email,
				"role":  user.Role,
			},
			"token": token,
		}

		c.JSON(http.StatusOK, res)
	})
}
