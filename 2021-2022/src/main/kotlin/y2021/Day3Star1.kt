package y2021
import java.io.File

fun main(args: Array<String>) {
    val lines = File("inputs/day3").readLines()
    val digits = lines[0].length
    var gamma = 0
    repeat(digits) { digit ->
        val sum = lines.sumOf { it[digit] - '0' }
        gamma *= 2;
        if (sum > lines.size - sum) {
            gamma += 1
        }
    }
    var epsilon = gamma.inv().and(1.shl(digits)-1)



    println("$gamma $epsilon ${gamma*epsilon}")
}