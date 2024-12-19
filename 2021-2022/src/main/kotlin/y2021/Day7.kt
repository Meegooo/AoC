package y2021

import java.io.File
import java.util.*
import kotlin.math.abs

fun main(args: Array<String>) {
    val positions = Scanner(File("inputs/day7")).nextLine().split(",").map { it.toInt() }
    val maxPosition = positions.maxOf { it }

    val min = (1..maxPosition).map { it + 1 }
        .associateWith { currentPos ->
            positions.sumOf {
                val dist = abs(it - currentPos)
                dist * (dist+1) / 2
            }
        }.minByOrNull { it.value }!!
    println(min)
}