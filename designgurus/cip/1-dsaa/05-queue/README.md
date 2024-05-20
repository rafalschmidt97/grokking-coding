# Queue

Queue in computer science is a type of data structure where the element that
enters first is the one that gets accessed first (FIFO).

## FIFO

Terminology:

- Enqueue: add an element at the end.
- Dequeue: removing an element from the front.
- Front: This is the start of the Queue, where the first element resides.
- Rear: This is the end of the Queue, where the last element is placed.

## Types

- Simple Queue: Elements are inserted at the rear and removed from the front,
  following the FIFO principle.

- Deque (Double Ended Queue): Elements can be added or removed from both ends of
  the Queue.

- Circular Queue: The last element points to the first, making a loop.

- Priority Queue: In this type, not all elements are equal. Some are considered
  more important.

- Affinity Queue: In this type, every element has an affinity and is placed with
  an element having the same affinity; otherwise placed at the rear.

## Operations

There are four key operations that you can perform on a queue:

- Enqueue/Offer: This is how we add an element to the queue. The element is
  added to the end (rear). O(1)
- Dequeue/Poll: This is the removal of an element from the queue. The element
  removed is from the front (first). O(1)
- Peek/Front: This operation allows us to see the element on the front of the
  queue without removing it. O(1)
- Is Empty: This operation checks if the queue is empty. O(1)

## Consider

- queue overflow and underflow - verify if not full

## Implementing Stack Data Structure

- Using Linked List
- Using an Array
