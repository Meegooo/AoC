package y2022

import kotlin.math.abs

fun main() {
    val visited = HashSet<Pair<Int, Int>>()
    val knots = Array(10) { 0 to 0}
    visited.add(knots[knots.size-1])
    object {}.javaClass.getResource("Day9")?.readText()!!.lines().forEach {
        val (direction, num) = it.split(" ")
        repeat(num.toInt()) {
            when (direction) {
                "U" -> knots[0] = knots[0].first to knots[0].second - 1
                "D" -> knots[0] = knots[0].first to knots[0].second + 1
                "L" -> knots[0] = knots[0].first - 1 to knots[0].second
                "R" -> knots[0] = knots[0].first + 1 to knots[0].second
            }
            for (k in 1 until knots.size) {
                knots[k] = newTail(knots[k-1], knots[k])
            }
            visited.add(knots.last())
        }
        draw(knots, visited)
        println("------------------------------------------------------------------------------------")
    }
    println(visited.size)
}

fun newTail(head: Pair<Int, Int>, tail: Pair<Int, Int>): Pair<Int, Int> {
    if (abs(head.first - tail.first) <= 1 && abs(head.second - tail.second) <= 1) return tail
    val dx = tail.first - head.first
    val dy = tail.second - head.second
    if (abs(dx) == 0) return head.first to (if (dy < 0) head.second - 1 else head.second + 1)
    if (abs(dy) == 0) return (if (dx < 0) head.first - 1 else head.first + 1) to head.second
    if (abs(dx) <= 1) return head.first to (if (dy < 0) head.second - 1 else head.second + 1)
    if (abs(dy) <= 1) return (if (dx < 0) head.first - 1 else head.first + 1) to head.second
    if (abs(dx) == 2 && abs(dy) == 2) return (if (dx < 0) head.first - 1 else head.first + 1) to (if (dy < 0) head.second - 1 else head.second + 1)
    throw IllegalStateException("$head, $tail")
}

fun draw(knots: Array<Pair<Int, Int>>, visited: HashSet<Pair<Int, Int>>) {
    for (row in -11..5) {
        for (col in -11..11) {
            print(
                when (col to row) {
                    knots[0] -> 'H'
                    in knots -> knots.indexOf(col to row)
                    in visited -> '#'
                    else -> '.'
                }
            )
        }
        println()
    }
    println()
}