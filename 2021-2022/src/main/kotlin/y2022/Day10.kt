package y2022

import kotlin.math.abs

var cycles = 0
var register = 1
var score = 0

fun main() {
    object {}.javaClass.getResource("Day10")?.readText()!!.lines().forEach {
        if (it == "noop") nextCycle()
        else if (it.startsWith("addx")) {
            nextCycle();
            nextCycle()
            register += it.split(" ").last().toInt()
        }
    }
    println(score)
}

fun nextCycle() {
    if (abs(register-cycles%40) <= 1) print("██")
    else print("  ")
    if (cycles % 40 == 39) println()
    cycles++
    if ((cycles-20) % 40 == 0) {
        score += cycles * register
    }
}