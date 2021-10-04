// Copyright (C) 2021 Greenbone Networks GmbH
//
// SPDX-License-Identifier: AGPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package util

import (
	"fmt"
	"sync"
)

type Comparable interface {
	Compare(Comparable) bool
}

// QueueList is a mixture of a List and a Queue. It is able to remove Items
// within the Queue without changing the order. It is also thread safe.
type QueueList struct {
	items []Comparable
	sync.RWMutex
}

// NewQueueList returns a new empty QueueList
func NewQueueList() *QueueList {
	return &QueueList{
		items: make([]Comparable, 0),
	}
}

func (ql *QueueList) String() string {
	ret := "["
	for i, item := range ql.items {
		ret += fmt.Sprintf(" %d: %s ", i+1, item)
	}
	ret += "]"
	return ret
}

// RemoveListItem removes an Item from a list without changing the order.
// Returns true if the item was in the list and false when not.
func (ql *QueueList) RemoveListItem(item Comparable) bool {
	ql.Lock()
	defer ql.Unlock()
	if item == nil {
		return false
	}
	i := 0
	for ; i < len(ql.items); i++ {
		if ql.items[i].Compare(item) {
			ql.items = append(ql.items[:i], ql.items[i+1:]...)
			return true
		}
	}
	return false
}

// Append adds a item to the and of a Queue
func (ql *QueueList) Enqueue(item Comparable) {
	ql.Lock()
	defer ql.Unlock()
	if item != nil {
		ql.items = append(ql.items, item)
	}
}

// Contains checks if an Item is already contained in a QueuList
func (ql *QueueList) Contains(item Comparable) bool {
	ql.RLock()
	defer ql.RUnlock()
	if item == nil {
		return false
	}
	i := 0
	for ; i < len(ql.items); i++ {
		if ql.items[i].Compare(item) {
			return true
		}
	}
	return false
}

// Get returns the item inside the queue list for which Contains would return
// true. This function is used when the item in the list is required and the
// Compare Method of the item do not have to contain the same values.
func (ql *QueueList) Get(item Comparable) (Comparable, bool) {
	ql.RLock()
	defer ql.RUnlock()
	i := 0
	if item == nil {
		return nil, false
	}
	for ; i < len(ql.items); i++ {
		if ql.items[i].Compare(item) {
			return ql.items[i], true
		}
	}
	return nil, false
}

func (ql *QueueList) IsEmpty() bool {
	return len(ql.items) == 0
}

func (ql *QueueList) Size() int {
	return len(ql.items)
}

func (ql *QueueList) Front() Comparable {
	ql.RLock()
	defer ql.RUnlock()
	if ql.IsEmpty() {
		return nil
	}
	return ql.items[0]
}

// Dequeue removes and returns first List Item. Returns nil when list is
// empty
func (ql *QueueList) Dequeue() Comparable {
	ql.Lock()
	defer ql.Unlock()
	if ql.IsEmpty() {
		return nil
	}
	ret := ql.items[0]

	ql.items = ql.items[1:]

	return ret
}
