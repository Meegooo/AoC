package y2022

import kotlin.collections.HashMap
import kotlin.math.max
import kotlin.math.min


fun main() {
    val map = DefaultHashMap<Pair<Int, Int>, Boolean>(false)
    object {}.javaClass.getResource("Day14")!!.readText().lines().forEach { line ->
        line.split(" -> ").map { pair -> pair.split(",").map(String::toInt).let { it[0] to it[1] } }.zipWithNext { l, r ->
            if (l.first == r.first) {
                for (col in min(l.second, r.second) .. max(l.second, r.second)) {
                    map[l.first to col] = true
                }
            } else if (l.second == r.second) {
                for (row in min(l.first, r.first) .. max(l.first, r.first)) {
                    map[row to l.second] = true
                }
            }
        }
    }

    val mapWithSand = DefaultHashMap(map)
    val maxY = map.keys.maxOf { it.second }
    var sandCount = 0;
    outer@while (true) {
        //drop sand
        var sandPos = 500 to 0
        inner@while (true) {
            val newSand = sandPos.move(mapWithSand, maxY+2) //remove maxY+2 for part1
            if (newSand == null) {
                mapWithSand[sandPos] = true
                break@inner
            } else {
//                if (sandPos.second >= maxY) break@outer
                sandPos = newSand
            }
        }
        sandCount++
        if (sandPos == 500 to 0 ) break@outer

        for (row in 0..9) {
            for (col in 494..503) {
                if (map[col to row]) print('#')
                else if (mapWithSand[col to row]) print('o')
                else print('.')
            }
            println()
        }
        println("-------------------")
    }
    println(sandCount)

}


private fun Pair<Int, Int>.move(map: Map<Pair<Int, Int>, Boolean>, maxY: Int = Integer.MAX_VALUE): Pair<Int, Int>? {
    return if (this.second+1 == maxY) {
        null
    } else  if (!map.containsKey(this.first to this.second+1)) {
        this.first to this.second+1
    } else if (!map.containsKey(this.first-1 to this.second+1)) {
        this.first-1 to this.second+1
    } else if (!map.containsKey(this.first+1 to this.second+1)) {
        this.first+1 to this.second+1
    } else {
        null
    }
}

class DefaultHashMap<K, V> : HashMap<K, V> {
    private val defaultValue: V

    constructor(defaultValue: V) : super() {
        this.defaultValue = defaultValue
    }

    constructor(map: DefaultHashMap<out K, out V>) : super(map) {
        this.defaultValue = map.defaultValue
    }
    override fun get(key: K): V {
        return if (super.containsKey(key)) super.get(key)!!
        else defaultValue;
    }
}
