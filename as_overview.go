package ripestat

import (
	"encoding/json"
	"errors"
	"fmt"
)

/*
	RIPE Stat API reply to as-overview query
*/
type ripeReplyAsOverview struct {
	ripeReply

	Data struct {
		Holder string `json:"holder"`
	} `json:"data"`
}

type AsOverview struct {
	/*
		Name of the AS holder
	*/
	AsName string `json:"asname"`
}

/*
	Query for AS Overview

	More info at https://stat.ripe.net/docs/data_api#as-overview
*/
func (rs *RipeStat) GetAsOverview(in uint64) (ret AsOverview, err error) {
	d, err := rs.getData("as-overview", fmt.Sprintf("%d", in))
	if err != nil {
		return
	}

	err = rs.checkReply(d)
	if err != nil {
		return
	}

	var r ripeReplyAsOverview
	err = json.Unmarshal(d, &r)
	if err != nil {
		return
	}

	if len(r.Data.Holder) == 0 {
		err = errors.New("No AS holder returned")
		return
	}
	ret.AsName = r.Data.Holder

	return
}
