package ripestat

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
	Common part of reply from th RIPE Stat API
*/
type ripeReply struct {
	Status   string     `json:"status"`
	Messages [][]string `json:"messages"`
}

type RipeStat struct {
	sourceApp string
}

func New() *RipeStat {
	var rs RipeStat
	return &rs
}

/*
	See section Regular Usage at https://stat.ripe.net/docs/data_api
*/
func (rs *RipeStat) SetSourceApp(a string) {
	rs.sourceApp = a
}

/*
	Make query to the RIPE Stat API and return response
*/
func (rs *RipeStat) getData(t, q string) (ret []byte, err error) {
	url := fmt.Sprintf("https://stat.ripe.net/data/%s/data.json?resource=%s", t, q)
	if len(rs.sourceApp) > 0 {
		url += fmt.Sprintf("&sourceapp=%s", rs.sourceApp)
	}
	r, err := http.Get(url)
	if err != nil {
		return
	}
	defer r.Body.Close()

	switch r.StatusCode {
	case http.StatusOK, http.StatusBadRequest:
		ret, err = ioutil.ReadAll(r.Body)
	case http.StatusNotFound:
		err = errors.New(fmt.Sprintf("Query %s NOT found in the RIPE Stat database", q))
	default:
		err = errors.New(fmt.Sprintf("RIPE Stat database responded with unhandled HTTP code %d", r.StatusCode))
	}

	return
}

/*
	Check status of the response for errors and/or maintenance
*/
func (rs *RipeStat) checkReply(in []byte) (err error) {
	var r ripeReply
	err = json.Unmarshal(in, &r)
	if err != nil {
		return
	}

	switch r.Status {
	case "ok":
	case "error":
		err = errors.New(r.Messages[0][1])
	case "maintenance":
		err = errors.New("RIPE Stat API is in the maintenance")
	default:
		err = errors.New(fmt.Sprintf("RIPE Stat API responded with unknown status (%s)", r.Status))
	}

	return
}
