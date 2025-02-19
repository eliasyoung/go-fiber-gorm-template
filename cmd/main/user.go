package main

import (
	"github.com/eliasyoung/fiber-flavor/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateUserPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (app *application) createUserHandler(c *fiber.Ctx) error {
	reqPayload := CreateUserPayload{}
	err := c.BodyParser(&reqPayload)
	if err != nil {
		return c.JSON(pkg.MessageResponse(-2, err.Error(), err.Error()))
	}

	user, err := app.store.CreateUser(reqPayload.Username, reqPayload.Email)
	if err != nil {
		return c.JSON(pkg.MessageResponse(-2, err.Error(), err.Error()))
	}

	user.Languages = pkg.NilSliceFormater(user.Languages)

	return c.JSON(pkg.SuccessResponse(*user))
}

func (app *application) getAllUsersHandler(c *fiber.Ctx) error {
	users, err := app.store.GetAllUsers()
	if err != nil {
		return c.JSON(pkg.MessageResponse(-2, err.Error(), err.Error()))
	}

	for i, user := range users {
		users[i].Languages = pkg.NilSliceFormater(user.Languages)
	}

	return c.JSON(pkg.SuccessResponse(users))
}

func (app *application) getUserByIdHandler(c *fiber.Ctx) error {
	uid := c.Params("userId")
	uuid, err := uuid.Parse(uid)
	if err != nil {
		return c.JSON(pkg.MessageResponse(-2, err.Error(), err.Error()))
	}

	user, err := app.store.GetUserById(uuid)
	if err != nil {
		return c.JSON(pkg.MessageResponse(-2, err.Error(), err.Error()))
	}

	user.Languages = pkg.NilSliceFormater(user.Languages)

	return c.JSON(pkg.SuccessResponse(user))

}
