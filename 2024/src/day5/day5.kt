package day5

import readInputText

fun isValid(numbers: List<String>, validPairs: Set<String>): Boolean {
    return numbers.zipWithNext().all { (a, b) ->
        val pair = "$a|$b"
        pair in validPairs
    }
}

fun sort(numbers: List<String>, validPairs: Set<String>): List<String> {
    val result = numbers.toMutableList()
    var iterations = 0
    val maxIterations = 1000

    while (!isValid(result, validPairs) && iterations < maxIterations) {
        for (i in 0 until result.size - 1) {
            val pair = "${result[i]}|${result[i + 1]}"
            if (pair !in validPairs) {
                result[i] = result[i + 1].also { result[i + 1] = result[i] }
            }
        }
        iterations++
    }

    if (!isValid(result, validPairs)) {
        throw Error("Valid combination not found after $maxIterations iterations")
    }
    return result
}

fun main() {
    val (rules, groups) = readInputText(5).split("\n\n").map { it.split("\n") }
    val pairs = rules.toSet()

    val sum = groups.sumOf { group ->
        val split = group.split(',')
        val isValid = isValid(split, pairs)
        if (isValid) {
            0
        } else {
            sort(split, pairs)[split.size / 2].toInt()
        }
    }
    println(sum)
}