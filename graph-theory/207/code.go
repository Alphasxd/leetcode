// 207. 课程表 https://leetcode.cn/problems/course-schedule/

// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
// 在选修某些课程之前需要一些先修课程。先修课程按数组 prerequisites 给出，
// 其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

// 示例 1：
// 输入：numCourses = 2, prerequisites = [[1,0]]
// 输出：true

// 示例 2：
// 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
// 输出：false

package leetcode

func canFinish(numCourses int, prerequisites [][]int) bool {
    // 记录每个节点的后继节点，即邻接表，等同于将 prerequisites 转换为邻接表
    edges := make([][]int, numCourses)
    // 记录每个节点的入度
    inDegree := make([]int, numCourses)
    for _, info := range prerequisites {
        // 以 info[1] 为起点，info[0] 为终点，info[1] -> info[0]
        edges[info[1]] = append(edges[info[1]], info[0])
        inDegree[info[0]]++
    }

    // 将入度为 0 的节点加入队列
    queue := []int{}
    for i := 0; i < numCourses; i++ {
        if inDegree[i] == 0 {
            queue = append(queue, i)
        }
    }

    // BFS
    for len(queue) > 0 {
        // 出队
        node := queue[0]
        queue = queue[1:]
        // 课程数减一
        numCourses--
        // 遍历当前节点的所有后继节点
        for _, next := range edges[node] {
            // 将后继节点的入度减一
            inDegree[next]--
            // 如果入度为 0，加入队列
            if inDegree[next] == 0 {
                queue = append(queue, next)
            }
        }
    }
    return numCourses == 0
}
