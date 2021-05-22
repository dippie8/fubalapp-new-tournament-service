package tournament

type Standing struct {
	Id     string `bson:"_id"`
	Win    int    `bson:"Win"`
	Played int    `bson:"Played"`
}
