package mongodb

type Standing struct {
	Username string `bson:"_id"`
	Win      int    `json:"win"`
	Played   int    `json:"played"`
	Elo      int    `json:"elo"`
}

