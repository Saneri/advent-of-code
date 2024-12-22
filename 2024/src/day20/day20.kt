package day20

import readInput
import toCoordinates

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

fun main() {
    val grid = toCoordinates(readInput(20)).toMutableMap()

    val start = grid.entries.first { it.value == 'S' }.key
    val shortcuts = mutableListOf<Node>()
    var sum = 0

    val visited = mutableSetOf<Pair<Int, Int>>()
    val queue = ArrayDeque<Node>()
    queue.add(Node(start, 0))

    while (queue.isNotEmpty()) {
        val current = queue.removeFirst()
        shortcuts.filter { it.position == current.position }.forEach { shortcut ->
            val saved = current.picoseconds - shortcut.picoseconds
            if (saved >= 100) {
                sum++
            }
        }

        visited.add(current.position)

        Direction.entries.forEach {
            val next = move(current.position, it)
            if ((grid[next] == '.' || grid[next] == 'E') && !visited.contains(next)) {
                queue.add(Node(next, current.picoseconds + 1))
            }
            if (grid[next] == '#') {
                Direction.entries.forEach { dir ->
                    val nextShortcut = move(next, dir)
                    if ((grid[nextShortcut] == '.' || grid[nextShortcut] == 'E')
                        && !visited.contains(nextShortcut)
                    ) {
                        shortcuts.add(Node(nextShortcut, current.picoseconds + 2))
                    }
                }
            }
        }
    }
    println(sum)
}