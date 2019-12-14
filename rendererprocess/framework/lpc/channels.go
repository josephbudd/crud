// +build js, wasm

package lpc

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/pkg/errors"

	"github.com/josephbudd/crud/domain/lpc"
	"github.com/josephbudd/crud/domain/lpc/message"
)

/*
	DO NOT EDIT THIS FILE.

	USE THE TOOL kicklpc TO ADD OR REMOVE LPC Messages.

	kicklpc will edit this file for you.

*/

// Sending is a channel that sends to the main process.
type Sending chan interface{}

// Receiving is a channel that receives from the main process.
type Receiving chan interface{}

var (
	send    Sending
	receive Receiving
	eoj     chan struct{}
	global  js.Value
	alert   js.Value
)

func init() {
	send = make(chan interface{}, 1024)
	receive = make(chan interface{}, 1024)
	eoj = make(chan struct{}, 1024)
	g := js.Global()
	global = g
	alert = g.Get("alert")
}

// Channels returns the renderer connection channels.
func Channels() (sendChan, receiveChan chan interface{}, eojChan chan struct{}) {
	sendChan = send
	receiveChan = receive
	eojChan = eoj
	return
}

// Payload converts unmarshalled msg to the correct marshalled payload.
func (sending Sending) Payload(msg interface{}) (payload []byte, err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "sending.Payload")
		}
	}()

	var bb []byte
	var id uint64
	switch msg := msg.(type) {
	case *message.LogRendererToMainProcess:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 0
	case *message.InitRendererToMainProcess:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 1
	case *message.AddContactRendererToMainProcess:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 2
	case *message.EditContactRendererToMainProcess:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 3
	case *message.GetEditContactRendererToMainProcess:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 4
	case *message.GetEditSelectContactsPageRendererToMainProcess:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 5
	case *message.GetPrintContactRendererToMainProcess:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 6
	case *message.GetPrintSelectContactsPageRendererToMainProcess:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 7
	case *message.GetRemoveContactRendererToMainProcess:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 8
	case *message.GetRemoveSelectContactsPageRendererToMainProcess:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 9
	case *message.ReloadContactsRendererToMainProcess:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 10
	case *message.RemoveContactRendererToMainProcess:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 11
	default:
		bb = []byte("Unknown!")
		id = 999
	}
	pl := &lpc.Payload{
		ID:    id,
		Cargo: bb,
	}
	payload, err = json.Marshal(pl)
	return
}

// Cargo returns a marshalled payload's unmarshalled cargo.
func (receiving Receiving) Cargo(payloadbb []byte) (cargo interface{}, err error) {

	defer func() {
		if err != nil {
			err = errors.WithMessage(err, "receiving.Cargo")
		}
	}()

	payload := lpc.Payload{}
	if err = json.Unmarshal(payloadbb, &payload); err != nil {
		return
	}
	switch payload.ID {
	case 0:
		msg := &message.LogMainProcessToRenderer{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 1:
		msg := &message.InitMainProcessToRenderer{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 2:
		msg := &message.AddContactMainProcessToRenderer{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 3:
		msg := &message.EditContactMainProcessToRenderer{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 4:
		msg := &message.GetEditContactMainProcessToRenderer{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 5:
		msg := &message.GetEditSelectContactsPageMainProcessToRenderer{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 6:
		msg := &message.GetPrintContactMainProcessToRenderer{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 7:
		msg := &message.GetPrintSelectContactsPageMainProcessToRenderer{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 8:
		msg := &message.GetRemoveContactMainProcessToRenderer{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 9:
		msg := &message.GetRemoveSelectContactsPageMainProcessToRenderer{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 10:
		msg := &message.ReloadContactsMainProcessToRenderer{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 11:
		msg := &message.RemoveContactMainProcessToRenderer{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	default:
		errMsg := fmt.Sprintf("no case found for payload id %d", payload.ID)
		err = errors.New(errMsg)
	}
	return
}