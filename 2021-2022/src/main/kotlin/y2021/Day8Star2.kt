package y2021

import java.io.File
import java.util.*

fun main(args: Array<String>) {
    val sc = Scanner(File("inputs/day8"))

    var star1 = 0;
    var star2 = 0;

    val digitsMap = mapOf(
        setOf('a', 'b', 'c', 'e', 'f', 'g') to 0,
        setOf('c', 'f') to 1,
        setOf('a', 'c', 'd', 'e', 'g') to 2,
        setOf('a', 'c', 'd', 'f', 'g') to 3,
        setOf('b', 'c', 'd', 'f') to 4,
        setOf('a', 'b', 'd', 'f', 'g') to 5,
        setOf('a', 'b', 'd', 'e', 'f', 'g') to 6,
        setOf('a', 'c', 'f') to 7,
        setOf('a', 'b', 'c', 'd', 'e', 'f', 'g') to 8,
        setOf('a', 'b', 'c', 'd', 'f', 'g') to 9,
    )

    while (sc.hasNext()) {
        val line = sc.nextLine();
        val digits = Array<Set<Char>>(10) { HashSet() }
        val inputs = line.split(" | ").map { it -> it.split(" ").map { it.toCharArray().toSet() } }
        mappings().forEach { mapping ->
            val correctMapping = inputs[0].map { set -> set.map { mapping[it] }.toSet() }.all { digitsMap.containsKey(it) }
            if (correctMapping) {
                var currentNum = 0
                inputs[1].forEach { set ->
                    val mapped = set.map { mapping[it] }.toSet()
                    val digit = digitsMap[mapped]!!
                    if (digit == 1 || digit == 4 || digit == 7 || digit == 8) {
                        star1++
                    }
                    if (digit == -1) throw IllegalArgumentException("$mapped not found for $line\n${digits.joinToString { it.toString() }}")
                    else {
                        currentNum *= 10
                        currentNum += digit
                    };
                }
                star2 += currentNum
            }
        }



    }

    println(star1)
    println(star2)
}

fun mappings(): Sequence<Map<Char, Char>> {
    val arrayStart = (1..7).toList().toIntArray()
    val array = (1..7).toList().toIntArray()
    val letters = mapOf(1 to 'a', 2 to 'b', 3 to 'c', 4 to 'd', 5 to 'e', 6 to 'f', 7 to 'g')
    return sequence {
        yield(mapOf('a' to 'a', 'b' to 'b', 'c' to 'c', 'd' to 'd', 'e' to 'e', 'f' to 'f', 'g' to 'g'))
        Permutations.nextPermutation(array)
        while (!array.contentEquals(arrayStart)) {
            yield(array.mapIndexed { index, it -> letters[it]!! to letters[index + 1]!! }.toMap())
            Permutations.nextPermutation(array)
        }

    }

}


object Permutations {
    fun nextPermutation(nums: IntArray) {
        var i = nums.size - 2
        while (i >= 0 && nums[i + 1] <= nums[i]) {
            i--
        }
        if (i >= 0) {
            var j = nums.size - 1
            while (nums[j] <= nums[i]) {
                j--
            }
            swap(nums, i, j)
        }
        reverse(nums, i + 1)
    }

    private fun reverse(nums: IntArray, start: Int) {
        var i = start
        var j = nums.size - 1
        while (i < j) {
            swap(nums, i, j)
            i++
            j--
        }
    }

    private fun swap(nums: IntArray, i: Int, j: Int) {
        val temp = nums[i]
        nums[i] = nums[j]
        nums[j] = temp
    }
}