import kotlin.io.path.Path
import kotlin.io.path.readText

fun readInput(number: Int): List<String> =
    Path("src/day$number/input.txt").readText().trim().lines()

fun readInputText(number: Int): String =
    Path("src/day$number/input.txt").readText()

fun toCoordinates(input: List<String>): Map<Pair<Int, Int>, Char> =
    input.flatMapIndexed { rowIndex, row ->
        row.mapIndexed { colIndex, char ->
            Pair(rowIndex, colIndex) to char
        }
    }.toMap()
