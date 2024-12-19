package y2021
import java.io.File
import java.util.*
import kotlin.collections.HashMap
import kotlin.math.abs
import kotlin.math.max
import kotlin.math.min

private const val considerDiagonals = true

fun main(args: Array<String>) {
    val sc = Scanner(File("inputs/day5"))
    val visited = HashMap<Point, Int>()

    var multiples = 0;

    while (sc.hasNextLine()) {
        val line = sc.nextLine();
        val points = line.split(" -> ").map { Point(it) }
        points[0].sequenceToPoint(points[1]).forEach {
            val current = visited.getOrPut(it) {0}
            visited[it] = current+1
            if (current == 1) {
                multiples++
            }
        }
    }

    println(multiples)

}

data class Point(val x: Int, val y: Int) {
    companion object {
        operator fun invoke(str: String): Point {
            val map = str.split(",").map { it.toInt() }
            return Point(map[0], map[1])
        }
    }

    fun sequenceToPoint(p: Point): Sequence<Point> {
        if (this.x == p.x) {
            val from = min(this.y, p.y)
            val to = max(this.y, p.y)
            return sequence {
                for (y in from..to) {
                    yield(Point(x, y))
                }
            }
        } else if (this.y == p.y) {
            val from = min(this.x, p.x)
            val to = max(this.x, p.x)
            return sequence {
                for (x in from..to) {
                    yield(Point(x, y))
                }
            }
        }
        else if (considerDiagonals and (abs(this.x - p.x) == abs(this.y - p.y))) { //diagonals
            val topPoint =  if (this.y < p.y) this else p;
            val bottomPoint =  if (this.y > p.y) this else p;
            val slope = if (topPoint.x < bottomPoint.x) 1 to 1 else -1 to 1
            var currentX = topPoint.x
            var currentY = topPoint.y
            return sequence {
                while (currentY <= bottomPoint.y) {
                    yield(Point(currentX, currentY))
                    currentX += slope.first
                    currentY += slope.second
                }
            }
        } else return emptySequence()
    }
}