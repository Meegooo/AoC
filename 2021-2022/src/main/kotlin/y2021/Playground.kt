package y2021

fun main() {
    while (true) {
        val fishies = readLine()!!.split(",").map { it.toInt() }
            .groupingBy { it }.eachCount()
            .entries.sortedBy { it.key }.map { "${it.key}->${it.value}" }
        println(fishies)
    }
}