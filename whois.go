package ripestat

import (
	"encoding/json"
	"errors"
)

/*
	RIPE Stat API reply to whois query
*/
type ripeReplyWhois struct {
	ripeReply

	Data struct {
		Records [][]struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"records"`
	} `json:"data"`
}

type Whois struct {
	/*
		ISO 3166-1 alpha-2 code of the country the Network belongs to
	*/
	CountryCode string `json:"country"`
	/*
		Network description
	*/
	Netname string `json:"netname"`
}

/*
	Query for Whois

	More info at https://stat.ripe.net/docs/data_api#whois
*/
func (rs *RipeStat) GetWhois(in string) (ret Whois, err error) {
	d, err := rs.getData("whois", in)
	if err != nil {
		return
	}

	err = rs.checkReply(d)
	if err != nil {
		return
	}

	var r ripeReplyWhois
	err = json.Unmarshal(d, &r)
	if err != nil {
		return
	}

	if len(r.Data.Records) != 1 {
		err = errors.New("More than one records found, unable to parse")
		return
	}

	if len(r.Data.Records[0]) < 2 {
		err = errors.New("Missing data in the records section, unable to parse")
		return
	}

	for _, v := range r.Data.Records[0] {
		switch v.Key {
		case "netname":
			ret.Netname = v.Value
		case "country":
			ret.CountryCode = v.Value
		}
	}

	if len(ret.Netname) == 0 {
		err = errors.New("Netname not found")
		return
	}

	if len(ret.CountryCode) != 2 {
		err = errors.New("Country code not found")
		return
	}

	return
}
