package day20

import readInput
import toCoordinates
import kotlin.math.abs

enum class Direction {
    UP, RIGHT, DOWN, LEFT
}

data class Node(
    val position: Pair<Int, Int>,
    val picoseconds: Int
)

fun move(current: Pair<Int, Int>, direction: Direction): Pair<Int, Int> {
    return when (direction) {
        Direction.UP -> current.first - 1 to current.second
        Direction.RIGHT -> current.first to current.second + 1
        Direction.DOWN -> current.first + 1 to current.second
        Direction.LEFT -> current.first to current.second - 1
    }
}

const val MAX_DISTANCE = 20
const val MIN_SAVED = 50

fun main() {
    val grid = toCoordinates(readInput(20)).toMutableMap()

    val start = grid.entries.first { it.value == 'S' }.key
    val shortcuts = mutableListOf<Node>()
    var sum = 0

    val visited = mutableSetOf<Pair<Int, Int>>()
    val queue = ArrayDeque<Node>()
    queue.add(Node(start, 0))

    fun checkShortcuts(start: Pair<Int, Int>, picoseconds: Int) {
        for (dx in -MAX_DISTANCE..MAX_DISTANCE) {
            val maxDy = MAX_DISTANCE - abs(dx)
            for (dy in -maxDy..maxDy) {
                val next = start.first + dx to start.second + dy
                val distance = abs(dx) + abs(dy)
                if (distance <= MAX_DISTANCE && (grid[next] == '.' || grid[next] == 'E')
                    && !visited.contains(next)
                ) {
                    shortcuts.add(Node(next, picoseconds + distance))
                }
            }
        }
    }

    while (queue.isNotEmpty()) {
        val current = queue.removeFirst()
        shortcuts.filter { it.position == current.position }.forEach { shortcut ->
            val saved = current.picoseconds - shortcut.picoseconds
            if (saved >= MIN_SAVED) {
                sum++
            }
        }

        visited.add(current.position)

        checkShortcuts(current.position, current.picoseconds)
        Direction.entries.forEach {
            val next = move(current.position, it)
            if ((grid[next] == '.' || grid[next] == 'E') && !visited.contains(next)) {
                queue.add(Node(next, current.picoseconds + 1))
            }
        }
    }
    println(sum)
}