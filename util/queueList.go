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

import "sync"

// QueueList is a mixture of a List and a Queue. It is able to remove Items
// within the Queue without changing the order. It is also thread safe.
type QueueList struct {
	items []string
	mutex *sync.RWMutex
}

// NewQueueList returns a new empty QueueList
func NewQueueList() *QueueList {
	return &QueueList{
		items: make([]string, 0),
		mutex: &sync.RWMutex{},
	}
}

// RemoveListItem removes an Item from a list without changing the order.
// Returns true if the item was in the list and false when not.
func (ql *QueueList) RemoveListItem(item string) bool {
	ql.mutex.Lock()
	defer ql.mutex.Unlock()
	i := 0
	for ; i < len(ql.items); i++ {
		if ql.items[i] == item {
			ql.items = append(ql.items[:i], ql.items[i+1:]...)
			return true
		}
	}
	return false
}

// Append adds a item to the and of a Queue
func (ql *QueueList) Enqueue(item string) {
	ql.mutex.Lock()
	defer ql.mutex.Unlock()
	ql.items = append(ql.items, item)
}

// Contains checks if an Item is already contained in a QueuList
func (ql *QueueList) Contains(item string) bool {
	ql.mutex.RLock()
	defer ql.mutex.RUnlock()
	i := 0
	for ; i < len(ql.items); i++ {
		if ql.items[i] == item {
			return true
		}
	}
	return false
}

func (ql *QueueList) IsEmpty() bool {
	return len(ql.items) == 0
}

func (ql *QueueList) Size() int {
	return len(ql.items)
}

func (ql *QueueList) Front() string {
	ql.mutex.RLock()
	defer ql.mutex.RUnlock()
	return ql.items[0]
}

// Dequeue removes and returns first List Item. Returns false when list is
// empty
func (ql *QueueList) Dequeue() (string, bool) {
	ql.mutex.Lock()
	defer ql.mutex.Unlock()
	if ql.IsEmpty() {
		return "", false
	}
	ret := ql.items[0]

	ql.items = ql.items[1:]

	return ret, true
}
