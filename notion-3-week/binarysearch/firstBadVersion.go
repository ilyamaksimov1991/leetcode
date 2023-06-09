package binarysearch

// time O(logN)
// memory O(1)
func firstBadVersion(n int) int {
	left, right := 0, n-1
	badVersion := n

	for left <= right {
		med := (left + right) / 2
		if isBadVersion(med) {
			badVersion = med
			right = med - 1
		} else {
			left = med + 1
		}
	}

	return badVersion
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	return false
}
