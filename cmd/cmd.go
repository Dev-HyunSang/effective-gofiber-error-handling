package cmd

import (
	"github.com/dev-hyunsang/effective-gofiber-error-handling/database"
	"github.com/dev-hyunsang/effective-gofiber-error-handling/ent"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"time"
)

type ToDo struct {
	ToDoUUID    uuid.UUID `json:"todo_uuid"`
	ToDoContext string    `json:"todo_context"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SuccessResponse struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Data        *ent.ToDo
	RespondedAt time.Time `json:"responded_at"`
}

type SuccessDataResopnse struct {
	Code        int         `json:"code"`
	Message     string      `json:"message"`
	Data        []*ent.ToDo `json:"data"`
	RespondedAt time.Time   `json:"responded_at"`
}

type ErrorResopnse struct {
	Code        int       `json:"code"`
	Message     string    `json:"message"`
	RespondedAt time.Time `json:"responded_at"`
}

func Create(c *fiber.Ctx) error {
	req := new(ToDo)
	if err := c.BodyParser(req); err != nil {

	}

	todoUUID := uuid.New()
	newToDo := ToDo{
		ToDoUUID:    todoUUID,
		ToDoContext: req.ToDoContext,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	data, err := database.Create(newToDo)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(ErrorResopnse{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			RespondedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Code:        fiber.StatusOK,
		Message:     "성공적으로 할일을 만들었어요!.",
		Data:        data,
		RespondedAt: time.Now(),
	})
}

func AllToDoRead(c *fiber.Ctx) error {
	data, err := database.AllToDoRead()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResopnse{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			RespondedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(SuccessDataResopnse{
		Code:        fiber.StatusOK,
		Message:     "성공적으로 모든 할일을 가지고 왔습니다.",
		Data:        data,
		RespondedAt: time.Now(),
	})
}

func ParametersRead(c *fiber.Ctx) error {
	StringtodoUUID := c.Params("uuid")
	todoUUID, err := uuid.Parse(StringtodoUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResopnse{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			RespondedAt: time.Now(),
		})
	}

	data, err := database.ParticularToDoRead(todoUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResopnse{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			RespondedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(SuccessDataResopnse{
		Code:        fiber.StatusOK,
		Message:     "성공적으로 해당 할 일을 가지고 불러왔습니다.",
		Data:        data,
		RespondedAt: time.Now(),
	})
}

func Update(c *fiber.Ctx) error {
	req := new(ToDo)
	if err := c.BodyParser(req); err != nil {
		log.Println(err)
	}

	_, err := database.Update(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResopnse{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			RespondedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Code:        fiber.StatusOK,
		Message:     "성공적으로 ToDo를 수정 하였습니다.",
		Data:        nil,
		RespondedAt: time.Now(),
	})
}

func Delete(c *fiber.Ctx) error {
	stringToDoUUID := c.Params("uuid")

	todoUUID, err := uuid.Parse(stringToDoUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResopnse{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			RespondedAt: time.Now(),
		})
	}

	_, err = database.Delete(todoUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResopnse{
			Code:        fiber.StatusInternalServerError,
			Message:     err.Error(),
			RespondedAt: time.Now(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Code:        fiber.StatusOK,
		Message:     "성공적으로 할일을 제거 했습니다.",
		RespondedAt: time.Now(),
	})
}
