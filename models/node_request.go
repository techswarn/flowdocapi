package models

import (
	"github.com/go-playground/validator/v10"
	"fmt"
)

type Links struct {
	LinkLabel string `json:"linklabel"`
	Url string `json:"url"`
}


type NodeRequest struct {
	NodeType string `json:"nodetype", validate:"required"`
	Label string `json:"label", validate:"required"`
	Source string `json:"source". validate:"required"`
	EdgeType string `json:"edgetype", validate:"required"`
	Animated bool `json:"animated", validate:"required"`
	Heading string `json:"heading"`
	Description string `json:"description"`
	Error   string `json:"error"`
	Links   []Links `json:"links"`
}

func (node_details *NodeRequest) ValidateStruct() []*ErrorResponse{

	fmt.Printf("Blog request %#v \n", node_details)
	var errors []*ErrorResponse
	validate := validator.New()

	// validate the struct
	err := validate.Struct(node_details)

	// if the validation is failed
    // insert the error inside "errors" variable

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
			var element ErrorResponse
			element.ErrorMessage = getErrorMessage(err)
			element.Field = err.Field()
			errors = append(errors, &element)
		}
	}
    // return the validation errors
	return errors
}