// Package strtouint provides easy conversion from string to uint
// useful for converting id from request params to database row id
package strtouint

import "strconv"

// Parse converts string to uint
func Parse(str string) (uint, error) {
	id, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	uid := uint(id)
	return uid, nil
}
