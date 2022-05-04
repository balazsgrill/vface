package vface

import "testing"

func assertListElement(t *testing.T, list *SliceList[ITexteditModel], i uint, expected string) {
	if list.slice[i].GetContent() != expected {
		t.Errorf("Value at %d (%s) differs from expected %s", i, list.slice[i].GetContent(), expected)
	}
}

func TestAppend(t *testing.T) {
	list := &SliceList[ITexteditModel]{}
	list.Insert(list.Size(), &TexteditModel{Content: "a"})
	list.Insert(list.Size(), &TexteditModel{Content: "b"})

	assertListElement(t, list, 0, "a")
	assertListElement(t, list, 1, "b")
}

func TestInsertFirst(t *testing.T) {
	list := &SliceList[ITexteditModel]{}
	list.Insert(list.Size(), &TexteditModel{Content: "a"})
	list.Insert(list.Size(), &TexteditModel{Content: "b"})
	list.Insert(0, &TexteditModel{Content: "c"})

	assertListElement(t, list, 0, "c")
	assertListElement(t, list, 1, "a")
	assertListElement(t, list, 2, "b")
}

func TestInsertMiddle(t *testing.T) {
	list := &SliceList[ITexteditModel]{}
	list.Insert(list.Size(), &TexteditModel{Content: "a"})
	list.Insert(list.Size(), &TexteditModel{Content: "b"})
	list.Insert(1, &TexteditModel{Content: "c"})

	assertListElement(t, list, 0, "a")
	assertListElement(t, list, 1, "c")
	assertListElement(t, list, 2, "b")
}

func TestRemoveFirst(t *testing.T) {
	list := &SliceList[ITexteditModel]{}
	list.Insert(list.Size(), &TexteditModel{Content: "a"})
	list.Insert(list.Size(), &TexteditModel{Content: "b"})
	list.Remove(0)

	assertListElement(t, list, 0, "b")
}

func TestRemoveMiddle(t *testing.T) {
	list := &SliceList[ITexteditModel]{}
	list.Insert(list.Size(), &TexteditModel{Content: "a"})
	list.Insert(list.Size(), &TexteditModel{Content: "b"})
	list.Insert(list.Size(), &TexteditModel{Content: "c"})
	list.Remove(1)

	assertListElement(t, list, 0, "a")
	assertListElement(t, list, 1, "c")
}
