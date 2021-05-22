package tournament

type Player struct {
	name 	string	`bson:"_id"`
	gold 	int		`bson:"goldmedals"`
	silver 	int		`bson:"silvermedals"`
	bronze 	int 	`bson:"bronzemedals"`
}