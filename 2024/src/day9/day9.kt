package day9

import readInputText

fun main() {
    val input = readInputText(9).map { it.toString().toLong() } // Parse to Long
    val memory: MutableList<Long?> = mutableListOf() // Store nullable Long values
    input.forEachIndexed { index, value ->
        if (index % 2 == 0) {
            repeat(value.toInt()) { memory.add((index / 2).toLong()) } // Convert index to Long
        } else {
            repeat(value.toInt()) { memory.add(null) }
        }
    }

    var firstspace = memory.indexOfFirst { it == null }
    var lastfile = memory.indexOfLast { it != null }
    while (firstspace < lastfile) {
        memory[firstspace] = memory[lastfile].also { memory[lastfile] = memory[firstspace] }
        firstspace = memory.indexOfFirst { it == null }
        lastfile = memory.indexOfLast { it != null }
    }
    var sum = 0L
    memory.filterNotNull().forEachIndexed { index, value ->
        sum += index * value
    }
    println(sum)
}
