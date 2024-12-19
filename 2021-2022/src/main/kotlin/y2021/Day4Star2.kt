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
        val iterator = boards.iterator()
        while (iterator.hasNext()) {
            val board = iterator.next()
            val winner = board.mark(num)
            if (winner != -1 && boards.size > 1) {
                iterator.remove()
            } else if (winner != -1) {
                println(winner)
                return
            }
        }
    }

}