# Linked List

A linked list is a linear data structure where elements are stored in nodes, and
each node points to the next one in the sequence, forming a chain-like
structure.

Characteristics:

- Dynamic Size: Easily grows and shrinks in size.
- Efficient Insertions/Deletions: Quick at the beginning and middle of the list.

Node Structure:

- Value: The data element stored in the node.
- Next: A reference to the next node in the list.

## Types of Linked Lists

- Singly Linked List - Each node points to the next node.
- Doubly Linked List - Each node has pointers to both the next and the previous
  nodes.
- Circular Linked List - The last node in the list points back to the first
  node.

## Operations

Singly Linked:

- Insert
  - At the Beginning - O(1): Add a new node before the current head of the list.
  - At the End - O(n): Traverse the list to the last node and add the new node
    after it.
  - After a Given Node - O(n): Traverse the list to the desired node and insert
    the new node after it.
- Delete
  - From the Beginning - O(1): Set the second node as the new head.
  - From the End - O(n): Traverse the list and remove the reference to the last
    node.
  - A Given Node - O(n): Traverse to the node before the one to delete, then
    remove the reference to the node-to-delete.
- Search - O(n) -Traverse the list from the head, comparing each node's data
  with the search value until a match is found or the end of the list is
  reached.

## Skipped exercises

- Reverse linked list
