// Copyright (C) 2013 Mark Stahl

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package rt

import (
	"fmt"
	"time"
)

type Promise struct {
	ID   uint64
	Addr Mailbox

	Value  uint64
	Valued chan bool

	Behaviors map[string]Value

	Blocking []Message
}

func CreatePromise() *Promise {
	id := NewID(PROMISE)

	n := 128
	promise := &Promise{ID: id, Addr: make(Mailbox, n), Behaviors: map[string]Value{}, Blocking: []Message{}}
	promise.Valued = make(chan bool, 1)

	RT.Heap.Insert(id, promise)
	go promise.New()

	return promise
}

func (promise *Promise) New() {
	for {
		select {
		case <-promise.Valued:
			for _, msg := range promise.Blocking {
				forwardMessage(promise, msg)
			}
			promise.Blocking = []Message{}
		case msg := <-promise.Address():
			msg.ForwardMessage(promise)
		}
	}
}

// Any Definition body or Block will return the last expression to be
// evaluated.  If the last expression is a Message (unary, binary, or keyword)
// then the result will be a Promise. To return a promise, we don't know when
// the promised the value of the Message will be available so we must send
// it an asynchronous message (because it could be returned from a remote
// machine) request the Promise's value on behalf of the received message. 
// Therefore we send the "value" asynchronous message to the Promise, but
// instead of creating a new Promise, we use the same Promise ID of the 
// original message.
// 
func (p *Promise) Return(am *AsyncMsg) {
	async := &AsyncMsg{[]uint64{0, 0}, "value", am.PromisedTo}
	p.Address() <- async
}

func (p *Promise) String() string {
	for p.Value == 0 {
		time.Sleep(10 * time.Millisecond)
	}

	val := RT.Heap.Lookup(p.Value)
	switch val.(type) {
	case *Object:
		obj := val.(*Object)
		id := obj.ID & 0xFFFFFFFFF
		return fmt.Sprintf("%s (0x%x @ %s:%d)", obj.Expr, id, RT.IPAddr, RT.Port)
	case *Peer:
		peer := val.(*Peer)
		id := peer.ID & 0xFFFFFFFFF
		expr := peer.RequestValueExpr()
		if expr == "" {
			return fmt.Sprintf("%s (0x%x @ %s:%d)", "Remote", id, peer.IPAddr, RT.Port)
		}

		return expr
	}

	fmt.Printf("OBJ: %#v\n", val)
	return "Unknown"
}

func (p *Promise) OID() uint64 {
	return p.ID
}

func (p *Promise) Address() Mailbox {
	return p.Addr
}

func (p *Promise) LookupBehavior(name string) Value {
	return p.Behaviors[name]
}
