package y2022

import java.util.*

fun main() {
    var start = ArrayList<Pair<Int, Int>>()
    var end = 0 to 0;
    val graph = object {}.javaClass.getResource("Day12")!!.readText().lines().mapIndexed { rowIdx, chars ->
        chars.mapIndexed { colIdx, c ->
            val actualChar = if (c == 'a') {
                start.add( rowIdx to colIdx)
                'a'
            } else if (c == 'E') {
                end = rowIdx to colIdx
                'z'
            } else {
                c
            }
            actualChar - 'a'
        }.toIntArray()
    }.toTypedArray()

    println(start.map { bfs(it, graph, end) }.map{ println(it); it}.min())

}

private operator fun Array<IntArray>.get(to: Pair<Int, Int>): Int {
    return this[to.first][to.second]
}

fun bfs(start: Pair<Int, Int>, graph: Array<IntArray>, end: Pair<Int, Int>): Int {
    var queue: Deque<Pair<Int, Int>> = ArrayDeque()
    var queue2: Deque<Pair<Int, Int>> = ArrayDeque()
    val visited = HashSet<Pair<Int, Int>>()
    var steps = 0
    queue.addLast(start)
    while (queue.isNotEmpty() || queue2.isNotEmpty()) {
        if (queue.isEmpty()) {
            queue = queue2
            queue2 = ArrayDeque()
            steps++
//            println("${visited.size} ${queue.size}")
        }
        val from = queue.pollFirst()
        if (visited.contains(from)) continue
        if (from == end) {
            return steps
        }

        if (from.first > 0) { //can go up
            val to = from.first - 1 to from.second
            if (!visited.contains(to) && graph[to] - graph[from] <= 1) queue2.add(to)
        }
        if (from.first < graph.size - 1) { //can go down
            val to = from.first + 1 to from.second
            if (!visited.contains(to) && graph[to] - graph[from] <= 1) queue2.add(to)
        }
        if (from.second > 0) { //can go up
            val to = from.first to from.second - 1
            if (!visited.contains(to) && graph[to] - graph[from] <= 1) queue2.add(to)
        }
        if (from.second < graph[0].size - 1) { //can go down
            val to = from.first to from.second + 1
            if (!visited.contains(to) && graph[to] - graph[from] <= 1) queue2.add(to)
        }
        visited.add(from)

    }

    return Integer.MAX_VALUE
}
