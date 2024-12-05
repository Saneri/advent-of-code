package day4

import Grid
import readInput

const val KEYWORD = "XMAS"

fun main() {
    val input = readInput(4).map { line -> line.toCharArray().toTypedArray() }.toTypedArray()
    var countA = 0
    var countB = 0
    for (row in input.indices) {
        for (col in 0 until input[row].size) {
            val grid = Grid(input, Pair(row, col))
            if (checkKeyword(grid) { grid.up() }) countA++
            if (checkKeyword(grid) { grid.upRight() }) countA++
            if (checkKeyword(grid) { grid.right() }) countA++
            if (checkKeyword(grid) { grid.downRight() }) countA++
            if (checkKeyword(grid) { grid.down() }) countA++
            if (checkKeyword(grid) { grid.downLeft() }) countA++
            if (checkKeyword(grid) { grid.left() }) countA++
            if (checkKeyword(grid) { grid.upLeft() }) countA++

            if (checkXFormation(grid)) countB++
        }
    }
    println("a: $countA")
    println("b: $countB")
}

fun checkXFormation(grid: Grid<Char>): Boolean {
    if (grid.currentValue() != 'A') {
        return false
    }
    val chars: MutableList<Char> = mutableListOf()
    val position = grid.currentCell
    if (grid.upRight()) chars.add(grid.currentValue())
    grid.currentCell = position
    if (grid.downLeft()) chars.add(grid.currentValue())
    grid.currentCell = position
    var countM = chars.count { it == 'M' }
    var countS = chars.count { it == 'S' }
    if (countM != 1 || countS != 1) return false
    if (grid.downRight()) chars.add(grid.currentValue())
    grid.currentCell = position
    if (grid.upLeft()) chars.add(grid.currentValue())
    grid.currentCell = position
    countM = chars.count { it == 'M' }
    countS = chars.count { it == 'S' }
    return countM == 2 && countS == 2
}

fun checkKeyword(grid: Grid<Char>, move: () -> Boolean): Boolean {
    val position = grid.currentCell
    if (grid.currentValue() != KEYWORD[0]) {
        return false
    }
    for (i in 1 until KEYWORD.length) {
        if (!move() || grid.currentValue() != KEYWORD[i]) {
            grid.currentCell = position
            return false
        }
    }
    grid.currentCell = position
    return true
}