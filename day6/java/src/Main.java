
void main() throws IOException, RuntimeException {
//    var content = Files.readAllLines(Path.of("../test.data"));
    var content = Files.readAllLines(Path.of("../input.data"));
    if (content.getLast().isEmpty()) {
        content = content.subList(0, content.size()-1);
    }
    IO.println("Hello and welcome!");
    IO.println(content);
    int operatorRow = content.size() - 1;
    var rows = content.subList(0,operatorRow).stream().map((x) -> Arrays.stream(x.trim().split("\\s+")).map(Long::parseLong).toList()).toList();
    var operators = content.get(operatorRow).trim().split("\\s+");
    long res = 0;
    for(int i = 0; i < operators.length; i++) {
        int finalI = i;  // Needed because `i` can vary and closures need "final variables" that doesn't vary
        var colRes = switch (operators[i]) {
            case "*" -> rows.stream().reduce(1L, (acc, row) -> acc * row.get(finalI), Long::sum);
            case "+" -> rows.stream().reduce(0L, (acc, row) -> acc + row.get(finalI), Long::sum);
            default -> throw new RuntimeException("Wrong operator");
        };
        res += colRes;
    }
    IO.println(String.format("Problem 1: %d", res));
}
