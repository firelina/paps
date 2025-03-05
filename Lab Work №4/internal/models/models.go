package models

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserResponse struct {
	ID int `json:"user_id"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

type CreateScenarioRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CreateScenarioResponse struct {
	ID int `json:"scenario_id"`
}

type Scenario struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Material struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateMaterialRequest struct {
	Name string `json:"name"`
}

type CreateMaterialResponse struct {
	ID int `json:"material_id"`
}
