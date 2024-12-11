package day9

import readInputText

// difficult b part, quite inefficient code and hard to read code that takes 10s+ to run but works
fun main() {
    val input = readInputText(9).map { it.toString().toLong() }
    val memory: MutableList<Long?> = mutableListOf()
    input.forEachIndexed { index, value ->
        if (index % 2 == 0) {
            repeat(value.toInt()) { memory.add((index / 2).toLong()) }
        } else {
            repeat(value.toInt()) { memory.add(null) }
        }
    }

    fun findFirstNullSection(memory: List<Long?>, startIndex: Int): Pair<Int, Int> {
        val firstSpace = memory.withIndex()
            .firstOrNull { it.index >= startIndex && it.value == null }?.index ?: -1
        if (firstSpace < 0) {
            throw Exception("Next null section not found")
        }
        var lastSpace = firstSpace
        while (lastSpace < memory.size && memory[lastSpace] == null) {
            lastSpace++
        }
        lastSpace--
        return Pair(firstSpace, lastSpace)
    }

    fun findLastNonNullSection(memory: List<Long?>, number: Long): Pair<Int, Int> {
        val lastFile = memory.indexOfLast { it == number }
        var firstFile = lastFile
        while (firstFile > 0 && memory[firstFile] == number) {
            firstFile--
        }
        firstFile++
        return Pair(firstFile, lastFile)
    }

    val number = memory.filterNotNull().maxOrNull() ?: -1
    if (number < 0) {
        throw Exception("values not found")
    }

    for (i in number downTo 0) {
        var (firstSpace, lastSpace) = findFirstNullSection(memory, 0)
        val (firstFile, lastFile) = findLastNonNullSection(memory, i)

        while (firstSpace != -1 && firstSpace < lastFile) {
            val nullSectionSize = lastSpace - firstSpace + 1
            val numberSectionSize = lastFile - firstFile + 1
            if (nullSectionSize >= numberSectionSize) {
                val sectionSize = minOf(nullSectionSize, numberSectionSize)
                for (j in 0 until sectionSize) {
                    memory[firstSpace + j] =
                        memory[firstFile + j].also {
                            memory[firstFile + j] = memory[firstSpace + j]
                        }
                }
                break
            } else {
                val new = findFirstNullSection(memory, lastSpace + 1)
                firstSpace = new.first
                lastSpace = new.second
            }
        }

    }

    var sum = 0L
    memory.forEachIndexed { index, value ->
        if (value != null) {
            sum += index * value
        }
    }
    println(sum)
}
