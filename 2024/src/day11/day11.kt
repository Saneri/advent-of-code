package day11

import readInput

fun main() {
    val input = readInput(11)[0].split(" ").map(String::toLong).toMutableList()
    var totalSum = 0L
    val cache = mutableMapOf<Pair<Long, Int>, Long>()

    fun blink(stone: Long, depth: Int): Long {
        if (depth == 0) return 1
        val cachedValue = cache[Pair(stone, depth)]
        if (cachedValue != null) {
            return cachedValue
        }

        var count = 0L
        when {
            stone == 0L -> count += blink(1, depth - 1)
            stone.toString().length % 2 == 0 -> {
                count += blink(
                    stone.toString().substring(0, stone.toString().length / 2).toLong(), depth - 1
                )
                count += blink(
                    stone.toString().substring(stone.toString().length / 2).toLong(), depth - 1
                )
            }
//65601038650482
            else -> count += blink(stone * 2024, depth - 1)
        }
        cache[Pair(stone, depth)] = count
        return count
    }

    for (startStone in input) {
        val finalCount = blink(startStone, 75)
        totalSum += finalCount
    }

    println(totalSum)
}
