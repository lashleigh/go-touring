package main

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

// Action is one of create, update, destroy
type Message struct {
	Action string `json:"action"`
	Place  *Place `json:"place"`
}
type Place struct {
	Guid    string  `json:"guid"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
	Address string  `json:"address"`
}

type Day struct {
	Id       string `json:"_id"`
	EndingAt Place  `json:"end"`
	PrevDay  *Day
	NextDay  *Day
	Polyline [][]float64
	//Waypoints []
}

func (m *Message) route_place() {
	switch m.Action {
	case "create":
		m.Place.insert()
	case "update":
		m.Place.update()
	case "destroy":
		// TODO
	default:
		log.Println("Could not route", m)
	}
}

// If the place already exists then we update the coords
// if the place doesn't exist then we create it
func (p *Place) update() {
	log.Println("UPDATE", p)

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	coll := session.DB("gotour").C("places")
	err = coll.Update(bson.M{"guid": p.Guid}, bson.M{"$set": bson.M{"lat": p.Lat, "lng": p.Lng}})

	if err != nil {
		panic(err)
	}
}

func (p *Place) insert() {
	log.Println("CREATE", p)
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	coll := session.DB("gotour").C("places")
	err = coll.Insert(bson.M{"lat": p.Lat, "lng": p.Lng, "guid": p.Guid, "address": p.Address})

	if err != nil {
		panic(err)
	}

	// Add yourself to the global places array. Turrible.
	places = append(places, p)
}

func get_places(places []*Place) []*Place {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	coll := session.DB("gotour").C("places")
	err = coll.Find(bson.M{}).All(&places)

	if err != nil {
		panic(err)
	}
	return places
}
