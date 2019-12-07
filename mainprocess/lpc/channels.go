package lpc

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"

	"github.com/josephbudd/crud/domain/lpc"
	"github.com/josephbudd/crud/domain/lpc/message"
)

// Sending is a chanel that sends to the renderer.
type Sending chan interface{}

// Receiving is a channel that receives from the renderer.
type Receiving chan interface{}

var (
	send    Sending
	receive Receiving
)

func init() {
	send = make(chan interface{}, 1024)
	receive = make(chan interface{})
}

// Channels returns the renderer connection channels.
func Channels() (sendChan Sending, receiveChan Receiving) {
	sendChan = send
	receiveChan = receive
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
	case *message.LogMainProcessToRenderer:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 0
	case *message.InitMainProcessToRenderer:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 1
	case *message.AddContactMainProcessToRenderer:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 2
	case *message.EditContactMainProcessToRenderer:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 3
	case *message.GetEditContactMainProcessToRenderer:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 4
	case *message.GetEditSelectContactsPageMainProcessToRenderer:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 5
	case *message.GetPrintContactMainProcessToRenderer:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 6
	case *message.GetPrintSelectContactsPageMainProcessToRenderer:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 7
	case *message.GetRemoveContactMainProcessToRenderer:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 8
	case *message.GetRemoveSelectContactsPageMainProcessToRenderer:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 9
	case *message.ReloadContactsMainProcessToRenderer:
		if bb, err = json.Marshal(msg); err != nil {
			return
		}
		id = 10
	case *message.RemoveContactMainProcessToRenderer:
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
		msg := &message.LogRendererToMainProcess{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 1:
		msg := &message.InitRendererToMainProcess{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 2:
		msg := &message.AddContactRendererToMainProcess{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 3:
		msg := &message.EditContactRendererToMainProcess{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 4:
		msg := &message.GetEditContactRendererToMainProcess{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 5:
		msg := &message.GetEditSelectContactsPageRendererToMainProcess{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 6:
		msg := &message.GetPrintContactRendererToMainProcess{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 7:
		msg := &message.GetPrintSelectContactsPageRendererToMainProcess{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 8:
		msg := &message.GetRemoveContactRendererToMainProcess{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 9:
		msg := &message.GetRemoveSelectContactsPageRendererToMainProcess{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 10:
		msg := &message.ReloadContactsRendererToMainProcess{}
		if err = json.Unmarshal(payload.Cargo, msg); err != nil {
			return
		}
		cargo = msg
	case 11:
		msg := &message.RemoveContactRendererToMainProcess{}
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
