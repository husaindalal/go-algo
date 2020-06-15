package fb

/*
We have a list of objects representing a news feed like so:

class ObjectType(Enum):
	NORMAL = 1
	PHOTO = 2
	VIDEO = 3

feed_objects = [
	{
		'id': 1,
		'subject': 'This is a subject',
		'body': 'This is a body',
		'type': ObjectType.PHOTO,
		'score': 100
	},
	{
		'id': 2,
		'subject': 'This is a subject 2',
		'body': 'This is a body 2',
		'type': ObjectType.PHOTO,
		'score': 99
	},
	{
		'id': 3,
		'subject': 'This is a subject 3',
		'body': 'This is a body 3',
		'type': ObjectType.NORMAL,
		'score': 85
	},
	{
		'id': 4,
		'subject': 'This is a subject 4',
		'body': 'This is a body 4',
		'type': ObjectType.VIDEO,
		'score': 95
	},
	{
		'id': 5,
		'subject': 'This is a subject 5',
		'body': 'This is a body 5',
		'type': ObjectType.PHOTO,
		'score’: 94
	},
	[...]
]

Write a method that groups content where type = PHOTO together into one object (we'll call it a 'rollup') to be displayed in the feed, with a limit of 3 items per rollup. How you structure this is up to you.

The input to the method is the `feed_objects` list. Make sure the output includes all of the elements of the original input (including non `PHOTO` types)

Be sure to write some unit tests.

Few notes:
 - The returned feed should be ordered by `score`
 - We want higher scored items to be grouped together

[P1 N V P2 P3 N N P P ] -- immutable
[[P1P2P3] N V N N PP ]

sort by score
iterate through array
    if not P next
    if P then
        mark this position
        add the P to array
        if limit 3 or end of array then insert into the position

*/

import "fmt"
import "sort"

type FeedObject struct {
	id int
	score int
	objectType string

	// TODO other elements
}

// type GroupedFeed struct {
//     objects [][]FeedObject
// }

func groupObjects(objs []FeedObject) [][]FeedObject {

	result := [][]FeedObject{}

	// sort
	sort.SliceStable(objs, func(i, j int) bool {
		return objs[i].score > objs[j].score
	})

	fmt.Printf("sorted %v\n", objs)


	photoLoc := -1
	photos := []FeedObject{}

	for _, obj := range objs {
		if obj.objectType != "photo" {
			result = append(result, []FeedObject{obj})

		} else {
			photos = append(photos, obj)
			if len(photos) == 1 {
				result = append(result, photos)
				photoLoc = len(result) -1
			} else if len(photos) <= 3 {
				// clear photos for next round
				// photos = append(photos, obj)
				result[photoLoc] = photos
			}

			if len(photos) == 3 {
				photos = []FeedObject{}
			}
		}
		fmt.Printf("photoLoc %v\n", photoLoc)

	}

	// iterate through array
	//     if not P next
	//     if P then
	//         mark this position
	//         add the P to array
	//         if limit 3 or end of array then insert into the position
	return result

}

func main() {
	objs := []FeedObject{
		{
			id: 1,
			score: 18,
			objectType: "photo",
		},
		{
			id: 2,
			score: 20,
			objectType: "photo",
		},
		{
			id: 3,
			score: 50,
			objectType: "video",
		},
		{
			id: 4,
			score: 20,
			objectType: "normal",
		},
		{
			id: 5,
			score: 15,
			objectType: "photo",
		},
		{
			id: 6,
			score: 10,
			objectType: "photo",
		},
		{
			id: 6,
			score: 10,
			objectType: "normal",
		},
		{
			id: 6,
			score: 10,
			objectType: "photo",
		},
		{
			id: 6,
			score: 10,
			objectType: "photo",
		},
		{
			id: 6,
			score: 10,
			objectType: "photo",
		},
	}
	result := groupObjects(objs)

	fmt.Printf("result %v\n", result)
}

