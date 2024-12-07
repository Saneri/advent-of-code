package day6

import Grid
import readInput

enum class Direction {
    UP, RIGHT, DOWN, LEFT
}

fun turnRight(direction: Direction): Direction = when (direction) {
    Direction.UP -> Direction.RIGHT
    Direction.RIGHT -> Direction.DOWN
    Direction.DOWN -> Direction.LEFT
    Direction.LEFT -> Direction.UP
}

fun moveForward(grid: Grid<Char>, direction: Direction): Boolean =
    when (direction) {
        Direction.UP -> grid.up()
        Direction.RIGHT -> grid.right()
        Direction.DOWN -> grid.down()
        Direction.LEFT -> grid.left()
    }

fun checkForward(grid: Grid<Char>, direction: Direction): Char? = when (direction) {
    Direction.UP -> grid.checkUp()
    Direction.RIGHT -> grid.checkRight()
    Direction.DOWN -> grid.checkDown()
    Direction.LEFT -> grid.checkLeft()
}

data class Visit(
    val position: Pair<Int, Int>,
    val direction: Direction
)

fun stepsToGetOut(
    startPosition: Pair<Int, Int>,
    input: Array<Array<Char>>,
    obstacle: Pair<Int, Int>
): Boolean {
    val grid = Grid(input, startPosition)
    grid.grid[obstacle.first][obstacle.second] = '#'
    var direction = Direction.UP
    var out = false
    var stuck = false
    val visited = mutableSetOf<Visit>()
    while (!out && !stuck) {
        while (checkForward(grid, direction) != '#') {
            val currentVisit = Visit(grid.currentCell, direction)
            if (visited.contains(currentVisit)) stuck = true
            visited.add(currentVisit)
            if (!moveForward(grid, direction)) {
                out = true
                break
            }
        }
        direction = turnRight(direction)
    }
    grid.grid[obstacle.first][obstacle.second] = '.'
    return stuck
}

fun main() {
    val input = readInput(6).map {
        it.toCharArray().toTypedArray()
    }.toTypedArray()

    val start = input.find { it.indexOf('^') != -1 }?.let {
        Pair(input.indexOf(it), it.indexOf('^'))
    }
    if (start == null) throw Exception("start not found")

    val grid = Grid(input, start)
    var direction = Direction.UP
    var out = false
    val uniquePositions = mutableSetOf<Pair<Int, Int>>()
    while (!out) {
        while (checkForward(grid, direction) != '#') {
            uniquePositions.add(grid.currentCell)
            if (!moveForward(grid, direction)) {
                out = true
                break
            }
        }
        direction = turnRight(direction)
    }
    val count = uniquePositions.count { position ->
        stepsToGetOut(start, input, position)
    }
    println(count)
}