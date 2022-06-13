package structure

import (
	"github.com/jinzhu/copier"
	"github.com/ulule/deepcopier"
)

type Option struct {
	IgnoreEmpty bool
	DeepCopy    bool
}

// Copy 结构体映射
func Copy(s, ts interface{}) error {
	return copier.Copy(s, ts)
}

//func CopyWithOption(s interface{}, ts interface{}, opt Option) (err error) {
//	return copier.CopyWithOption(s, ts, copier.Option(opt))
//}

func CopyWithIgnoreEmpty(s interface{}, ts interface{}) (err error) {
	return deepcopier.Copy(ts).To(s)
}

// MapMerge
func MapMerge(s, d map[string]interface{}) map[string]interface{} {
	list := make(map[string]interface{})
	for k, v := range s {
		for x, y := range d {
			if k == x {
				list[k] = y
			} else {
				if _, ok := list[k]; !ok {
					list[k] = v
				}
				if _, ok := list[x]; !ok {
					list[x] = y
				}
			}
		}
	}
	return list
}
