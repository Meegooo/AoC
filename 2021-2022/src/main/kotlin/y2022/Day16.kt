package y2022

import kotlin.math.max

fun main() {
    Day16.part1()
}

object Day16 {
    lateinit var usefulNodes: List<Node>
    fun part1() {
        val nodes = javaClass.getResource(javaClass.simpleName)!!.readText().lines().map { line ->
            val match =
                """Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? (\w\w(?:, \w\w)*)""".toRegex().find(line)
            val name = match!!.groupValues[1]
            val rate = match.groupValues[2].toInt()
            val neighbours = match.groupValues[3].split(", ").toSet()
            Node(name, rate, neighbours)
        }.associateBy { it.name }

        val (distances, nameToIdx) = getPairwiseDistance(nodes)

        for (a in distances) {
            for (index in a.indices) {
                a[index] += 1
            }
        }

        usefulNodes = nodes.filterValues { it.flow > 0 || it.name == "AA" }.values.toList()

        usefulNodes.forEachIndexed { idx, node -> println("$idx. ${node.name}") }
        val usefulFlows = usefulNodes.mapIndexed { index, node -> index to node.flow }.toMap()
        val indexOfAA = usefulNodes.indexOfFirst { it.name == "AA" }
        val usefulDistances = Array(usefulNodes.size) { IntArray(usefulNodes.size) }

        usefulNodes.forEachIndexed { leftIndex, leftNode ->

            usefulNodes.forEachIndexed { rightIndex, rightNode ->
                usefulDistances[leftIndex][rightIndex] =
                    distances[nameToIdx[leftNode.name]!!][nameToIdx[rightNode.name]!!]
            }
        }
        dfsWithAnElephant(usefulDistances, usefulFlows, indexOfAA, indexOfAA, 0, 0, setOf(0).toMutableSet(), 26, 0, 0)
        println(maxTotalFlow)
    }

    var maxTotalFlow = 0

    fun dfs(
        graph: Array<IntArray>,
        usefulFlows: Map<Int, Int>,
        current: Int,
        visited: MutableSet<Int>,
        timeRemaining: Int,
        flowRate: Int,
        totalFlow: Int
    ) {
        maxTotalFlow = max(maxTotalFlow, totalFlow + flowRate * timeRemaining)
        val neighbours = graph[current]
        neighbours.forEachIndexed { dest, weight ->
            if (timeRemaining == 30) {
                println(dest)
            }
            if (dest !in visited && dest != current && timeRemaining >= weight) {
                visited.add(dest)
                dfs(
                    graph,
                    usefulFlows,
                    dest,
                    visited,
                    timeRemaining - weight,
                    flowRate + usefulFlows[dest]!!,
                    totalFlow + flowRate * weight
                )
                visited.remove(dest)
            }
        }
    }

