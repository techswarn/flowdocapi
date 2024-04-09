package handlers

import(
	"net/http"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/techswarn/flowdocapi/services"
	"github.com/techswarn/flowdocapi/models"
	_ "github.com/techswarn/flowdocapi/utils"
)

func Addnodehandler(c *fiber.Ctx) error {
	//Create a struct using Node and edge Model
	var nodeInput *models.NodeRequest = new(models.NodeRequest)

	// parse the request into "userInput" variable
	if err := c.BodyParser(nodeInput); err != nil {
		// if parsing is failed, return an error
		return c.Status(http.StatusBadRequest).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	    // validate the request
	errors := nodeInput.ValidateStruct()

	// if validation is failed, return the validation errors
	if errors != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "validation failed",
			Data:    errors,
		})
	}

	fmt.Printf("%#v", nodeInput)
	return c.JSON(models.Response[string]{
		Success: true,
		Message: "token data",
		Data:    nodeInput,
	})
}