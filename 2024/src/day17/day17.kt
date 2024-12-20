package day17

import readInputText
import kotlin.math.pow

fun main() {
    val input = readInputText(17).split("\n\n")
    val registers = input[0].split("\n").associate { line ->
        val (key, value) = line.split(": ")
        key.split(" ")[1] to value.toInt()
    }.toMutableMap()

    var a = registers["A"] ?: 0
    var b = registers["B"] ?: 0
    var c = registers["C"] ?: 0

    fun combo(value: Int): Int {
        if (value < 4) {
            return value
        }
        when (value) {
            4 -> return a
            5 -> return b
            6 -> return c
        }
        throw Exception("invalid operand $value")
    }

    val instructions = input[1].split(": ")[1].split(",").map { it.toInt() }
    var pointer = 0

    val out = mutableListOf<Char>()

    while (true) {
        if (pointer + 1 >= instructions.size) break
        val instruction = instructions[pointer]
        val operand = instructions[pointer + 1]
        var jump: Int? = null
        when (instruction) {
            0 -> a /= 2.0.pow(combo(operand)).toInt()
            1 -> b = b xor operand
            2 -> b = combo(operand) % 8
            3 -> if (a != 0) jump = operand
            4 -> b = b xor c
            5 -> out.addAll((combo(operand) % 8).toString().toCharArray().toList())
            6 -> b = a / 2.0.pow(combo(operand)).toInt()
            7 -> c = a / 2.0.pow(combo(operand)).toInt()
        }
        if (jump != null) {
            pointer = jump
        } else {
            pointer += 2
        }
    }

    println(out.joinToString(","))
}