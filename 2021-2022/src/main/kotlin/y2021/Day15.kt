package y2021

import org.jgrapht.Graph
import org.jgrapht.alg.shortestpath.DijkstraShortestPath
import org.jgrapht.graph.*
import java.io.File
import java.util.*
import kotlin.collections.ArrayList


fun main(args: Array<String>) {
    val sc = Scanner(File("inputs/day15"))
    val g: Graph<Pair<Int, Int>, DefaultWeightedEdge> = SimpleDirectedWeightedGraph(DefaultWeightedEdge::class.java)
    val tile = ArrayList<List<Int>>()
    var row = 0
    while (sc.hasNext()) {
        val line = sc.nextLine().split("").mapNotNull { it.toIntOrNull() }
        tile.add(line)
    }

    var rows = tile.size
    val cols = tile[0].size
    val arr = Array(rows*5) { IntArray(cols*5) }

    repeat(5) { yTile ->
        repeat(5) { xTile ->
            repeat(tile.size) { y ->
                repeat(tile[0].size) { x ->
                    arr[yTile*rows + y][xTile*cols + x] = (tile[y][x] + yTile + xTile - 1) % 9 + 1
                    g.addVertex(yTile*rows + y to xTile*cols + x)

                }
            }
        }
    }

    repeat(arr.size) { y ->
        repeat(arr[y].size) { x ->
            if (y > 0) { //top
                val weight = arr[y-1][x].toDouble()
                val edge = g.addEdge(y to x, y-1 to x)
                g.setEdgeWeight(edge, weight)
            }
            if (y < arr.size-1) { //bottom
                val weight = arr[y+1][x].toDouble()
                val edge = g.addEdge(y to x, y+1 to x)
                g.setEdgeWeight(edge, weight)
            }

            if (x > 0) { //left
                val weight = arr[y][x-1].toDouble()
                val edge = g.addEdge(y to x, y to x-1)
                g.setEdgeWeight(edge, weight)
            }
            if (x < arr[y].size-1) { //right
                val weight = arr[y][x+1].toDouble()
                val edge = g.addEdge(y to x, y to x+1)
                g.setEdgeWeight(edge, weight)
            }
        }
    }

    val dijkstra = DijkstraShortestPath(g)
    val path = dijkstra.getPath(0 to 0, arr.size - 1 to arr[0].size - 1)
    println(path.weight)


}