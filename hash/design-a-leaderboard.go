package hash

import "container/list"

/*
Design a Leaderboard class, which has 3 functions:

addScore(playerId, score): Update the leaderboard by adding score to the given player's score. If there is no player with such id in the leaderboard, add him to the leaderboard with the given score.
top(K): Return the score sum of the top K players.
reset(playerId): Reset the score of the player with the given id to 0. It is guaranteed that the player was added to the leaderboard before calling this function.
Initially, the leaderboard is empty.



Example 1:

Input:
["Leaderboard","addScore","addScore","addScore","addScore","addScore","top","reset","reset","addScore","top"]
[[],[1,73],[2,56],[3,39],[4,51],[5,4],[1],[1],[2],[2,51],[3]]
Output:
[null,null,null,null,null,null,73,null,null,null,141]

Explanation:
Leaderboard leaderboard = new Leaderboard ();
leaderboard.addScore(1,73);   // leaderboard = [[1,73]];
leaderboard.addScore(2,56);   // leaderboard = [[1,73],[2,56]];
leaderboard.addScore(3,39);   // leaderboard = [[1,73],[2,56],[3,39]];
leaderboard.addScore(4,51);   // leaderboard = [[1,73],[2,56],[3,39],[4,51]];
leaderboard.addScore(5,4);    // leaderboard = [[1,73],[2,56],[3,39],[4,51],[5,4]];
leaderboard.top(1);           // returns 73;
leaderboard.reset(1);         // leaderboard = [[2,56],[3,39],[4,51],[5,4]];
leaderboard.reset(2);         // leaderboard = [[3,39],[4,51],[5,4]];
leaderboard.addScore(2,51);   // leaderboard = [[2,51],[3,39],[4,51],[5,4]];
leaderboard.top(3);           // returns 141 = 51 + 51 + 39;


Constraints:

1 <= playerId, K <= 10000
It's guaranteed that K is less than or equal to the current number of players.
1 <= score <= 100
There will be at most 1000 function calls.
*/

type Player struct {
	PlayerID int
	Score    int
}

type LeaderBoard struct {
	Lookup map[int]*list.Element
	List   list.List
}

// O(n)
func (lb *LeaderBoard) AddScore(playerID int, score int) {
	// add to list.
}

// O(count)
func (lb *LeaderBoard) Top(count int) int {
	return 0
}

// O(1)
func (lb *LeaderBoard) Reset(playerID int) {

}

// O(n)
func (lb *LeaderBoard) Print() {

}

/*
	lb := hash.NewLeaderBoard()
	lb.AddScore(1, 75)
	lb.AddScore(2, 77)
	lb.AddScore(3, 66)

	top := lb.Top(1)
	fmt.Printf("result %v %v %v %v\n", top, "x","y","z")

	top = lb.Top(3)
	fmt.Printf("result %v %v %v %v\n", top, "x","y","z")

*/
//
//func NewLeaderBoard() PriorityQueue {
//	pq := make(PriorityQueue, 0)
//	heap.Init(&pq)
//
//	return pq
//}
//
////type LeaderBoard struct {
////	queue PriorityQueue
////}
//
//func (pq *PriorityQueue) AddScore(playerID int, score int) {
//	heap.Push(pq, &Player{
//		PlayerID: playerID,
//		Score:    score,
//	})
//}
//
//// sum of top count players
//func (pq *PriorityQueue) Top(count int) int {
//	topScore := 0
//
//	for i:=0; i< count || len(*pq) > 0; i++ {
//		top := heap.Pop(pq).(*Player)
//		topScore = topScore + top.Score
//	}
//
//	return topScore
//}
//
//
//
//// A PriorityQueue implements heap.Interface and holds Items.
//type PriorityQueue []*Player
//
//func (pq PriorityQueue) Len() int { return len(pq) }
//
//func (pq PriorityQueue) Less(i, j int) bool {
//	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
//	return pq[i].Score > pq[j].Score
//}
//
//func (pq PriorityQueue) Swap(i, j int) {
//	pq[i], pq[j] = pq[j], pq[i]
//	pq[i].index = i
//	pq[j].index = j
//}
//
//func (pq *PriorityQueue) Push(x interface{}) {
//	n := len(*pq)
//	item := x.(*Player)
//	item.index = n
//	*pq = append(*pq, item)
//}
//
//func (pq *PriorityQueue) Pop() interface{} {
//	old := *pq
//	n := len(old)
//	item := old[n-1]
//	old[n-1] = nil  // avoid memory leak
//	item.index = -1 // for safety
//	*pq = old[0 : n-1]
//	return item
//}
//
////// update modifies the priority and value of an Item in the queue.
////func (pq *PriorityQueue) update(item *Player, value string, priority int) {
////	item.Score = value
////	item.priority = priority
////	heap.Fix(pq, item.index)
////}
