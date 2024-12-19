package y2021
import java.io.File
import java.util.*

fun main(args: Array<String>) {
    val sc = Scanner(File("inputs/day2"))
    var x = 0
    var y = 0
    var aim = 0
    while (sc.hasNext()) {
        val command = sc.nextLine().split(" ")
        when (command[0]) {
            "forward" -> {
                x += command[1].toInt()
                y += command[1].toInt()*aim
            }
            "down" -> aim += command[1].toInt()
            "up" -> aim -= command[1].toInt()
        }
    }

    println("$x $y ${x*y}")
}