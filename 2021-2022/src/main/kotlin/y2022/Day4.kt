package y2022

fun main() =
    println("${f { a, b -> a.containsAll(b) || b.containsAll(a) }} ${f { a, b -> a.intersect(b).isNotEmpty() }}")

fun f(r: (Set<Int>, Set<Int>) -> Boolean) = object {}.javaClass.getResource("Day4")?.readText()!!.lines().count { s ->
    s.split(",", "-").chunked(2).map { (it[0].toInt()..it[1].toInt()).asIterable().toSet() }
        .zipWithNext(r)[0]
}