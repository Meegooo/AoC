package y2021

import java.io.File
import java.math.BigInteger
import java.util.*

fun main(args: Array<String>) {
    val sc = Scanner(File("inputs/day10"))
    var points = 0
    val autocomplete = ArrayList<BigInteger>()
    while (sc.hasNext()) {
        var points2 = BigInteger.ZERO
        val line = sc.nextLine().toCharArray()
        val stack = Stack<Char>()
        var broken = false
        line.forEach { c ->
            if (c in setOf('[', '(', '<', '{')) {
                stack.add(c)
            } else {
                if (c == ')' && stack.pop() != '(') {
                    points += 3
                    broken = true
                }
                else if (c == ']' && stack.pop() != '[') {
                    points += 57
                    broken = true
                }
                else if (c == '}' && stack.pop() != '{') {
                    points += 1197
                    broken = true
                }
                else if (c == '>' && stack.pop() != '<') {
                    points += 25137
                    broken = true
                }
            }
        }

        if (!broken) {
            while (!stack.isEmpty()) {
                val c = stack.pop()
                points2 *= BigInteger("5");
                points2 += when (c) {
                    '(' -> BigInteger("1")
                    '[' -> BigInteger("2")
                    '{' -> BigInteger("3")
                    '<' -> BigInteger("4")
                    else -> throw RuntimeException("Uwot")
                }
            }
            autocomplete.add(points2)
        }


    }
    println(points)
    autocomplete.sort()
    println(autocomplete)
    println(autocomplete[autocomplete.size/2])
    //NO 2138262343
}