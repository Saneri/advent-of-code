class Grid<T>(private val grid: Array<Array<T>>, var currentCell: Pair<Int, Int>) {
    private var rows = grid.size
    private var cols = grid[0].size

    fun up(): Boolean {
        return if (currentCell.first > 0) {
            currentCell = currentCell.copy(first = currentCell.first - 1)
            true
        } else {
            false
        }
    }

    fun down(): Boolean {
        return if (currentCell.first < rows - 1) {
            currentCell = currentCell.copy(first = currentCell.first + 1)
            true
        } else {
            false
        }
    }

    fun left(): Boolean {
        return if (currentCell.second > 0) {
            currentCell = currentCell.copy(second = currentCell.second - 1)
            true
        } else {
            false
        }
    }

    fun right(): Boolean {
        return if (currentCell.second < cols - 1) {
            currentCell = currentCell.copy(second = currentCell.second + 1)
            true
        } else {
            false
        }
    }

    fun upRight(): Boolean {
        return if (currentCell.first > 0 && currentCell.second < cols - 1) {
            currentCell =
                currentCell.copy(first = currentCell.first - 1, second = currentCell.second + 1)
            true
        } else {
            false
        }
    }

    fun upLeft(): Boolean {
        return if (currentCell.first > 0 && currentCell.second > 0) {
            currentCell =
                currentCell.copy(first = currentCell.first - 1, second = currentCell.second - 1)
            true
        } else {
            false
        }
    }

    fun downRight(): Boolean {
        return if (currentCell.first < rows - 1 && currentCell.second < cols - 1) {
            currentCell =
                currentCell.copy(first = currentCell.first + 1, second = currentCell.second + 1)
            true
        } else {
            false
        }
    }

    fun downLeft(): Boolean {
        return if (currentCell.first < rows - 1 && currentCell.second > 0) {
            currentCell =
                currentCell.copy(first = currentCell.first + 1, second = currentCell.second - 1)
            true
        } else {
            false
        }
    }

    fun currentValue(): T {
        return grid[currentCell.first][currentCell.second]
    }
}