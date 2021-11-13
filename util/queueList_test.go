package util

import (
	"fmt"
	"testing"
)

type item string

func (i item) Compare(c Comparable) bool {
	if it, ok := c.(item); ok {
		return i == item(it)
	}
	return false
}

func TestQueueList(t *testing.T) {
	// Create new QueueList
	ql := NewQueueList()

	// Test if new one is empty
	if !ql.IsEmpty() {
		t.Fatal("New queue list should be empty\n")
	}
	if ql.Size() != 0 {
		t.Fatal("New queue list should have a size of 0\n")
	}

	// Fill with 50 values, ordered
	for i := 0; i < 50; i++ {
		ql.Enqueue(item(fmt.Sprint(i)))
	}

	// Check for elements containing
	if ql.IsEmpty() {
		t.Fatal("Queue list should not be empty\n")
	}
	if ql.Size() != 50 {
		t.Fatal("Queue list should contain 50 items\n")
	}

	// Check for values
	if ql.Contains(item("foo")) {
		t.Fatal("Queue list should not contain foo\n")
	}
	if !ql.Contains(item("23")) {
		t.Fatal("Queue list should contain 23\n")
	}

	// Remove items and save them for later
	remove1 := item("15")
	remove2 := item("37")
	if !ql.RemoveListItem(remove1) {
		t.Fatalf("Unable to remove %s\n", remove1)
	}
	fmt.Printf("%v\n", ql.items)
	if !ql.RemoveListItem(remove2) {
		t.Fatalf("Unable to remove %s\n", remove2)
	}

	// Check for failing RemoveListItem
	if ql.RemoveListItem(item("foo")) {
		t.Fatal("Should be unable to remove foo\n")
	}

	// Check for first Elements
	if ql.Front() != item("0") {
		t.Fatal("0 should be the first item\n")
	}
	if i := ql.Dequeue().(item); i != "0" {
		t.Fatal("First item should be 0 and dequeue should have worked\n")
	}
	if i := ql.Dequeue().(item); i != "1" {
		t.Fatal("First item should be 1 and dequeue should have worked\n")
	}

	// Remove rest of the Elements and check correct order
	for i := 2; i < 50; i++ {
		if item(fmt.Sprint(i)) == remove1 || item(fmt.Sprint(i)) == remove2 {
			continue
		}
		if i := ql.Dequeue().(item); i != item(fmt.Sprint(i)) {
			t.Fatalf("Expected: %s, Got %s\n", fmt.Sprint(i), i)
		}
	}

	// Queue List should be empty again
	if !ql.IsEmpty() {
		t.Fatal("Queue should be empty\n")
	}

}
