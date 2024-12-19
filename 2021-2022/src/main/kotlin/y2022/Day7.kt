package y2022

import kotlin.math.min


fun main() {
    val root = Node(-1L, null, "")
    var pwd = root
    object {}.javaClass.getResource("Day7")?.readText()!!.split("$").forEach { s: String ->
        if (s.isBlank()) return@forEach
        val command = s.lineSequence().first().trim()
        when {
            command == "cd /" -> pwd = root
            command == "cd .." -> pwd = pwd.parent!!
            command.startsWith("cd") -> pwd = pwd.children[command.split(" ")[1]]!!
            command.startsWith("ls") -> {
                s.trim().lineSequence().drop(1).forEach {
                    val (size, filename) = it.split(" ");
                    if (size == "dir") {
                        pwd.addFile(-1, filename)
                    } else {
                        pwd.addFile(size.toLong(), filename)
                    }
                }
            }
        }
    }

    println(part1(root))
    println(part2(root))
}

fun part1(root: Node): Long {
    if (root.fileSize >= 0) {
        return 0
    }
    val recurse = root.children.values.sumOf { part1(it) }
    return if (root.size > 100000) {
        recurse
    } else {
        root.size + recurse
    }
}

fun part2(root: Node): Long {
    val filesystemSize = 70000000
    val usedSpace = root.size
    val needSpace = 30000000
    val goal = needSpace - (filesystemSize-usedSpace)
    var smallest = Long.MAX_VALUE;
    fun recurse(current: Node) {
        if (current.fileSize > 0 || current.size < goal) {
            return
        }
        smallest = min(smallest, current.size);
        current.children.values.forEach { recurse(it) }
    }
    recurse(root)
    return smallest
}

//9725484

class Node(val fileSize: Long, val parent: Node?, val name: String) {
    val size: Long by lazy {
            if (fileSize == -1L) {
                children.values.sumOf { it.size }
            } else {
                fileSize;
            }
        }

    val children: MutableMap<String, Node> = HashMap()

    fun addFile(size: Long, name: String) {
        val n = Node(size, this, name)
        children[name] = n
    }

    override fun toString(): String {
        return "Node(fileSize=$fileSize, name='$name', children=$children)"
    }

}