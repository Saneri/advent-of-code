package day2

import readInput

fun isSorted(report: List<Int>, comparator: (Int, Int) -> Boolean): Boolean {
    for (i in 1 until report.size) {
        if (!comparator(report[i - 1], report[i])) {
            return false
        }
    }
    return true
}

fun isSafe(report: List<Int>): Boolean {
    val isDescending = isSorted(report) { a, b -> a - b in 1..3 }
    val isAscending = isSorted(report) { a, b -> b - a in 1..3 }
    return isAscending || isDescending
}

fun isSafeWithOneRemoval(report: List<Int>): Boolean {
    if (isSafe(report)) {
        return true
    }
    for (i in report.indices) {
        val modifiedReport = report.toMutableList().apply { removeAt(i) }
        if (isSafe(modifiedReport)) {
            return true
        }
    }
    return false
}

fun main() {
    val input = readInput(2)
    val count = input.map {
        val report = it.split(' ').map { level -> level.toInt() }
        isSafeWithOneRemoval(report)
    }.count { it }
    println(count)
}