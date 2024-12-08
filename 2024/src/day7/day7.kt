package day7

import readInput

fun main() {
    val input = readInput(7)
    val sum = input.sumOf {
        val (start, rest) = it.split(": ")
        val testValue = start.toLong()
        val numbers = rest.split(' ').map(String::toLong)
        val operators = listOf('+', '*', '|')

        fun evaluate(ans: Long, index: Int, operator: Char): Boolean {
            var number = numbers[index]
            when (operator) {
                '+' -> number += ans
                '*' -> number *= ans
                '|' -> number = "$ans$number".toLong()
            }
            if (index < numbers.size - 1) {
                return operators.any { op -> evaluate(number, index + 1, op) }
            }
            return number == testValue
        }

        if (operators.any { op -> evaluate(0, 0, op) }) testValue else 0
    }
    println(sum)
}