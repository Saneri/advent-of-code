package day13

import readInputText
import kotlin.math.roundToLong

fun solveCramersRule(
    a11: Double, a12: Double, b1: Double,
    a21: Double, a22: Double, b2: Double
): Pair<Double?, Double?> {
    val detA = a11 * a22 - a12 * a21
    if (detA == 0.0) {
        return null to null
    }
    val detX = b1 * a22 - b2 * a12
    val detY = a11 * b2 - a21 * b1
    val x = detX / detA
    val y = detY / detA

    return x to y
}


fun isInteger(value: Double): Boolean {
    val rounded = value.roundToLong()
    return kotlin.math.abs(value - rounded) < 1e-9
}

fun main() {
    val input = readInputText(13).split("\n\n").map { it.split("\n") }
    var spent = 0L
    for (line in input) {
        val values =
            "\\d+".toRegex().findAll(line.toString()).map { i -> i.value.toDouble() }.toList()

        val add = 10000000000000
        val (x, y) =
            solveCramersRule(
                values[0],
                values[2],
                values[4] + add,
                values[1],
                values[3],
                values[5] + add
            )
        if (x != null && y != null) {
            if (isInteger(x) && isInteger(y)) {
                spent += x.roundToLong() * 3 + y.roundToLong()
            }
        } else {
            throw Exception("Solution not found")
        }
    }
    println(spent)
}