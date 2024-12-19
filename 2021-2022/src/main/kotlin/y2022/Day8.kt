package y2022

import java.lang.Integer.max


fun main1() {
    val graph = ArrayList<List<Int>>()
    object {}.javaClass.getResource("Day8")?.readText()!!.lines().forEach {
        graph.add(it.map(Char::digitToInt).toList())
    }
    val x = graph[0].size
    val y = graph.size
    val state = Array(y) { Array(x) { State.NOT_VISIBLE } }

    var count = 0
    for (row in 0 until y) {
        inner@for (column in 0 until x) {
            if (row == 0 || row == y - 1 || column == 0 || column == x - 1) {
                state[row][column] = State.VISIBLE
                count++
                break@inner
            }
            val value = graph[row][column]
            var maxAbove = -1
            for (i in 0 until row) {
                maxAbove = max(maxAbove, graph[i][column])
            }
            var maxBelow = -1
            for (i in row + 1 until y) {
                maxBelow = max(maxBelow, graph[i][column])
            }
            var maxLeft = -1
            for (i in 0 until column) {
                maxLeft = max(maxLeft, graph[row][i])
            }
            var maxRight = -1
            for (i in column + 1 until y) {
                maxRight = max(maxRight, graph[row][i])
            }
            if (value > listOf(maxAbove, maxBelow, maxLeft, maxRight).min()) {
                count++
                state[row][column] = State.VISIBLE
            }
        }
    }
    println(count)
}

fun main() {
    val graph = ArrayList<List<Int>>()
    object {}.javaClass.getResource("Day8")?.readText()!!.lines().forEach {
        graph.add(it.map(Char::digitToInt).toList())
    }
    val x = graph[0].size
    val y = graph.size

    var maxScore = 0
    for (row in 0 until y) {
        inner@for (column in 0 until x) {
            if (row == 0 || row == y - 1 || column == 0 || column == x - 1) {
                print(0)
                continue@inner
            }
            val value = graph[row][column]
            var countAbove = 0
            for (i in (row-1) downTo(0)) {
                countAbove++
                if (value <= graph[i][column])
                break
            }
            var countBelow = 0
            for (i in row + 1 until y) {
                countBelow++
                if (value <= graph[i][column])
                break
            }
            var countLeft = 0
            for (i in (column-1) downTo 0) {
                countLeft++
                if (value <= graph[row][i])
                break
            }
            var countRight = 0
            for (i in column + 1 until x) {
                countRight++
                if (value <= graph[row][i]) 
                break
            }
            val fold = listOf(countAbove, countBelow, countLeft, countRight).fold(1) { o1, o2 -> o1 * o2 }
            print(fold)
            maxScore = max(maxScore, fold)
        }
        println()
    }
    println(maxScore)
}
enum class State {
    NOT_VISIBLE, VISIBLE
}