package day22

import readInput

fun generateSecret(number: Long): Long {
    val first = ((number * 64) xor number) % 16777216
    val second = (first.floorDiv(32) xor first) % 16777216
    return ((second * 2048) xor second) % 16777216
}

fun main() {
    val input = readInput(22).map(String::toLong)

    val lists = mutableListOf<List<Int>>()
    input.forEach { number ->
        val onesDigits = mutableListOf<Int>()
        var num = number

        repeat(2000) {
            onesDigits.add((num % 10).toInt())
            num = generateSecret(num)
        }

        lists.add(onesDigits)
    }
    val mappings = mutableListOf<Map<List<Int>, Int>>()

    lists.forEach { onesDigits ->
        val mapping = mutableMapOf<List<Int>, Int>()
        val differences = mutableListOf<Int>()
        for (i in 0 until onesDigits.size - 1) {
            differences.add(onesDigits[i + 1] - onesDigits[i])
        }

        for (i in 0 until differences.size - 4) {
            val key = differences.subList(i, i + 4).toList()
            val value = onesDigits[i + 4]
            if (!mapping.containsKey(key)) {
                mapping[key] = value
            }
        }

        mappings.add(mapping)
    }

    val mergedMapping = mutableMapOf<List<Int>, Int>()
    mappings.forEach { mapping ->
        mapping.forEach { (key, value) ->
            mergedMapping[key] = mergedMapping.getOrDefault(key, 0) + value
        }
    }

    val max = mergedMapping.maxByOrNull { it.value }
    if (max != null) {
        println(max.value)
    }
}