package y2021

import java.io.File
import java.util.*

private const val considerDiagonals = true

fun main(args: Array<String>) {
    val sc = Scanner(File("inputs/day6-test"))
    val fishies = sc.nextLine().split(",").map { Fishy(it.toInt()) }.toMutableList()
    repeat(256) {
        fishies.addAll(fishies.mapNotNull { it.tick() })
    }
    println(fishies.size)

}

data class Fishy(var timer: Int) {
    fun tick(): Fishy? {
        timer-=1;
        return if (timer == -1) {
            timer = 6
            Fishy(8)
        } else null
    }
}