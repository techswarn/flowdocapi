package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/techswarn/flowdocapi/models"
	"github.com/techswarn/flowdocapi/services"
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
	fmt.Printf("Node request: %#v \n", nodeInput)
	data, err := services.CreateNode(*nodeInput)
	
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Response[[]*models.ErrorResponse]{
			Success: false,
			Message: "Error while inserting records",
			Data:    errors,
		})
	}

	fmt.Printf("%#v", data)
	return c.JSON(models.Response[string]{
		Success: true,
		Message: "Node added successfully",
		Data: "data",
	})
}

func GetNodeshandler(c *fiber.Ctx) error {
	var nodes []models.Node = services.GetAllNodes()
	fmt.Printf("List of Nodes: %#v \n", nodes)

	return c.Status(http.StatusOK).JSON(models.Response[[]models.Node]{
		Success: true,
		Message: "List of Nodes",
		Data:  nodes,
	})
}

func GetEdgeshandler(c *fiber.Ctx) error {
	var edges []models.Edge = services.GetAllEdges()
	fmt.Printf("List of Nodes: %#v \n", edges)

	return c.Status(http.StatusOK).JSON(models.Response[[]models.Edge]{
		Success: true,
		Message: "List of edges",
		Data:  edges,
	})
}

func GetArticle(c *fiber.Ctx) error {
	var article []models.Article = services.GetArticle()
	fmt.Printf("List of Nodes: %#v \n", article)

	return c.Status(http.StatusOK).JSON(models.Response[[]models.Article]{
		Success: true,
		Message: "Articles",
		Data:  article,
	})
}