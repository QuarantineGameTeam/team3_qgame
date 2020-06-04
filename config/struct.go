package conf

type User struct {
	ID        int64          `json:"id"`
	Name      string         `json:"name"`
	Team      sql.NullString `json:"team"`
	Role      sql.NullString `json:"role"`
	Health    float32        `json:"health"`
	Strength  float32        `json:"strength"`
	Defence   float32        `json:"defence"`
	Intellect float32        `json:"intellect"`
	Level     float32        `json:"level"`
}