    fun dfsWithAnElephant(
        graph: Array<IntArray>,
        usefulFlows: Map<Int, Int>,
        currentHuman: Int,
        currentElephant: Int,
        cooldownHuman: Int,
        cooldownElephant: Int,
        visited: MutableSet<Int>,
        timeRemaining: Int,
        flowRate: Int,
        totalFlow: Int
    ) {
        val timePassed = 26 - timeRemaining + 1
//        //DEBUG
//        if (timePassed == 1 && currentHuman == 6 && currentElephant == 3 && cooldownElephant == 2 && cooldownHuman == 3) {
//            println("$timePassed. Human -> ${usefulNodes[currentHuman].name} ($cooldownHuman). Elephant -> ${usefulNodes[currentElephant].name} ($cooldownElephant). Flow rate $flowRate, total $totalFlow")
//        }
//        if (timePassed == 2 && currentHuman == 6 && currentElephant == 3 && cooldownElephant == 1 && cooldownHuman == 2) {
//            println("$timePassed. Human -> ${usefulNodes[currentHuman].name} ($cooldownHuman). Elephant -> ${usefulNodes[currentElephant].name} ($cooldownElephant). Flow rate $flowRate, total $totalFlow")
//        }
//        if (timePassed == 3 && currentHuman == 6 && currentElephant == 3 && cooldownElephant == 0 && cooldownHuman == 1) {
//            println("$timePassed. Human -> ${usefulNodes[currentHuman].name} ($cooldownHuman). Elephant -> ${usefulNodes[currentElephant].name} ($cooldownElephant). Flow rate $flowRate, total $totalFlow")
//        }
//        if (timePassed == 3 && currentHuman == 6 && currentElephant == 5 && cooldownElephant == 5 && cooldownHuman == 1) {
//            println("$timePassed. Human -> ${usefulNodes[currentHuman].name} ($cooldownHuman). Elephant -> ${usefulNodes[currentElephant].name} ($cooldownElephant). Flow rate $flowRate, total $totalFlow")
//        }
//        if (timePassed == 4 && currentHuman == 6 && currentElephant == 5 && cooldownElephant == 4 && cooldownHuman == 0 && flowRate == 20) {
//            println("$timePassed. Human -> ${usefulNodes[currentHuman].name} ($cooldownHuman). Elephant -> ${usefulNodes[currentElephant].name} ($cooldownElephant). Flow rate $flowRate, total $totalFlow")
//        }
//        if (timePassed == 4 && currentHuman == 1 && currentElephant == 5 && cooldownElephant == 4 && cooldownHuman == 4 && flowRate == 41) {
//            println("$timePassed. Human -> ${usefulNodes[currentHuman].name} ($cooldownHuman). Elephant -> ${usefulNodes[currentElephant].name} ($cooldownElephant). Flow rate $flowRate, total $totalFlow")
//        }
//        if (timePassed == 5 && currentHuman == 1 && currentElephant == 5 && cooldownElephant == 3 && cooldownHuman == 3 && flowRate == 41) {
//            println("$timePassed. Human -> ${usefulNodes[currentHuman].name} ($cooldownHuman). Elephant -> ${usefulNodes[currentElephant].name} ($cooldownElephant). Flow rate $flowRate, total $totalFlow")
//        }
//        if (timePassed == 6 && currentHuman == 1 && currentElephant == 5 && cooldownElephant == 2 && cooldownHuman == 2 && flowRate == 41) {
//            println("$timePassed. Human -> ${usefulNodes[currentHuman].name} ($cooldownHuman). Elephant -> ${usefulNodes[currentElephant].name} ($cooldownElephant). Flow rate $flowRate, total $totalFlow")
//        }
//        if (timePassed == 7 && currentHuman == 1 && currentElephant == 5 && cooldownElephant == 1 && cooldownHuman == 1 && flowRate == 41) {
//            println("$timePassed. Human -> ${usefulNodes[currentHuman].name} ($cooldownHuman). Elephant -> ${usefulNodes[currentElephant].name} ($cooldownElephant). Flow rate $flowRate, total $totalFlow")
//        }
//        if (timePassed == 8 && currentHuman == 1 && currentElephant == 5 && cooldownElephant == 0 && cooldownHuman == 0 && flowRate == 41) {
//            println("$timePassed. Human -> ${usefulNodes[currentHuman].name} ($cooldownHuman). Elephant -> ${usefulNodes[currentElephant].name} ($cooldownElephant). Flow rate $flowRate, total $totalFlow")
//        }
//        if (timePassed == 8 && currentHuman == 2 && currentElephant == 4 && cooldownElephant == 4 && cooldownHuman == 2 && flowRate == 76) {
//            println("$timePassed. Human -> ${usefulNodes[currentHuman].name} ($cooldownHuman). Elephant -> ${usefulNodes[currentElephant].name} ($cooldownElephant). Flow rate $flowRate, total $totalFlow")
//        }
//        if (timePassed == 9 && currentHuman == 2 && currentElephant == 4 && cooldownElephant == 3 && cooldownHuman == 1 && flowRate == 76) {
//            println("$timePassed. Human -> ${usefulNodes[currentHuman].name} ($cooldownHuman). Elephant -> ${usefulNodes[currentElephant].name} ($cooldownElephant). Flow rate $flowRate, total $totalFlow")
//        }
//        if (timePassed == 10 && currentHuman == 2 && currentElephant == 4 && cooldownElephant == 2 && cooldownHuman == 0 && flowRate == 76) {
//            println("$timePassed. Human -> ${usefulNodes[currentHuman].name} ($cooldownHuman). Elephant -> ${usefulNodes[currentElephant].name} ($cooldownElephant). Flow rate $flowRate, total $totalFlow")
//        }
//        if (timePassed == 10 && currentHuman == 2 && currentElephant == 4 && cooldownElephant == 2 && cooldownHuman == 1000 && flowRate == 78) {
//            println("$timePassed. Human -> ${usefulNodes[currentHuman].name} ($cooldownHuman). Elephant -> ${usefulNodes[currentElephant].name} ($cooldownElephant). Flow rate $flowRate, total $totalFlow")
//        }
//        if (timePassed == 11 && currentHuman == 2 && currentElephant == 4 && cooldownElephant == 1 && cooldownHuman == 999  && flowRate == 78) {
//            println("$timePassed. Human -> ${usefulNodes[currentHuman].name} ($cooldownHuman). Elephant -> ${usefulNodes[currentElephant].name} ($cooldownElephant). Flow rate $flowRate, total $totalFlow")
//        }
//        if (timePassed == 12 && currentHuman == 2 && currentElephant == 4 && cooldownElephant == 0 && cooldownHuman == 998  && flowRate == 78) {
//            println("$timePassed. Human -> ${usefulNodes[currentHuman].name} ($cooldownHuman). Elephant -> ${usefulNodes[currentElephant].name} ($cooldownElephant). Flow rate $flowRate, total $totalFlow")
//        }
//        if (timePassed == 12 && currentHuman == 2 && currentElephant == 4 && cooldownElephant >0 && cooldownHuman == 998  && flowRate == 81) {
//            println("$timePassed. Human -> ${usefulNodes[currentHuman].name} ($cooldownHuman). Elephant -> ${usefulNodes[currentElephant].name} ($cooldownElephant). Flow rate $flowRate, total $totalFlow")
//        }


        if (timeRemaining == 0) {
            maxTotalFlow = max(maxTotalFlow, totalFlow)
            return
        }
        if (cooldownHuman > 0 && cooldownElephant > 0) {
            dfsWithAnElephant(
                graph,
                usefulFlows,
                currentHuman,
                currentElephant,
                cooldownHuman - 1,
                cooldownElephant - 1,
                visited,
                timeRemaining - 1,
                flowRate,
                totalFlow + flowRate
            )
            return
        }

        //human
        if (cooldownHuman == 0) {
            val neighbours = graph[currentHuman]
            var visitedInside = false
            neighbours.forEachIndexed { dest, weight ->
                if (dest !in visited && dest != currentHuman && timeRemaining >= weight && weight != Int.MIN_VALUE) {
                    visited.add(dest)
                    visitedInside = true
                    dfsWithAnElephant(
                        graph,
                        usefulFlows,
                        dest,
                        currentElephant,
                        weight,
                        cooldownElephant,
                        visited,
                        timeRemaining,
                        flowRate + usefulFlows[currentHuman]!!,
                        totalFlow
                    )
                    visited.remove(dest)
                }
            }
            if (!visitedInside) {
                dfsWithAnElephant(
                    graph,
                    usefulFlows,
                    currentHuman,
                    currentElephant,
                    1000,
                    cooldownElephant,
                    visited,
                    timeRemaining,
                    flowRate + usefulFlows[currentHuman]!!,
                    totalFlow
                )
            }

        }

        if (cooldownElephant == 0) {
            val neighbours = graph[currentElephant]
            var visitedInside = false
            neighbours.forEachIndexed { dest, weight ->
                if (dest !in visited && dest != currentElephant && timeRemaining >= weight && weight != Int.MIN_VALUE) {
                    visited.add(dest)
                    visitedInside = true
                    dfsWithAnElephant(
                        graph,
                        usefulFlows,
                        currentHuman,
                        dest,
                        cooldownHuman,
                        weight,
                        visited,
                        timeRemaining,
                        flowRate + usefulFlows[currentElephant]!!,
                        totalFlow
                    )
                    visited.remove(dest)
                }
            }
            if (!visitedInside) {
                dfsWithAnElephant(
                    graph,
                    usefulFlows,
                    currentHuman,
                    currentElephant,
                    cooldownHuman,
                    1000,
                    visited,
                    timeRemaining,
                    flowRate + usefulFlows[currentElephant]!!,
                    totalFlow
                )
            }
        }
    }


    fun getPairwiseDistance(nodes: Map<String, Node>): Pair<Array<IntArray>, Map<String, Int>> {
        val idxToName = nodes.keys.toList()
        val nameToIdx = idxToName.mapIndexed { index, s -> s to index }.toMap()

        val distances = Array(idxToName.size) { IntArray(idxToName.size) { Int.MAX_VALUE } }

        for ((_, node) in nodes) {
            for (neighbour in node.neighbours) {
                val idxLeft = nameToIdx[node.name]!!
                val idxRight = nameToIdx[nodes[neighbour]!!.name]!!
                distances[idxLeft][idxRight] = 1
            }
        }

        for (i in idxToName.indices) {
            distances[i][i] = 0
        }

        for (k in idxToName.indices) {
            for (i in idxToName.indices) {
                for (j in idxToName.indices) {
                    if (distances[i][k] != Int.MAX_VALUE && distances[k][j] != Int.MAX_VALUE &&
                        distances[i][j] > distances[i][k] + distances[k][j]
                    ) {
                        distances[i][j] = distances[i][k] + distances[k][j]
                    }
                }
            }
        }
        return (distances to nameToIdx)
    }

    class Node(val name: String, val flow: Int, val neighbours: Set<String>)

}

//2283 too low