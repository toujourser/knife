package version

import (
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func NewVersion(version string) Version {
	result := Version{}
	parts := strings.Split(version, ".")
	for index, v := range parts {
		if index == 0 {
			result.Major, _ = strconv.Atoi(v)
		} else if index == 1 {
			result.Minor, _ = strconv.Atoi(v)
		} else if index == 2 {
			result.Patch, _ = strconv.Atoi(v)
		}
	}
	return result
}

func (v Version) Compare(other Version) int {
	if v.Major < other.Major {
		return -1
	} else if v.Major > other.Major {
		return 1
	}

	// Compare minor version
	if v.Minor < other.Minor {
		return -1
	} else if v.Minor > other.Minor {
		return 1
	}
	// Compare patch version
	if v.Patch < other.Patch {
		return -1
	} else if v.Patch > other.Patch {
		return 1
	}
	return 0
}

// Compare 比较两个版本号
//func (v Version) Compare(other Version) int {
//	switch {
//	case v.Major < other.Major:
//		return -1
//	case v.Major > other.Major:
//		return 1
//	case v.Minor < other.Minor:
//		return -1
//	case v.Minor > other.Minor:
//		return 1
//	case v.Patch < other.Patch:
//		return -1
//	case v.Patch > other.Patch:
//		return 1
//	default:
//		return 0
//	}
//}
