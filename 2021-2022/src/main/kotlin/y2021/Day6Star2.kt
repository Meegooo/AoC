package y2021

import java.io.File
import java.math.BigInteger
import java.util.*

fun main(args: Array<String>) {
    val sc = Scanner(File("inputs/day6"))
    val fishies = sc.nextLine().split(",").map { it }.groupingBy { it.toInt() }.eachCount().mapValues { it.value.toBigInteger() }.toMutableMap()
    repeat(7) {
        fishies.putIfAbsent(it, BigInteger.ZERO)
    }

    var currentDay = 0
    var newFishiesYesterday = BigInteger.ZERO
    var newFishiesToday = BigInteger.ZERO
    repeat(256) {
        currentDay = (currentDay + 1) % 7
        val new = fishies[currentDay]!!
        fishies[currentDay] = fishies[currentDay]!! + newFishiesYesterday;
//        println("${fishies.values.sum() + new}, $currentDay")
        newFishiesYesterday = newFishiesToday
        newFishiesToday = new
        println("$currentDay=" + fishies.entries.sortedBy { it.key }.map { "${it.key}->${it.value}" }.toString() + " [$new]")
    }


    print(fishies.values.reduce { acc, bigInteger -> acc.add(bigInteger) } + newFishiesYesterday)
}