package main

import "encoding/json"

func uppercaseCount(target string) int {
	var ans int
	for _, c := range target {
		if t := c - 'A'; t >= 0 && t < 26 {
			ans++
		}
	}

	return ans
}

func marshalToJson(v interface{}) string {
	ans, _ := json.Marshal(v)
	return string(ans)
}
