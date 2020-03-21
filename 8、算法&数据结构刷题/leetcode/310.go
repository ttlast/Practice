func findMinHeightTrees(n int, edges [][]int) []int {

    // topological sort 拓扑排序
    var retSlice []int
    duSlice := make([]int, n)
    edgSlice := make([][]int, n)
    for i := range edges {
        u := edges[i][0]
        v := edges[i][1]
        duSlice[u] ++
        duSlice[v] ++
        edgSlice[u] = append(edgSlice[u], v)
        edgSlice[v] = append(edgSlice[v], u)
    }

    var startSlice []int
    for i:=0; i<n; i++ {
        if duSlice[i] <= 1 {
            duSlice[i] = 0
            startSlice = append(startSlice, i)
        }
    }

    allNum := n
    for true {
        if allNum <= len(startSlice) {
            retSlice = startSlice
            break;
        }

        allNum -= len(startSlice)

        var tmpMap = make(map[int]int)
        for _,u := range startSlice {
            for _, v := range edgSlice[u] {
                if duSlice[v] <= 0 {
                    continue
                }
                duSlice[v] --
                tmpMap[v] = 1
            }
        }

        startSlice = startSlice[0:0]
        for key,_ := range tmpMap {
            if duSlice[key] <= 1 {
                duSlice[key] = 0
                startSlice = append(startSlice, key)
            }
        }
    }

    return retSlice
}