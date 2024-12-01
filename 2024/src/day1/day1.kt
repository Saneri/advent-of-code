package day1

import readInput
import kotlin.math.abs

fun main() {
    val input = readInput(1)
    val pairs = input.map { it.split("\\s+".toRegex()) }

    val leftSides = pairs.map { it[0].trim().toInt() }.sorted()
    val rightSides = pairs.map { it[1].trim().toInt() }.sorted()

    val distance =
        leftSides.zip(rightSides).sumOf { (left, right) -> abs(left - right) }
    println("a: $distance")

    val rightSideCounts = rightSides.groupingBy { it }.eachCount()
    val similarity = leftSides.sumOf { (rightSideCounts[it] ?: 0) * it }
    println("b: $similarity")
}