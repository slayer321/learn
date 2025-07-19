package main

import "fmt"

type ValueSlice struct {
	timeStamp int
	value     string
}

type TimeMap struct {
	storeMap map[string][]ValueSlice
}

func Constructor() TimeMap {

	return TimeMap{
		storeMap: make(map[string][]ValueSlice),
	}

}

func (this *TimeMap) Set(key string, value string, timestamp int) {

	if _, ok := this.storeMap[key]; !ok {
		this.storeMap[key] = []ValueSlice{
			{
				timeStamp: timestamp,
				value:     value,
			},
		}
	} else {
		this.storeMap[key] = append(this.storeMap[key], ValueSlice{timeStamp: timestamp, value: value})
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {

	// if val, ok := this.storeMap[timestamp]; ok {
	// 	if finalVal, ok := val[key]; ok {
	// 		return finalVal
	// 	}
	// }
	res := ""
	if keyValueArray, ok := this.storeMap[key]; ok {
		l := 0
		r := len(keyValueArray) - 1

		for l <= r {
			mid := l + (r-l)/2

			if keyValueArray[mid].timeStamp == timestamp {
				return keyValueArray[mid].value
			}

			if timestamp > keyValueArray[mid].timeStamp && mid == 0 && len(keyValueArray) == 1 {
				return keyValueArray[mid].value

			} else if timestamp > keyValueArray[mid].timeStamp {
				res = keyValueArray[mid].value
				l = mid + 1

			} else if timestamp < keyValueArray[mid].timeStamp {
				r = mid - 1
			}
		}
	}

	return res
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

func main() {

	obj := Constructor()

	obj.Set("foo", "bar", 1)
	val := obj.Get("foo", 1)
	fmt.Printf("val: %v\n", val)
	val = obj.Get("foo", 3)
	fmt.Printf("val: %v\n", val)

	obj.Set("foo", "bar2", 4)
	val = obj.Get("foo", 4)
	fmt.Printf("val: %v\n", val)
	val = obj.Get("foo", 5)
	fmt.Printf("val: %v\n", val)

}
