package hash

/*
  map[string][]string
  map of word-regex and list of root words that match

  during build, create regex of each word. eg bingo = _ingo, b_ngo, bi_go, bin_o, bing_

during get generate regex and for each regex search in map.
if match found, then iterate thru the []string and make sure root word doesnt match searched word.
*/
