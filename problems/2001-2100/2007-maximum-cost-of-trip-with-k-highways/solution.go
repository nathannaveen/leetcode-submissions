type key struct {
    city int
    toll int
}

type key2 struct {
    a int
    b bool
}

type key3 struct {
    used int
    cur int
}

var dp = map[key3] key2 {}

func maximumCost(n int, highways [][]int, k int) int {
    dp = map[key3] key2 {}

    edges := map[int] []key {}

    for _, highway := range highways {
        edges[highway[0]] = append(edges[highway[0]], key{highway[1], highway[2]})
        edges[highway[1]] = append(edges[highway[1]], key{highway[0], highway[2]})
    }

    res := -1

    for i := 0; i < n; i++ {
        helperRes, foundEnd := helper(edges, k, 1 << i, i)

        if foundEnd && helperRes > res {
            res = helperRes
        }
    }

    return res
}

func helper(edges map[int] []key, k, used, cur int) (int, bool) {
    if k == 0 {
        return 0, true
    }

    if val, ok := dp[key3{used, cur}]; ok {
        return val.a, val.b
    }

    res := 0
    foundKEnd := false

    for _, edge := range edges[cur] {
        if used & (1 << edge.city) == 0 {
            helperRes, endK := helper(edges, k - 1, used | (1 << edge.city), edge.city)

            if endK && edge.toll + helperRes >= res {
                res = edge.toll + helperRes
                foundKEnd = true
            }
        }
    }

    dp[key3{used, cur}] = key2{res, foundKEnd}

    return res, foundKEnd
}