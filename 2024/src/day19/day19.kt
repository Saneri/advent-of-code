package day19

import readInputText

fun main() {
    val input = readInputText(19).split("\n\n")
    val patterns = input[0].split(", ")
    val designs = input[1].split("\n")

    val cache = mutableMapOf<String, Long>()

    fun countWays(design: String): Long {
        if (design.isEmpty()) return 1
        if (cache.containsKey(design)) return cache[design]!!

        var count = 0L
        for (pattern in patterns) {
            if (design.startsWith(pattern)) {
                val rest = design.removePrefix(pattern)
                count += countWays(rest)
            }
        }
        cache[design] = count
        return count
    }

    val sum = designs.sumOf { countWays(it) }
    println(sum)
}