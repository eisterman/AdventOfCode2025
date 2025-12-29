import java.io.File

fun main() {
//    val content = File("../test.data").readText();
    val content = File("../input.data").readText();
    val rows = content.trim().lines()
    val emptyRow = rows.indexOf("")
    val freshRows = rows.slice(0..<emptyRow)
    val availableRows = rows.slice(emptyRow + 1..<rows.size)
    val freshRanges = freshRows.map { r -> r.split("-").map(String::toLong).let { it[0]..it[1] } }
    val toTestList = availableRows.map { it.toLong() }
    val freshIds = toTestList.filter { i -> freshRanges.any { range -> i in range } }
    val res1 = freshIds.size
    println("Problem 1: $res1")
    // Problem 2
    val result = mutableListOf<LongRange>()
    val rangesToCheck = freshRanges.sortedBy { it.last - it.first }.toMutableList()
    toCheck@ while (rangesToCheck.isNotEmpty()) {
        var range = rangesToCheck.removeLast()
        for (savedRange in result) {
            // Check reject
            if (savedRange.first <= range.first && range.last <= savedRange.last) {
                continue@toCheck
            }
            if (range.last in savedRange) {
                // Check restriction on the right of the range
                range = range.first..<savedRange.first  // same as ..(val+1)
            } else if (range.first in savedRange) {
                // Check restriction on the left of the range
                range = (savedRange.last + 1)..range.last
            }
            // Split can happen even with sort because we trim the range, and we can insert in the result very small results
            if (range.first <= savedRange.first && savedRange.last <= range.last) {
                range = range.first..<savedRange.first
                val splinterRange = (savedRange.last + 1)..range.last
                rangesToCheck.addLast(splinterRange)
            }
        }
        result.addLast(range)
    }
    val res2 = result.sumOf { it.last - it.first + 1 } // count cannot be used because it uses Int (32bit)
    println("Problem 2: $res2")
}