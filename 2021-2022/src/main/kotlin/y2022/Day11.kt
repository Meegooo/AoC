package y2022

import java.util.*

fun main() {
    val text = object {}.javaClass.getResource("Day11")!!.readText()
    val monkeys = text.split("\n\n").map(Monkey::construct)

    val newDivisor = monkeys.map { it.test }.reduce(Long::times)
    println(newDivisor)
    repeat(10000) {
        for (monkey in monkeys) {
            while (monkey.hasNextItem()) {
                val (newItem, newMonkey) = monkey.inspectNextItem2(newDivisor)
                monkeys[newMonkey.toInt()].addItem(newItem)
            }
        }
    }
    monkeys.forEach { println(it.inspections) }
    println("------------")
    println(monkeys.map { it.inspections.toLong() }.sortedDescending().take(2).reduce(Long::times))

}

fun printMonkeys(monkeys: List<Monkey>) = monkeys.forEach { println(it) }

data class Monkey(val items: Queue<Long>, val operation: (Long) -> Long, val test: Long, val throwIfTrue: Long, val throwIfFalse: Long) {

    var inspections = 0

    companion object {
        fun construct(input: String): Monkey {
            val lines = input.lines()
            val items = LinkedList(lines[1].trim().removePrefix("Starting items: ").split(", ").map(String::toLong));
            val (leftOperand, operator, rightOperand) = lines[2].trim().removePrefix("Operation: new = ").split(" ")
            val operatorFunc: Long.(Long) -> Long = if (operator == "+") Long::plus else if (operator == "*") Long::times else throw IllegalArgumentException()
            val operation = {i: Long ->
                val l = if (leftOperand == "old") i else leftOperand.toLong()
                val r = if (rightOperand == "old") i else rightOperand.toLong()
                operatorFunc(l, r)
            }

            val test = lines[3].trim().removePrefix("Test: divisible by ").toLong();
            val ifTrue = lines[4].trim().removePrefix("If true: throw to monkey ").toLong()
            val ifFalse = lines[5].trim().removePrefix("If false: throw to monkey ").toLong()
            return Monkey(items, operation, test, ifTrue, ifFalse)
        }
    }

    fun hasNextItem(): Boolean {
        return items.isNotEmpty()
    }
    fun inspectNextItem(): Pair<Long, Long> {
        inspections++
        val item = items.poll()
        val newItem = (operation(item) / 3)
        return newItem to if (newItem % test == 0L) throwIfTrue else throwIfFalse
    }

    fun inspectNextItem2(divisor: Long): Pair<Long, Long> {
        inspections++
        val item = items.poll()
        val newItem = (operation(item) % divisor)
        return newItem to if (newItem % test == 0L) throwIfTrue else throwIfFalse
    }

    fun addItem(item: Long) {
        items.add(item)
    }
}