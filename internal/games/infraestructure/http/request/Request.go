package request


type CreateMiniGameRequest struct {
	Type        string                 `json:"type" binding:"required,oneof=memorama escucha traduccion"`
	Language    string                 `json:"language" binding:"required,oneof=tzeltal zapoteco maya"`
	Level       string                 `json:"level" binding:"required,oneof=principiante intermedio avanzado experto"`
	ContentJSON map[string]interface{} `json:"content_json" binding:"required"`
	IsActive    bool                   `json:"is_active"`
}

type CreateGameSessionRequest struct {
    UserID      string                 `json:"user_id" binding:"required"`
    MiniGameID  string                 `json:"minigame_id" binding:"required"`
    CurrentData map[string]interface{} `json:"current_data,omitempty"`
}


type UpdateGameSessionRequest struct {
    Status string `json:"status" binding:"required,oneof=jugando pausado completado abandonado"`
    Score  int    `json:"score" binding:"min=0"`
}