package y2021

import java.io.File
import java.util.*
import kotlin.collections.ArrayList

fun main(args: Array<String>) {
    val sc = Scanner(File("inputs/day9"))
    val arr = ArrayList<List<Int>>()
    var width = 0
    var height = 0
    while (sc.hasNext()) {
        val line: List<Int> = sc.nextLine().toCharArray().map { it-'0' }
        height++
        arr.add(line)
        width = line.size
    }

    val visited = Array(height) { BooleanArray(width) }

    val basins = ArrayList<Int>()

    repeat(height) { y ->
        repeat(width) { x ->
            val basin = traverse(arr, visited, x, y)
            if (basin != 0) {
                basins.add(basin)
            }
        }
    }
    println(basins)
    println(basins.sorted().takeLast(3).reduce { acc, i -> acc*i })
}

fun traverse(arr: ArrayList<List<Int>>, visited: Array<BooleanArray>, x: Int, y: Int, currentSize: Int = 0): Int {
    if (visited[y][x] || arr[y][x] == 9) return currentSize
    visited[y][x] = true
    var size = currentSize+1

    val left = arr[y].getOrElse(x-1) {11}
    val right = arr[y].getOrElse(x+1) {11}
    val up = arr.getOrNull(y-1)?.get(x) ?: 11
    val down = arr.getOrNull(y+1)?.get(x) ?: 11

    if (up < 9) {
        size += traverse(arr, visited, x, y-1)
    }
    if (down < 9) {
        size += traverse(arr, visited, x, y+1)
    }
    if (left < 9) {
        size += traverse(arr, visited, x-1, y)
    }
    if (right < 9) {
        size += traverse(arr, visited, x+1, y)
    }
    return size
}