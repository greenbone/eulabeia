package util

import (
	"fmt"
	"testing"
)

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
		ql.Enqueue(fmt.Sprint(i))
	}

	// Check for elements containing
	if ql.IsEmpty() {
		t.Fatal("Queue list should not be empty\n")
	}
	if ql.Size() != 50 {
		t.Fatal("Queue list should contain 50 items\n")
	}

	// Check for values
	if ql.Contains("foo") {
		t.Fatal("Queue list should not contain foo\n")
	}
	if !ql.Contains("23") {
		t.Fatal("Queue list should contain 23\n")
	}

	// Remove items and save them for later
	remove1 := "15"
	remove2 := "37"
	if !ql.RemoveListItem(remove1) {
		t.Fatalf("Unable to remove %s\n", remove1)
	}
	fmt.Printf("%v\n", ql.items)
	if !ql.RemoveListItem(remove2) {
		t.Fatalf("Unable to remove %s\n", remove2)
	}

	// Check for failing RemoveListItem
	if ql.RemoveListItem("foo") {
		t.Fatal("Should be unable to remove foo\n")
	}

	// Check for first Elements
	if ql.Front() != "0" {
		t.Fatal("0 should be the first item\n")
	}
	if item, ok := ql.Dequeue(); item != "0" || !ok {
		t.Fatal("First item should be 0 and dequeue should have worked\n")
	}
	if item, ok := ql.Dequeue(); item != "1" || !ok {
		t.Fatal("First item should be 1 and dequeue should have worked\n")
	}

	// Remove rest of the Elements and check correct order
	for i := 2; i < 50; i++ {
		if fmt.Sprint(i) == remove1 || fmt.Sprint(i) == remove2 {
			continue
		}
		if item, ok := ql.Dequeue(); item != fmt.Sprint(i) || !ok {
			t.Fatalf("Expected: %s, Got %s\n", fmt.Sprint(i), item)
		}
	}

	// Queue List should be empty again
	if !ql.IsEmpty() {
		t.Fatal("Queue should be empty\n")
	}

}
