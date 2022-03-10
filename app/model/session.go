package model

// Sessions model
type Sessions struct {
	IDPersonal int64    `json:"id_personal"`
	IDPeran    []string `json:"id_peran"`
	//UserID     int64  `json:"user_id"`
	//Name       string `json:"name"`
	//Username   string `json:"username"`
}
