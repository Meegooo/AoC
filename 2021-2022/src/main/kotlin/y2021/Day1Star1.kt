package y2021

import java.io.File
import java.util.*

fun main() {
    val sc = Scanner(File("inputs/day1"))
    var p3 = sc.nextLine().toInt()
    var p2 = sc.nextLine().toInt()
    var p1 = sc.nextLine().toInt()
    var previous = p1+p2+p3
    var counter = 0;
    while (sc.hasNext()) {
        val input = sc.nextLine().toInt()
        val current = input + p1 + p2
        if (current > previous) {
            counter++
        }
        previous = current;
        p3 = p2
        p2 = p1
        p1 = input
    }

    println(counter)
}