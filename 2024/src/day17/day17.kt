package day17

import kotlin.math.pow

fun main() {
    val program = listOf(0, 3, 5, 4, 3, 0)

    fun solve(aRegister: Long): List<Int> {
        var a = aRegister
        var b = 0L
        var c = 0L

        fun combo(value: Int): Long {
            if (value < 4) {
                return value.toLong()
            }
            when (value) {
                4 -> return a
                5 -> return b
                6 -> return c
            }
            throw Exception("invalid operand $value")
        }

        var pointer = 0

        val out = mutableListOf<Char>()

        while (true) {
            if (pointer + 1 >= program.size) break
            val instruction = program[pointer]
            val operand = program[pointer + 1]
            var jump: Int? = null
            when (instruction) {
                0 -> a /= (2.0.pow(combo(operand).toInt())).toLong()
                1 -> b = b xor operand.toLong()
                2 -> b = combo(operand) % 8
                3 -> if (a != 0L) jump = operand
                4 -> b = b xor c
                5 -> out.addAll((combo(operand) % 8).toString().toCharArray().toList())
                6 -> b = a / (2.0.pow(combo(operand).toInt())).toLong()
                7 -> c = a / (2.0.pow(combo(operand).toInt())).toLong()
            }
            if (jump != null) {
                pointer = jump
            } else {
                pointer += 2
            }
        }

        return out.map { it.digitToInt() }
    }

    val base = 8.0

    fun findValidNumbers(position: Int, currentNumber: Long): List<Long> {
        if (position < 0) {
            val result = solve(currentNumber)
            return if (result == program) {
                listOf(currentNumber)
            } else {
                emptyList()
            }
        }

        return (0..8).flatMap { multiplier ->
            val newNumber = currentNumber + (multiplier * base.pow(position)).toLong()
            val result = solve(newNumber)

            if (result.size > position && result[position] == program[position]) {
                findValidNumbers(position - 1, newNumber)
            } else {
                emptyList()
            }
        }
    }

    val validNumbers = findValidNumbers(program.size - 1, 0)
    println(validNumbers.min())
}
