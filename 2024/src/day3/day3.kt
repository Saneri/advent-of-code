package day3

import readInputText

const val MUL_REGEX = "mul\\((\\d{1,3},\\d{1,3})\\)"
const val DO_REGEX = "do\\(\\)"
const val DONT_REGEX = "don't\\(\\)"

fun main() {
    val input = readInputText(3)
    val regex = Regex("$MUL_REGEX|$DO_REGEX|$DONT_REGEX")
    var enabled = true
    val result = regex.findAll(input).sumOf { match ->
        when {
            match.value.matches(Regex(DONT_REGEX)) -> {
                enabled = false
                0
            }

            match.value.matches(Regex(DO_REGEX)) -> {
                enabled = true
                0
            }

            match.value.matches(Regex(MUL_REGEX)) && enabled -> {
                val (group) = match.destructured
                val values = group.split(",").map { value -> value.toInt() }
                values[0] * values[1]
            }

            else -> 0
        }


    }
    println(result)
}