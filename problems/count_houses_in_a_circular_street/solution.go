/**
 * Definition for a street.
 * type Street interface {
 *     OpenDoor()
 *     CloseDoor()
 *     IsDoorOpen() bool
 *     MoveRight()
 *     MoveLeft()
 * }
 */
func houseCount(street Street, k int) int {
    for i := 0; i <= k; i++ {
        street.OpenDoor()
        street.MoveRight()
    }
    street.CloseDoor()
    for i := 0; i <= k; i++ {
        street.MoveRight()
        if !street.IsDoorOpen() {
            return i + 1
        }
    }
    return 0
}