import kotlin.io.path.Path
import kotlin.io.path.readText

fun readInput(number: Int): List<String> =
    Path("src/day$number/input.txt").readText().trim().lines()

fun readInputText(number: Int): String =
    Path("src/day$number/input.txt").readText()
