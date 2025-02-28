type key struct {
    mask int
    curTime int
}

var dp = map[key] int {}

func minSessions(tasks []int, sessionTime int) int {
    dp = map[key] int {}
    return 1 + helper(tasks, sessionTime, 0, 0)
}

func helper(tasks []int, sessionTime, timeUsed, tasksDone int) int {
    if val, ok := dp[key{mask: tasksDone, curTime: timeUsed}]; ok {
        return val
    }

    res := 100

    for i := 0; i < len(tasks); i++ {
        if tasksDone & (1 << i) == 0 {
            x := 0
            if timeUsed + tasks[i] > sessionTime {
                x = 1 + helper(tasks, sessionTime, tasks[i], tasksDone | (1 << i))
            } else {
                x = helper(tasks, sessionTime, timeUsed + tasks[i], tasksDone | (1 << i))
            }

            res = min(res, x)
        }
    }

    if res == 100 {
        res = 0
    }

    dp[key{mask: tasksDone, curTime: timeUsed}] = res

    return res
}