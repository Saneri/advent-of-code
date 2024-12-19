package day16

import readInput
import toCoordinates
import java.util.*

enum class Direction {
    UP, RIGHT, DOWN, LEFT
}

enum class Action {
    STRAIGHT, RIGHT, LEFT
}

data class Node(
    val position: Pair<Int, Int>,
    val direction: Direction,
    val weight: Int,
    val parent: Node? = null
)

fun Node.move(action: Action): Node = when (action) {
    Action.STRAIGHT -> {
        val newPos = when (this.direction) {
            Direction.UP -> position.first - 1 to position.second
            Direction.RIGHT -> position.first to position.second + 1
            Direction.DOWN -> position.first + 1 to position.second
            Direction.LEFT -> position.first to position.second - 1

        }
        Node(newPos, this.direction, this.weight + 1)
    }

    Action.RIGHT -> {
        val newDir = when (direction) {
            Direction.UP -> Direction.RIGHT
            Direction.RIGHT -> Direction.DOWN
            Direction.DOWN -> Direction.LEFT
            Direction.LEFT -> Direction.UP
        }
        Node(this.position, newDir, this.weight + 1000)
    }

    Action.LEFT -> {
        val newDir = when (direction) {
            Direction.UP -> Direction.LEFT
            Direction.RIGHT -> Direction.UP
            Direction.DOWN -> Direction.RIGHT
            Direction.LEFT -> Direction.DOWN
        }
        Node(this.position, newDir, this.weight + 1000)
    }
}

fun main() {
    val grid = toCoordinates(readInput(16))
    val start = grid.entries.first { it.value == 'S' }.key
    val end = grid.entries.first { it.value == 'E' }.key
    val visited = mutableMapOf<Pair<Pair<Int, Int>, Direction>, Int>()
    val queue = PriorityQueue<Node>(compareBy { it.weight })
    queue.add(Node(start, Direction.RIGHT, 0))

    val fastPathPositions = mutableSetOf<Pair<Int, Int>>()
    var minWeight = Int.MAX_VALUE

    while (queue.isNotEmpty()) {
        val current = queue.poll()
        if (current.position == end) {
            if (current.weight <= minWeight) {
                minWeight = current.weight
                var node = current
                while (node.parent != null) {
                    fastPathPositions.add(node.position)
                    node = node.parent
                }
            }
            continue
        }
        visited[current.position to current.direction] = current.weight

        Action.entries.forEach { action ->
            val nextNode = current.move(action).copy(parent = current)
            if (visited[nextNode.position to nextNode.direction] == null && grid[nextNode.position] != '#') {
                queue.add(nextNode)
            }
        }
    }
    println(fastPathPositions.size)
}