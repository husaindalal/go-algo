package hash

/*

map[key][]<timestamp, value>
timestamp array is saved in sorted order
-- question, do the set have timestamps in any order or new ts cannot be less than before

get: go to key
search timestamps serially from back O(k) - k is avg timestamps per key
better - binary search. if not found take the next after last



 */
