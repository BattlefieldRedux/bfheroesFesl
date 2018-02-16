package fesl

import (
	"github.com/Synaxis/bfheroesFesl/inter/network"
	"github.com/Synaxis/bfheroesFesl/inter/network/codec"
	"github.com/sirupsen/logrus"
)

const (
	partition = "partition.partition"
)

type ansStart struct {
	Txn string      `fesl:"TXN"`
	ID  stPartition `fesl:"id"`
}

// Start handles pnow.Start
func (fm *FeslManager) Start(event network.EventClientCommand) {
	event.Client.WriteEncode(&codec.Packet{
		Payload: ansStart{
			Txn: pnowStart,
			ID: stPartition{1,
				event.Command.Message[partition]},
		},
		Step: event.Command.PayloadID,
		Type: pnow,
	})

	fm.Status(event)
	logrus.Println("=PNOW=")
}
