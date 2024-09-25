package integration

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesleyfebarretos/challenge-bravo/app/internal/infra/web/handler"
	"github.com/wesleyfebarretos/challenge-bravo/app/test/testutils"
)

func TestCreateUserHandler(t *testing.T) {
	t.Run("it should create an user", testutils.RunTest(func(t *testing.T) {
		newUser := handler.CreateUserRequest{
			Active:    testutils.Pointer(true),
			FirstName: "John",
			LastName:  "Doe",
			Password:  "123",
			Email:     "johndoe@test.com",
		}

		response := testutils.SendRequest(t, http.MethodPost, "user", newUser)

		body, _ := io.ReadAll(response.Body)

		// user, err := user_repository.New().GetOneByEmail(context.TODO(), "johndoe@test.com")
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(user)
		// fmt.Println("0000000000000000000000000000000")

		assert.Equal(t, http.StatusCreated, response.StatusCode)
		fmt.Println(string(body))
	}))
}
