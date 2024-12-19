package y2021

import java.io.File
import java.util.*

fun main(args: Array<String>) {
    val sc = Scanner(File("inputs/day9"))
    val arr = ArrayList<List<Int>>()
    var width = 0
    var height = 0
    var answer = 0;
    while (sc.hasNext()) {
        val line: List<Int> = sc.nextLine().toCharArray().map { it-'0' }
        height++
        arr.add(line)
        width = line.size
    }

    repeat(width) { y ->
        repeat(height) { x ->
            val current = arr[x][y]
            val up = arr[x].getOrElse(y-1) {11}
            val down = arr[x].getOrElse(y+1) {11}
            val left = arr.getOrNull(x-1)?.get(y) ?: 11
            val right = arr.getOrNull(x+1)?.get(y) ?: 11
            if (current < up && current < down && current < left && current < right) {
                answer+=current+1
            }
        }
    }

    println(answer)
}