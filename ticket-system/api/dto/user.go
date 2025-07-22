package dto

type RegisterUserRequestDTO struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginUserRequestDTO struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type UserResponseDTO struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}