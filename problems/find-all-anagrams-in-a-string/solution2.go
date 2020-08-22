package solution

// func findAnagrams(s string, p string) []int {
// 	res := make([]int, 0, len(s))
// 	// templ to recreate a map
// 	// and also to check if char is a char from a string p
// 	charPresenceMapTempl := make(map[byte]int, len(p))
// 	for i := range p {
// 		charPresenceMapTempl[p[i]]++
// 	}
// 	charPresenceMap, mapIsNew := createCharPresenceMap(charPresenceMapTempl), true
// 	anagramStartIndex := 0
// 	for i := 0; i < len(s); i++ {
// 		var v int
// 		var ok bool
// 		if v, ok = charPresenceMap[s[i]]; !ok {
// 			// if actual char is a char from string p
// 			// we start from the possible anagram sustring start index
// 			// and "return" all chars before char s[i] to charPresenceMap
// 			// to continue checking substring is anagram
// 			if _, ok = charPresenceMapTempl[s[i]]; ok {
// 				for ; s[anagramStartIndex] != s[i]; anagramStartIndex++ {
// 					// return back char in a map
// 					charPresenceMap[s[anagramStartIndex]]++
// 				}
// 				// start index is right after the same char as s[i]
// 				anagramStartIndex++
// 				continue
// 			}

// 			if !mapIsNew {
// 				// create new map again
// 				charPresenceMap, mapIsNew = createCharPresenceMap(charPresenceMapTempl), true
// 			}
// 			// set possible anagram start as next char
// 			anagramStartIndex = i + 1
// 			// no char in a map no party
// 			continue
// 		}
// 		// char is present in a map, so map will be changed
// 		mapIsNew = false
// 		if v > 1 {
// 			// just decrement the value in a map and go furhter
// 			charPresenceMap[s[i]]--
// 			continue
// 		}
// 		// only one char eq to current char is left in a map
// 		// remove it
// 		delete(charPresenceMap, s[i])
// 		if len(charPresenceMap) <= 0 {
// 			// // there is a moment when we find an angram
// 			// // save start index
// 			// // --
// 			// // in case when a substring is anagram
// 			// // roll back index to substring start index + 1
// 			// // to check any other sustring in substring coud start an anagram too
// 			// // --
// 			// // this is an index of anagram substring start
// 			// // (after the cycle end i will be incremented by 1)
// 			// i = i + 1 - len(p)
// 			// res = append(res, i)
// 			// // possible new anagram start index
// 			// anagramStartIndex = i + 1
// 			// // create new map again
// 			// charPresenceMap, mapIsNew = createCharPresenceMap(charPresenceMapTempl), true

// 			// return back first char in a map
// 			// to continue checking substring is anagram from the char
// 			// after the start of newly found anogram
// 			res = append(res, anagramStartIndex)
// 			charPresenceMap[s[anagramStartIndex]]++
// 			anagramStartIndex++
// 		}
// 	}

// 	return res
// }

// func createCharPresenceMap(templ map[byte]int) map[byte]int {
// 	charPresenceMap := make(map[byte]int, len(templ))
// 	for k, v := range templ {
// 		charPresenceMap[k] = v
// 	}
// 	return charPresenceMap
// }
