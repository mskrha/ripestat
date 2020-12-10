package ripestat

import (
	"encoding/json"
	"errors"
	"strconv"
)

/*
	RIPE Stat API reply to network-info query
*/
type ripeReplyNetworkInfo struct {
	ripeReply

	Data struct {
		ASNs   []string `json:"asns"`
		Prefix string   `json:"prefix"`
	} `json:"data"`
}

type NetworkInfo struct {
	/*
		Prefix the requested IP address belongs to
	*/
	Prefix string `json:"prefix"`
	/*
		AS number the requested IP address is announced from
	*/
	ASN uint64 `json:"asn"`
}

/*
	Query for Network Info

	More info at https://stat.ripe.net/docs/data_api#network-info
*/
func (rs *RipeStat) GetNetworkInfo(in string) (ret NetworkInfo, err error) {
	d, err := rs.getData("network-info", in)
	if err != nil {
		return
	}

	err = rs.checkReply(d)
	if err != nil {
		return
	}

	var r ripeReplyNetworkInfo
	err = json.Unmarshal(d, &r)
	if err != nil {
		return
	}

	if len(r.Data.Prefix) == 0 {
		err = errors.New("No prefix returned")
		return
	}
	ret.Prefix = r.Data.Prefix

	switch len(r.Data.ASNs) {
	case 0:
		err = errors.New("No ASNs returned, unable to get the originating ASN")
	case 1:
		ret.ASN, err = strconv.ParseUint(r.Data.ASNs[0], 10, 64)
	default:
		err = errors.New("More ASNs returned, unable to get the originating ASN")
	}

	return
}
