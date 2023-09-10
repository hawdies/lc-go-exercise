package everyday_ecercise

//现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修 bi 。
//
//例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
//返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回任意一种就可以了。如果不可能完成所有课程，返回一个空数组 。

// 该问题为拓扑排序问题，可使用深度优先搜索或广度有限搜索
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges = make([][]int, numCourses)
		// 遍历状态 0 -> 未遍历 1 -> 遍历中 2 -> 已遍历
		// 同一时刻只有1个节点在遍历中，如果2个或以上节点都在遍历中，说明遇到了环
		visited = make([]int, numCourses)
		result  []int
		// 标记是否存在环
		valid bool = true
		dfs   func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}

	if !valid {
		return []int{}
	}

	for i := 0; i < len(result)/2; i++ {
		result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i]
	}
	return result

}
