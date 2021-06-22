// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.1 DO NOT EDIT.
package gen

// Book defines model for Book.
type Book struct {
	Id   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Tag  *string `json:"tag,omitempty"`
}

// BookResponse defines model for BookResponse.
type BookResponse struct {
	// Embedded struct due to allOf(#/components/schemas/Book)
	Book `yaml:",inline"`
}

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// BookTags defines model for BookTags.
type BookTags []string

// ID defines model for ID.
type ID int64

// Limit defines model for Limit.
type Limit int32

// Order defines model for Order.
type Order string

// FindBooksParams defines parameters for FindBooks.
type FindBooksParams struct {

	// tags to filter by
	Tags *BookTags `json:"tags,omitempty"`

	// maximum number of results to return
	Order *Order `json:"order,omitempty"`

	// maximum number of results to return
	Limit *Limit `json:"limit,omitempty"`
}

// AddBookJSONBody defines parameters for AddBook.
type AddBookJSONBody Book

// AddBookJSONRequestBody defines body for AddBook for application/json ContentType.
type AddBookJSONRequestBody AddBookJSONBody

