package y2021

import java.io.File
import java.util.*
import kotlin.collections.HashMap


fun main(args: Array<String>) {
    val sc = Scanner(File("inputs/day14"))
    val rules = HashMap<String, Char>()
    var pairs = HashMap<String, Long>()

    val firstLine = sc.nextLine().toCharArray()
    for (i in firstLine.indices) {
        val key = String(charArrayOf(firstLine.getOrElse(i-1) {' '}, firstLine[i]))
        pairs[key] = pairs.getOrDefault(key, 0) + 1
    }

    sc.nextLine()
    while (sc.hasNext()) {
        val line = sc.nextLine();
        val key = line.take(2)
        val value = line.takeLast(1).toCharArray()[0]
        rules[key] = value
    }

    repeat(40) {
        val newPairs = HashMap<String, Long>()
        for ((pair, count) in pairs) {
            val center = rules[pair]
            if (center != null) {
                val polymerLeft = String(charArrayOf(pair[0], center))
                val polymerRight = String(charArrayOf(center, pair[1]))
                newPairs[polymerLeft] = newPairs.getOrDefault(polymerLeft, 0) + count
                newPairs[polymerRight] = newPairs.getOrDefault(polymerRight, 0) + count
            } else {
                newPairs[pair] = count
            }
        }
        pairs = newPairs
    }

    val counts = pairs.entries.map { it.key[1] to it.value }.groupingBy { it.first }.fold(0L) {acc, (_, value) -> acc + value}
    val max = counts.values.maxOrNull() ?: 0
    val min = counts.values.minOrNull() ?: 0

    println(counts)
    println(max-min)



}