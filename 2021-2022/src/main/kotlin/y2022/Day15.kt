package y2022

import kotlin.math.abs
import kotlin.math.max

fun main() {
    Day15.part2()
}

object Day15 {
    fun part1() {
        val y = 2000000
        var xOfBeaconsAtY = HashSet<Int>()
        val ranges = javaClass.getResource(javaClass.simpleName)!!.readText().lines().map { line ->
            val match =
                "Sensor at x=(-?\\d+), y=(-?\\d+): closest beacon is at x=(-?\\d+), y=(-?\\d+)".toRegex().find(line)
            val (sensor, beacon) = match!!.groupValues.drop(1).map { it.toInt() }.chunked(2).map { it[0] to it[1] }
            if (beacon.second == y) xOfBeaconsAtY.add(beacon.first)
            val distance = sensor.manhattanDistanceTo(beacon)
            val distanceToY = abs(y - sensor.second)
            sensor.first - (distance - distanceToY)..sensor.first + distance - distanceToY
        }.sortedWith(Comparator.comparingInt<IntRange> { it.first }.thenComparingInt { it.last })

        val combinedRanges = ArrayList<IntRange>()
        var currentRange = ranges[0]
        for (i in 1 until ranges.size) {
            val right = ranges[i]
            if (currentRange.last >= right.first) {
                currentRange = currentRange.first..max(currentRange.last, right.last)
            } else {
                combinedRanges.add(currentRange)
                currentRange = right
            }
        }
        combinedRanges.add(currentRange)

        println(ranges)
        println(combinedRanges)

        println(combinedRanges.map { it.last - it.first + 1 }.sum() - xOfBeaconsAtY.size)
    }


    fun part2() {
        val edge = 4000000
        val sensorBeacons = object {}.javaClass.getResource("Day15")!!.readText().lines().map { line ->
            val match =
                "Sensor at x=(-?\\d+), y=(-?\\d+): closest beacon is at x=(-?\\d+), y=(-?\\d+)".toRegex().find(line)
            val (sensor, beacon) = match!!.groupValues.drop(1).map { it.toInt() }.chunked(2).map { it[0] to it[1] }
            sensor to beacon
        }

        println(sensorBeacons.map {it.first.manhattanDistanceTo(it.second) }.sum()*4)
        val candidates = sensorBeacons.asSequence().flatMap { (sensor, beacon) ->
            println("Sensor: $sensor, beacon: $beacon")
            getPointsOfDistanceFrom(sensor, sensor.manhattanDistanceTo(beacon) + 1)
        }.filter { it.first in 0..edge && it.second in 0..edge }.toSet()

        println(candidates.size)
            candidates.forEach { candidate ->
                if (sensorBeacons.all { (sensor, beacon) ->
                        sensor.manhattanDistanceTo(beacon) < sensor.manhattanDistanceTo(
                            candidate
                        )
                    }) {
                    println(candidate)
                    println(candidate.first * 4000000 + candidate.second)
                }
            }


//
//        for (x in 0..20) {
//            for (y in 0..20) {
//                if (x to y == sensorBeacons[0].first) print("S")
//                else if (x to y == sensorBeacons[0].second) print("B")
//                else if (candidates.contains(x to y)) print("#")
//                else print(" ")
//            }
//            println()
//        }
    }

    fun getPointsOfDistanceFrom(from: Pair<Int, Int>, distance: Int): Set<Pair<Int, Int>> {
        val candidates = HashSet<Pair<Int, Int>>()
        var dx = distance
        var dy = 0
        candidates.add(dx + from.first to dy + from.second)
        //bottom left
        while (dx > 0) {
            dx--
            dy++
            candidates.add(dx + from.first to dy + from.second)
        }
        //bottom right
        while (dy > 0) {
            dx--
            dy--
            candidates.add(dx + from.first to dy + from.second)
        }
        //top right
        while (dx < 0) {
            dx++
            dy--
            candidates.add(dx + from.first to dy + from.second)
        }
        //top left
        while (dy < 0) {
            dx++
            dy++
            candidates.add(dx + from.first to dy + from.second)
        }
        return candidates
    }

    private fun Pair<Int, Int>.manhattanDistanceTo(second: Pair<Int, Int>): Int {
        return abs(this.first - second.first) + abs(this.second - second.second)
    }
}
