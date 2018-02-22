package theater

import (
	"github.com/Synaxis/bfheroesFesl/inter/network"
	"github.com/Synaxis/bfheroesFesl/inter/network/codec"

	"github.com/sirupsen/logrus"
)

type ansUBRA struct {
	TheaterID string `fesl:"TID"`
}

// UBRA - SERVER Called to  update server data
func (tM *Theater) UBRA(event network.EventClientCommand) {
	if !event.Client.IsActive {
		logrus.Println("Cli Left")
		return
	}

	event.Client.Answer(&codec.Pkt{
		Type: thtrUBRA,
		Content: ansUBRA{
			TheaterID: event.Command.Msg["TID"],
		},
	})

	gdata := tM.level.NewObject("gdata", event.Command.Msg["GID"])

	if event.Command.Msg["START"] == "1" {
		gdata.Set("AP", "0")
	}

}