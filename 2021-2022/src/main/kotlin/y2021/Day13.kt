package y2021

import java.io.File
import java.util.*
import kotlin.collections.HashSet


fun main(args: Array<String>) {
    val sc = Scanner(File("inputs/day13"))
    var dots = HashSet<Pair<Int, Int>>()
    var folds = false
    while (sc.hasNext()) {
        val line = sc.nextLine()
        if (line.isBlank()) {
            folds = true
        } else if (!folds) {
            val split = line.split(",").map { it.toInt() }
            dots.add(split[0] to split[1])
        } else {
            val rule = line.replace("fold along ", "")
            val num = rule.drop(2).toInt()
            val newDots = HashSet<Pair<Int, Int>>()
            if (rule.startsWith("x=")) {
                for (dot in dots) {
                    val newY = dot.second
                    val newX = if (dot.first > num) {
                        val d = dot.first - num
                        num - d
                    } else dot.first
                    newDots.add(newX to newY)
                }
            } else { //y
                for (dot in dots) {
                    val newX = dot.first
                    val newY = if (dot.second > num) {
                        val d = dot.second - num
                        num - d
                    } else dot.second
                    newDots.add(newX to newY)
                }
            }
            dots = newDots
        }
    }

    val maxX = dots.maxOf { it.first }
    val maxY = dots.maxOf { it.second }

    repeat(maxY+1) { y ->
        repeat(maxX+1) { x ->
            if (dots.contains(x to y)) {
                print("#")
            } else {
                print(".")
            }
        }
        println()
    }



}