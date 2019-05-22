package model

import "encoding/json"

type Profile struct {
	Name   string
	Gender string
	Age    string
	Height string
	//Weight     string
	Income    string
	Marriage  string
	Education string
	//Occupation string
	Registered string
	//House      string
	//Car        string
	ImageUrl string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s, &profile)
	return profile, err
}
