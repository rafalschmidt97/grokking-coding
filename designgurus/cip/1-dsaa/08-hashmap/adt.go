package main

import (
	"fmt"
)

type Record struct {
	Title         string
	PlacementInfo string
	Key           int
}

type HashTable struct {
	buckets    [][]Record
	max_length int
}

func NewHashTable(size int) *HashTable {
	max_length := size
	buckets := make([][]Record, size)
	for i := 0; i < size; i++ {
		buckets[i] = make([]Record, 0)
	}
	return &HashTable{buckets, max_length}
}

func (ht *HashTable) H(key int) int {
	return key % ht.max_length
}

func (ht *HashTable) Insert(item Record) bool {
	index := ht.H(item.Key)

	for _, record := range ht.buckets[index] {
		if record.Key == item.Key {
			return false
		}
	}

	ht.buckets[index] = append(ht.buckets[index], item)
	return true
}

func (ht *HashTable) Search(key int, returnedItem *Record) bool {
	index := ht.H(key)

	for _, record := range ht.buckets[index] {
		if record.Key == key {
			returnedItem.Key = record.Key
			returnedItem.Title = record.Title
			returnedItem.PlacementInfo = record.PlacementInfo
			return true
		}
	}

	return false
}

func (ht *HashTable) Delete(key int) bool {
	index := ht.H(key)
	chain := ht.buckets[index]

	for i, record := range chain {
		if record.Key == key {
			ht.buckets[index] = append(chain[:i], chain[i+1:]...)
			return true
		}
	}

	return false
}

func (ht *HashTable) ShowTable() {
	fmt.Println("Index\tValue (Key, Title, PlacementInfo)")
	for i := 0; i < ht.max_length; i++ {
		fmt.Print(i, "\t")
		chain := ht.buckets[i]
		if len(chain) == 0 {
			fmt.Println("[EMPTY BUCKET]")
		} else {
			for j, record := range chain {
				if j > 0 {
					fmt.Print("--> ")
				}
				fmt.Printf("(%d, %s, %s)", record.Key, record.Title, record.PlacementInfo)
			}
			fmt.Println()
		}
	}
}

type Solution struct{}

func main() {
	tableSize := 11
	hashTable := NewHashTable(tableSize)

	hashTable.Insert(Record{Key: 1701, Title: "Internet of Things", PlacementInfo: "G1 Shelf"})
	hashTable.Insert(Record{Key: 1712, Title: "Statistical Analysis", PlacementInfo: "G1 Shelf"})
	hashTable.Insert(Record{Key: 1718, Title: "Grid Computing", PlacementInfo: "H2 Shelf"})
	hashTable.Insert(Record{Key: 1735, Title: "UML Modeling", PlacementInfo: "G1 Shelf"})
	hashTable.Insert(Record{Key: 1752, Title: "Professional Practices", PlacementInfo: "G2 Shelf"})

	fmt.Println("\nHash Table after initial insertions:")
	hashTable.ShowTable()

	hashTable.Insert(Record{Key: 1725, Title: "Deep Learning with Python", PlacementInfo: "C3 Shelf"})

	fmt.Println("\nHash Table inserting Book Key 1725:")
	hashTable.ShowTable()

	hashTable.Delete(1701)
	hashTable.Delete(1718)

	fmt.Println("\nHash Table after deleting 1701 and 1718:")
	hashTable.ShowTable()
}
