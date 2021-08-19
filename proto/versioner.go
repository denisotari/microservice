package proto

type CurrentVersion struct {
	Version   string `json:"version"`
	CreatedAt int64  `json:"created_at"`
}
