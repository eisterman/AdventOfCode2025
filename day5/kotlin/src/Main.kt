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
}