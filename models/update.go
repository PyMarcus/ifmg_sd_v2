package models

type UpdateData struct {
    Name      string `json:"name"`
    Email     string `json:"email"`
    OldName   string `json:"old_name"`
    OldEmail  string `json:"old_email"`
}