package theater

import (
	"time"

	"github.com/Synaxis/bfheroesFesl/inter/network"
	"github.com/Synaxis/bfheroesFesl/inter/network/codec"

	"github.com/sirupsen/logrus"
)

type ansClientConnected struct {
	TheaterID   string `fesl:"TID"`
	ConnectedAt int64  `fesl:"TIME"`
	ConnTTL     int    `fesl:"activityTimeoutSecs"`
	Protocol    string `fesl:"PROT"`
}

// CONN - SHARED (???) called on connection
func (tm *Theater) CONN(event network.EventClientCommand) {
	if !event.Client.IsActive {
		logrus.Println("Cli Left")
		return
	}

	event.Client.Answer(&codec.Pkt{
		Type: thtrCONN,
		Content: ansClientConnected{
			TheaterID:   event.Command.Msg["TID"],
			ConnectedAt: time.Now().UTC().Unix(),
			ConnTTL:     int((60 * time.Minute).Seconds()),
			Protocol:    event.Command.Msg["PROT"],
		},
	})
}