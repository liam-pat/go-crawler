package model

import "encoding/json"

type Profile struct {
	Name       string
	Gender     string
	Age        string
	Height     string
	Income     string
	Marriage   string
	Education  string
	Registered string
	ImageUrl   string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	//tran interface to []uint8
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	// tran []uint8 to a struct
	err = json.Unmarshal(s, &profile)
	return profile, err
}
