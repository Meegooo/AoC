package y2021
import java.io.File
import kotlin.collections.ArrayList

fun main(args: Array<String>) {
    val lines = File("inputs/day3").readLines()
    var oxygen:List<String> = ArrayList(lines)
    var co2:List<String> = ArrayList(lines)
    val digits = lines[0].length
    repeat(digits) { digit ->
        oxygen = filter(oxygen, digit, false)
        co2 = filter(co2, digit, true)
    }

    val oxygenValue = oxygen[0].toInt(2)
    val co2Value = co2[0].toInt(2)
    println("$oxygenValue, $co2Value, ${oxygenValue*co2Value}")
}


fun filter(lines: List<String>, bitIdx: Int, invert: Boolean): List<String> {
    if (lines.size == 1) return lines
    val sum = lines.sumOf { it[bitIdx] - '0' }
    val required = if (!invert) {
        if (sum >= lines.size - sum) '1' else '0'
    } else {
        if (sum >= lines.size - sum) '0' else '1'
    }

    for (i in 1 downTo  -10) {

    }

    (1..2).forEach {  }
    return lines.filter { it[bitIdx] == required }
}