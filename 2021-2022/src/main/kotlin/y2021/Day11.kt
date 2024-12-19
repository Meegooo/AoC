package y2021

import java.io.File
import java.util.*

val flashed = HashSet<Pair<Int, Int>>()
fun main(args: Array<String>) {
    val sc = Scanner(File("inputs/day11"))
    val dim = 10
    val octopi = Array(dim) {IntArray(dim)}
    repeat(dim) {row ->
        val input = sc.nextLine().toCharArray().map { it - '0' }
        repeat(dim) { col ->
            octopi[row][col] = input[col]
        }
    }


    var steps = 0
    var activations = 0
    var previousActivations = 0
    var previousStepActivations = 0
    while(previousStepActivations < 100) {
        previousStepActivations = 0
        steps += 1
        flashed.clear()
        //Increase by 1
        repeat(dim) { row ->
            repeat(dim) { col ->
                increase(octopi, row, col)
            }
        }
        //Flash
        do {
            activations += previousActivations
            previousStepActivations += previousActivations
            previousActivations = 0
            repeat(dim) { row ->
                repeat(dim) { col ->
                    previousActivations += activate(octopi, row, col)
                }
            }
        } while (previousActivations > 0)

        flashed.forEach { octopi[it.first][it.second] = 0 }
    }

    println(activations)
    println(octopi.joinToString(separator = "\n") { it.joinToString(separator = "") { it.toString() } })
    println(steps)
}

fun activate(octopi: Array<IntArray>, row: Int, col: Int): Int {
    if (row < 0 || col < 0 || row >= octopi.size || col >= octopi[0].size) return 0
    if (octopi[row][col] > 9) {
        octopi[row][col] -= 10
        flashed.add(row to col)

        increase(octopi,row-1, col-1)
        increase(octopi,row+0, col-1)
        increase(octopi,row+1, col-1)

        increase(octopi,row-1, col+1)
        increase(octopi,row+0, col+1)
        increase(octopi,row+1, col+1)

        increase(octopi,row-1, col+0)
        increase(octopi,row+1, col+0)
        return 1
    }
    return 0
}

fun increase(octopi: Array<IntArray>, row: Int, col: Int) {
    if (row < 0 || col < 0 || row >= octopi.size || col >= octopi[0].size || (row to col) in flashed) return
    octopi[row][col] += 1
}