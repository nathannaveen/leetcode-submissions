/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    low := 0
    high := n
    
    for low < high {
        mid := low + (high - low) / 2
        
        if isBadVersion(mid) {
            high = mid
        } else {
            low = mid + 1
        }
    }
    
    return high
}