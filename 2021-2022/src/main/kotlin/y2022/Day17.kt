package y2022

import java.lang.IllegalArgumentException
import java.util.*


fun main() {
    Day17.part1()
}

object Day17 {

    var blocksMaterialized = 0
    fun part1() {
        val blocks = arrayOf(
            Block(Array(1) { BooleanArray(4) { true } }),
            Block(arrayOf(booleanArrayOf(false, true, false), booleanArrayOf(true, true, true), booleanArrayOf(false, true, false))),
            Block(arrayOf(booleanArrayOf(false, false, true), booleanArrayOf(false, false, true), booleanArrayOf(true, true, true))),
            Block(Array(4) { BooleanArray(1) { true } }),
            Block(Array(2) { BooleanArray(2) { true } })
        )

//        println(blocks.joinToString(separator = "\n\n"))

        val field = Field(javaClass.getResource(javaClass.simpleName)!!.readText().toCharArray())
        var blockIdx = 0
        field.spawnBlock(blocks[blockIdx])
        while (blocksMaterialized < 2022) {
            if (field.tick()) {
                blocksMaterialized++
                blockIdx++
                if (blockIdx == blocks.size) blockIdx = 0
                field.spawnBlock(blocks[blockIdx])
            }
        }
        println(field.height())
    }

    class Block(val shape: Array<BooleanArray>) {
        override fun toString(): String {
            return shape.joinToString(separator = "\n") {it.joinToString(separator = "") { if (it) "#" else "." }}
        }

        fun width() = shape[0].size
        fun height() = shape.size
        operator fun get(height: Int): BooleanArray? {
            if (height()-1-height < 0 || height()-1-height >= height()) return null
            return shape[height()-1-height]
        }
    }
    class Field(val airflowSequence: CharArray) {
        private val field = LinkedList<BooleanArray>()
        private var currentBlock: Block? = null
        private var blockBottomLeftOffset: Pair<Int, Int>? = null
        private var nextFlowTick: Int = 0
        private var nextDown = false;
        init {
            field.add(BooleanArray(7) {true})
        }
        fun spawnBlock(block: Block) {
            currentBlock = block
            blockBottomLeftOffset = 2 to 3
            println(blocksMaterialized)
            printField()
        }

        fun height() = field.size

        fun tick(): Boolean {
            if (blocksMaterialized == 9) {
                printField()
            }
            if (nextDown) {
                val newOffset = blockBottomLeftOffset!!.first to blockBottomLeftOffset!!.second-1
                if (collisionFits(newOffset)) {
                    blockBottomLeftOffset = newOffset
                } else {
                    materialize()
                    currentBlock = null
                    blockBottomLeftOffset = null
                    nextDown = !nextDown
                    return true
                }
            } else {
                val direction = airflowSequence[nextFlowTick++]
                if (nextFlowTick == airflowSequence.size) nextFlowTick = 0
                val newOffset = if (direction == '<') {
                    blockBottomLeftOffset!!.first-1 to blockBottomLeftOffset!!.second
                } else if (direction == '>') {
                    blockBottomLeftOffset!!.first+1 to blockBottomLeftOffset!!.second
                } else throw IllegalArgumentException()
                if (collisionFits(newOffset)) {
                    blockBottomLeftOffset = newOffset
                }
            }
            nextDown = !nextDown
            return false
        }

        private fun collisionFits(offset: Pair<Int, Int>): Boolean {
            if (offset.first !in 0 .. 7 - currentBlock!!.width())
                return false
            if (offset.second >= 0)
                return true

            val checkDepth = -offset.second
            val iterator = field.iterator()
            for (currentDepth in checkDepth-1 downTo  0) {
                if (currentDepth < currentBlock!!.height()) {
                    val fieldRow = iterator.next()
                    val blockRow = currentBlock!![currentDepth]
                    blockRow?.forEachIndexed { index, b ->
                        if (b && fieldRow[index + offset.first]) return false
                    }
                }
            }
            return true
        }

        private fun materialize() {
            val checkDepth = -blockBottomLeftOffset!!.second
            val iterator = field.iterator()
            for (currentDepth in checkDepth-1 downTo  0) {
                val fieldRow = iterator.next()
                val blockRow = currentBlock!![currentDepth]
                blockRow.forEachIndexed { index, b ->
                    fieldRow[index + blockBottomLeftOffset!!.first] = b || fieldRow[index + blockBottomLeftOffset!!.first]
                }
            }

            val remainingHeight = currentBlock!!.height()-checkDepth
            for (currentDepth in 0 until remainingHeight) {
                val newRow = BooleanArray(7) {false}
                val blockRow = currentBlock!![checkDepth+currentDepth]
                for (index in blockRow.indices) {
                    newRow[index + blockBottomLeftOffset!!.first] = blockRow[index]
                }
                field.addFirst(newRow)
            }
        }

        fun printField() {
            val currentBlockCopy = currentBlock!!
            if (blockBottomLeftOffset!!.second >= 0) {
                for (y in 0 until currentBlockCopy.height()) {
                    for (x in 0 until 7) {
                        if (x < blockBottomLeftOffset!!.first || x >= blockBottomLeftOffset!!.first+ currentBlockCopy.width()) print(".")
                        else if (currentBlockCopy[currentBlockCopy.height()-y-1][x-blockBottomLeftOffset!!.first]) print("@")
                        else print(".")
                    }
                    println()
                }
                for (i in 0 until blockBottomLeftOffset!!.second) {
                    println(".......")
                }
                field.iterator().forEach {
                    println(it.joinToString(separator = "") { if (it) "#" else "." })
                }
            } else {
                for (y in 0 until currentBlockCopy.height()+blockBottomLeftOffset!!.second) {
                    for (x in 0 until 7) {
                        if (x < blockBottomLeftOffset!!.first || x >= blockBottomLeftOffset!!.first+ currentBlockCopy.width()) print(".")
                        else if (currentBlockCopy[currentBlockCopy.height()-y-1][x-blockBottomLeftOffset!!.first]) print("@")
                        else print(".")
                    }
                    println()
                }

                val checkDepth = -blockBottomLeftOffset!!.second
                val iterator = field.iterator()
                for (currentDepth in checkDepth - 1 downTo  0) {
                    val fieldRow = iterator.next()
                    val blockRow = currentBlock!![currentDepth]
                    fieldRow.forEachIndexed { index, b ->
                        if (b) print("#")
                        else if (index - blockBottomLeftOffset!!.first in blockRow.indices && blockRow[index - blockBottomLeftOffset!!.first]) print("@")
                        else print(".")
                    }
                    println()
                }
                while (iterator.hasNext()) {
                    val fieldRow = iterator.next()
                    println(fieldRow.joinToString(separator = "") { if (it) "#" else "." })
                }

            }


            println()
        }
    }
}

//2283 too low