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
    val tileCount: Int
) {
    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (javaClass != other?.javaClass) return false
        other as Node
        return position == other.position && direction == other.direction
    }

    override fun hashCode(): Int {
        var result = position.hashCode()
        result = 31 * result + direction.hashCode()
        return result
    }
}

fun Node.move(action: Action): Node = when (action) {
    Action.STRAIGHT -> {
        val newPos = when (this.direction) {
            Direction.UP -> position.first - 1 to position.second
            Direction.RIGHT -> position.first to position.second + 1
            Direction.DOWN -> position.first + 1 to position.second
            Direction.LEFT -> position.first to position.second - 1

        }
        Node(newPos, this.direction, this.weight + 1, this.tileCount + 1)
    }

    Action.RIGHT -> {
        val newDir = when (direction) {
            Direction.UP -> Direction.RIGHT
            Direction.RIGHT -> Direction.DOWN
            Direction.DOWN -> Direction.LEFT
            Direction.LEFT -> Direction.UP
        }
        Node(this.position, newDir, this.weight + 1000, this.tileCount)
    }

    Action.LEFT -> {
        val newDir = when (direction) {
            Direction.UP -> Direction.LEFT
            Direction.RIGHT -> Direction.UP
            Direction.DOWN -> Direction.RIGHT
            Direction.LEFT -> Direction.DOWN
        }
        Node(this.position, newDir, this.weight + 1000, this.tileCount)
    }
}

fun main() {
    val grid = toCoordinates(readInput(16))
    val start = grid.entries.first { it.value == 'S' }.key
    val end = grid.entries.first { it.value == 'E' }.key
    val visited = mutableSetOf<Node>()
    val queue = PriorityQueue<Node>(compareBy { it.weight })
    queue.add(Node(start, Direction.RIGHT, 0, 1))

    while (queue.isNotEmpty()) {
        val current = queue.poll()
        if (current.position == end) {
            println(current.weight)
            break
        }
        visited.add(current)

        Action.entries.forEach { action ->
            val nextNode = current.move(action)
            if (!visited.contains(nextNode) && grid[nextNode.position] != '#') {
                queue.add(nextNode)
            }
        }
    }
}