# Hash table (aka Hashmap or Dictionary)

## Hashing

It's a technique that takes an input (or 'message') and returns a fixed-size
string, which looks random. The output, known as the hash value, is unique
(mostly) to the given input.

Use cases: quick data retrieval, data integrity checks, password security,
cryptography, data de-duplication, load balancing.

## Hash Tables (Hash Map)

Hash Table implements an associative array abstract data type (or Dictionary
ADT), mapping keys to values. A Hash Table uses a hash function to compute an
index into an array of buckets or slots where the desired value is stored.

- Keys: The inputs we feed into our hash function. They can be any data type -
  numbers, strings, or even objects. The crucial characteristic of keys is that
  they should be unique.

- Values: The actual data that we store in our Hash Table.

- Hash Function - H(x): Responsible for transforming keys into hash values,
  which dictate where we store our data in the table.

- Buckets: The hash value corresponds to a specific location or 'bucket' within
  the Hash Table. Think of buckets as shelves in library.

Basic operations that are performed on Hash Tables:

- Insert(key, value) operation: Calculates the hash index and stores the
  key-value pair in the slot at that index.
- Search(key) operation: Computes to find the slot and returns the value
  associated with the key, or null if not present.
- Delete(key) operation: Removes the key-value pair stored at index .

## Issues with Hash Tables

### Collision

Happens when two different inputs give the same hash. Without proper collision
handling, lookup performance degrades from O(1) to O(N).

#### Open Addressing (Closed Hashing)

Open addressing techniques resolve hash collisions by probing for the next
available slot within the hash table itself. This approach is called open
addressing since it opens up the entire hash table for finding an empty slot
during insertion.

It is also known as closed hashing because the insertion is restricted to
existing slots within the table without using any external data structures.

Depending on how you jump or probe to find the next empty slot, the closed
hashing is further divided into multiple types. Here are the main collision
resolution schemes used in the open-addressing scheme:

- Linear probing: Linear probing is the simplest way to handle collisions by
  linearly searching consecutive slots for the next empty location.
- Quadratic probing: When a collision occurs, the quadratic probing attempts to
  find an alternative empty slot for the key by using a quadratic function to
  probe the successive possible positions.
- Double hashing: It uses two hash functions to determine the next probing
  location in case of a successive collision.
- Random probing: Random probing uses a pseudo-random number generator (PRNG) to
  compute the step size or increment value for probes in case of collisions.

Implementation of insertion, search, and deletion operations is slightly
different for each type of the operations.

#### Separate Chaining (Open Hashing)

Selecting the right closed hashing technique for resolving collisions can be
tough. You need to keep the pros and cons of different strategies in mind and
then have to make a decision.

Separate chaining offers a rather simpler chaining mechanism to resolve
collisions. Each slot in the hash table points to a separate data structure,
such as a linked-list. This linked-list or chain stores multiple elements that
share the same hash index. When a collision occurs, new elements are simply
appended to the existing list in the corresponding slot.

Separate chaining is an "open hashing" technique because the hash table is
"open" to accommodate multiple elements within a single bucket (usually using a
chain) to handle collisions.

Downsides:

- Increased Memory Overhead: Separate chaining requires additional memory to
  store pointers or references to linked lists for each bucket, leading to
  increased memory overhead, especially when dealing with small data sets.
- Cache Inefficiency: As separate chaining involves linked lists, cache
  performance can be negatively impacted due to non-contiguous memory access
  when traversing the lists, reducing overall efficiency.
- External Fragmentation: The dynamic allocation of memory for linked lists can
  lead to external fragmentation, which may affect the performance of memory
  management in the system.
- Worst-Case Time Complexity: In the worst-case scenario, when multiple keys are
  hashed to the same bucket and form long linked lists, the time complexity for
  search, insert, and delete operations can degrade to O(n), making it less
  suitable for time-critical applications.
- Memory Allocation Overhead: Dynamic memory allocation for each new element can
  add overhead and might lead to performance issues when the system is under
  high memory pressure.

### Overflow

Occurs when the number of elements inserted exceeds the fixed capacity allocated
for the underlying bucket array.

Handling:

Ideally, for closed hashing, the load-factor a = n/m should not cross 0.5, where
n is the number of entries and m is table size. Otherwise, the hash table
experiences a significant increase in collisions, problems in searching, and
degrading performance and integrity of the table operations.

Chaining encounters overflow when chain lengths become too long, thereby
increasing the search time. The load-factor "a" can go up to 0.8/9 or before
performance is affected. Resizing the hash table can help alleviate the overflow
issues.

- Resizing is increasing the size of the hash table to avoid overflows and
  maintain certain load-factor. Once the load-factor of the hash table increases
  a certain threshold (e.g., for closed hashing and for chaining), resizing gets
  activated to increase the size of the table.

- Rehashing involves applying a new hash function (s) to all the entries in a
  hash table to make the table more evenly distributed. In context to resizing,
  it means recalculating hashes (according to the new table size) of all the
  entries in the old table and re-inserting those in the new table. Rehashing
  takes O(N) time for N entries.

## Selecting a Hash Function

Characteristics of a good hash function: uniformity and distribution,
efficiency, collision reduction.

Design techniques:

- Division method - Involves calculating the remainder obtained by dividing the
  key by the size of the hash table (the number of buckets). The remainder is
  taken as the hash code.

- Folding method - Involves dividing the key into smaller parts (subkeys) and
  then combining them to calculate the hash code.

  - Digit sum - Split the key into groups of digits, and their sum is taken as
    the hash code.
  - Reduction - Split the key in folds and use a reduction using any associative
    operation like XOR to reduce folds to a single value. Now, pass this single
    value through the simple division hash function to generate the final hash
    value.
  - Mid-square Method - Involves squaring the key, extracting the middle part,
    and then using it as the hash code.
