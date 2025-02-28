func canVisitAllRooms(rooms [][]int) bool {
    stack := []int{}; stack = append(stack, rooms[0]...) // Making a stack with the elements from rooms[0]
    roomsOpened := 1 // The number of rooms that have been opened
    openedRooms := map[int] bool{ 0 : true } // A map of opened rooms so we don't re-open a room that we have opened
    
    for len(stack) != 0 {
        pop := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        
        if !openedRooms[pop] {
            roomsOpened++
            stack = append(stack, rooms[pop]...)
            openedRooms[pop] = true
        }
        
    }
    
    return roomsOpened == len(rooms)
}