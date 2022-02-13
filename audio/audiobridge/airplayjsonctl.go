package audiobridge

import (
	"encoding/json"
	"fmt"
	"ledfx/integrations/airplay2"
)

type AirPlayAction int

const (
	AirPlayActionStopServer AirPlayAction = iota
	AirPlayActionGetClients
)

type AirPlayCTLJSON struct {
	Action AirPlayAction `json:"action"`
}

func (apctl AirPlayCTLJSON) AsJSON() ([]byte, error) {
	return json.Marshal(&apctl)
}

// AirPlay takes a marshalled AirPlayCTLJSON
//
// If AirPlayCTLJSON.Action == AirPlayActionStopServer, the server will stop.
//
// If AirPlayCTLJSON.Action == AirPlayActionGetClients, the first return value will be non-nil.
func (j *JsonCTL) AirPlay(jsonData []byte) (clients []*airplay2.Client, err error) {
	conf := AirPlayCTLJSON{}
	if err := json.Unmarshal(jsonData, &conf); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	switch conf.Action {
	case AirPlayActionStopServer:
		return nil, j.w.br.Controller().AirPlay().StopServer()
	case AirPlayActionGetClients:
		return j.w.br.Controller().AirPlay().Clients()
	}
	return nil, fmt.Errorf("unknown action '%d'", conf.Action)
}
