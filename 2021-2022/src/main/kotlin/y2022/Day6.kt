package y2022

import java.util.*
import kotlin.collections.HashMap

fun main() {
    object {}.javaClass.getResource("Day6")?.readText()!!.lines().forEach {
        var count = 0;
        val multiset = HashMap<Char, Int>();
        val queue: Queue<Char> = LinkedList()
        val chars = it.iterator()
        while (multiset.size!=14) {
            val c = chars.nextChar();
            queue.add(c)
            multiset.merge(c, 1) { o1, o2 -> o1+o2 }
            if (queue.size == 14+1) {
                multiset.compute(queue.poll()) {_, value -> if (value == 1) null else value!!-1}
            }
            count++;
        }
        println(count)
    }
}
