package dto

type CreateTicketRequestDTO struct {
    Title       string `json:"title"`
    Description string `json:"description"`
}

type AssignTicketRequestDTO struct {
    AssigneeID string `json:"assignee_id"`
}

type TicketResponseDTO struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    CreatorID   string `json:"creator_id"`
    AssigneeID  string `json:"assignee_id"`
}