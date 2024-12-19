package y2021

import org.jgrapht.Graph
import org.jgrapht.Graphs
import org.jgrapht.graph.DefaultEdge
import org.jgrapht.graph.DefaultUndirectedGraph
import java.io.File
import java.util.*
import kotlin.collections.HashSet


fun main(args: Array<String>) {
    val sc = Scanner(File("inputs/day12"))
    val g: Graph<String, DefaultEdge> = DefaultUndirectedGraph(DefaultEdge::class.java)
    while (sc.hasNext()) {
        val line = sc.nextLine().split("-")
        if (!g.containsVertex(line[0])) {
            g.addVertex(line[0])
        }
        if (!g.containsVertex(line[1])) {
            g.addVertex(line[1])
        }
        g.addEdge(line[0], line[1])
    }

    val visited = HashSet<String>()
    val allRoutes = HashSet<List<String>>()

    val currentRoute = ArrayList<String>()
    traverse("start", "end", visited, g, currentRoute, allRoutes)

    allRoutes.forEach {
        println(it)
    }

    println(allRoutes.size)


}

fun traverse(from: String, to: String, visited: Set<String>, g: Graph<String, DefaultEdge>, currentRoute: List<String>, allRoutes: HashSet<List<String>>, twice: Boolean = false) {
    if (from == to) {
        allRoutes.add(currentRoute)
        return
    }
    val newVisited = HashSet(visited)
    if (from.lowercase() == from) {
        newVisited.add(from)
    }
    val neighbours = Graphs.neighborListOf(g, from)
    for (vertex in neighbours) {
        if (vertex !in visited) {
            traverse(vertex, "end", newVisited, g, currentRoute + vertex, allRoutes, twice)
        } else if (!twice && vertex != "start" && vertex.lowercase() == vertex) {
            traverse(vertex, "end", newVisited, g, currentRoute + vertex, allRoutes, true)
        }
    }

}