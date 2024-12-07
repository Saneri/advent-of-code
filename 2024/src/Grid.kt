class Grid<T>(var grid: Array<Array<T>>, var currentCell: Pair<Int, Int>) {
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

    fun checkUp(): T? {
        return if (currentCell.first > 0) {
            grid[currentCell.first - 1][currentCell.second]
        } else {
            null
        }
    }

    fun checkDown(): T? {
        return if (currentCell.first < rows - 1) {
            grid[currentCell.first + 1][currentCell.second]
        } else {
            null
        }
    }

    fun checkLeft(): T? {
        return if (currentCell.second > 0) {
            grid[currentCell.first][currentCell.second - 1]
        } else {
            null
        }
    }

    fun checkRight(): T? {
        return if (currentCell.second < cols - 1) {
            grid[currentCell.first][currentCell.second + 1]
        } else {
            null
        }
    }

    fun currentValue(): T {
        return grid[currentCell.first][currentCell.second]
    }
}