package request

// UserRequest
// @Description User account information
type UserRequest struct {
	ID    string        `json:"id" validate:"userid"`
	Name  string        `json:"name" validate:"required,min=2,max=50"`
	Age   int           `json:"age" validate:"required,gte=1"`
	Roles []string      `json:"roles"`
	Email string        `json:"email" validate:"required,min=2,max=50"`
	Hobby *HobbyRequest `json:"hobby" validate:"omitempty"`
	Spec  *SpecRequest  `json:"spec" validate:"required"`
}

// HobbyRequest
// @Description Hobby information
type HobbyRequest struct {
	ID   int    `json:"id" validate:"gte=1"`
	Name string `json:"name" validate:"required,min=2,max=50"`
}

type SpecRequest struct {
	ID     int `json:"id" validate:"required,gte=1"`
	Tall   int `json:"tall" validate:"omitempty,gte=1"`
	Weight int `json:"weight" validate:"omitempty,gte=1"`
}
