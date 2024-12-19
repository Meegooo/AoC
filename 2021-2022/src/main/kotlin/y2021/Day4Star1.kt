package y2021
import java.io.File
import java.util.*
import kotlin.collections.ArrayList

fun main(args: Array<String>) {
    val sc = Scanner(File("inputs/day4"))
    val numbers = sc.nextLine().split(",").map { it.toInt() }
    val boards = ArrayList<BingoBoard>()
    while (sc.hasNext()) {
        sc.nextLine()
        val board = Array(5) { IntArray(5) }
        repeat(5) {
            val line = sc.nextLine().split(" ").filter { it.isNotEmpty() }.map { it.toInt() }.toIntArray()
            board[it] = line
        }

        boards.add(BingoBoard(board))
    }

    numbers.forEach { num ->
        boards.forEach { board ->
            val winner = board.mark(num)
            if (winner != -1) {
                println(winner)
                return
            }
        }
    }
}

data class BingoBoard(val board: Array<IntArray>) {
    val marked = Array(5) { BooleanArray(5) { false } }
    fun mark(called: Int): Int {
        repeat(5) { row ->
            repeat(5) { column ->
                if (board[row][column] == called) {
                    marked[row][column] = true
                }
            }
        }
        return check(called)
    }

    fun check(called: Int): Int {
        repeat(5) { row ->
            var brokeOut = false
            for (column in 0..4) {
                if (!marked[row][column]) {
                    brokeOut = true
                    break
                }
            }
            if (!brokeOut) { //WINNER
                return calculate(called)
            }
        }
        repeat(5) { column ->
            var brokeOut = false
            for (row in 0..4) {
                if (!marked[row][column]) {
                    brokeOut = true
                    break
                }
            }
            if (!brokeOut) { //WINNER
                return calculate(called)
            }
        }
        return -1
    }

    fun calculate(called: Int): Int {
        //Sum of all unmarked numbers
        var unmarkedSum = 0;
        repeat(5) { row ->
            repeat(5) { column ->
                if (!marked[row][column]) {
                    unmarkedSum += board[row][column]
                }
            }
        }
        return unmarkedSum * called
    }
}