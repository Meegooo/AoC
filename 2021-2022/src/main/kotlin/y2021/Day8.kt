package y2021

import java.io.File
import java.lang.IllegalArgumentException
import java.util.*
import kotlin.collections.HashSet

fun main(args: Array<String>) {
    val sc = Scanner(File("inputs/day8"))

    var star1 = 0;
    var star2 = 0;

    while (sc.hasNext()) {
        val line = sc.nextLine();
        val digits = Array<Set<Char>>(10) { HashSet() }
        val inputs = line.split(" | ").map { it -> it.split(" ").map { it.toCharArray().toSet() } }

        digits[1] = inputs[0].find { it.size == 2 }!!
        digits[4] = inputs[0].find { it.size == 4 }!!
        digits[7] = inputs[0].find { it.size == 3 }!!
        digits[8] = inputs[0].find { it.size == 7 }!!

        digits[6] = inputs[0].find { it.size == 6 && it.intersect(digits[1]).size == 1 } ?: emptySet<Char>().also { println("6 not found") }
        digits[9] = inputs[0].find { it.size == 6 && it.intersect(digits[4]).size == 4 } ?: emptySet<Char>().also { println("7 not found") }
        digits[0] = inputs[0].find { it.size == 6 && it != digits[6] && it != digits[9] } ?: emptySet<Char>().also { println("0 not found") }

        digits[3] = inputs[0].find { it.size == 5 && it.intersect(digits[1]).size == 2 } ?: emptySet<Char>().also { println("3 not found") }
        digits[2] = inputs[0].find { it.size == 5 && it.intersect(digits[9]).size == 4 } ?: emptySet<Char>().also { println("2 not found") }
        digits[5] = inputs[0].find { it.size == 5 && it.intersect(digits[9]).size == 5 } ?: emptySet<Char>().also { println("5 not found") }


        inputs[1].forEach {
            val digit = digits.indexOf(it)
            if (digit == 1 || digit == 4 || digit == 7 || digit == 8) {
                star1++
            }
            if (digit == -1) throw IllegalArgumentException("$it not found for $line\n${digits.joinToString { it.toString() }}")
            else star2+= digit;
        }
    }

    println(star1)
    println(star2)
}