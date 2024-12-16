package day14

import readInput

data class Robot(var position: Pair<Int, Int>, val velocity: Pair<Int, Int>)

const val HEIGHT = 103
const val WIDTH = 101

fun printRobotGrid(robots: List<Robot>, width: Int, height: Int) {
    val outputBuffer = StringBuilder()
    val grid = Array(height) { IntArray(width) }
    for (robot in robots) {
        val (x, y) = robot.position
        grid[y][x]++
    }


    for (y in 0 until height) {
        for (x in 0 until width) {
            if (grid[y][x] > 0) {
                outputBuffer.append(grid[y][x])
            } else {
                outputBuffer.append(".")
            }
        }
        outputBuffer.append("\n")
    }
    print(outputBuffer.toString())
}

fun circularDifference(pos: Int, vel: Int, range: Int): Int {
    val newPos = pos + vel
    if (newPos < 0) {
        return (newPos % range + range) % range
    }
    return newPos % range
}

fun main() {
    val input = readInput(14)
    val values = input.map { "-?\\d+".toRegex().findAll(it).map { i -> i.value.toInt() }.toList() }
    val robots = values.map { Robot(it[0] to it[1], it[2] to it[3]) }

    repeat(10000) { index ->
        for (robot in robots) {
            robot.position = Pair(
                circularDifference(robot.position.first, robot.velocity.first, WIDTH),
                circularDifference(robot.position.second, robot.velocity.second, HEIGHT)
            )
        }
        // first checked every frame, then saw a pattern which repeated every 101 frames
        if (index % 101 == 2236 % 101 - 1 && index > 7000) {
            printRobotGrid(robots, WIDTH, HEIGHT)
            println("Index: ${index + 1}")
            readlnOrNull()
        }
    }

}