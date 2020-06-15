package hash

/*
O(n2)
 for each element - [a, b]
	for each element after this element - [c, d]
		if [a, d] and [c, b] also present and [a, d] != [a, b] and [c, d] != [c, b]
        calc area. save min
*/
