
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
    // Problem 2
    var res2 = 0L;
    var buffers = Stream.generate(ArrayList<Character>::new).limit(content.size()).toList();
    for(int i = 0; i < content.getFirst().length(); i++) {
        int finalI = i;
        var chars = content.stream().map((row) -> row.charAt(finalI)).toList();
        if (chars.stream().allMatch((c) -> c == ' ')) {
            // all space, separation
            // Execute the operation on buffer and empty them
            res2 += processBuffers(buffers);
            continue;
        }
        IntStream.range(0, chars.size()).forEach((j) -> {
            buffers.get(j).addLast(chars.get(j));
        });
    }
    // Operate again if buffers are not empty
    if (!buffers.getFirst().isEmpty()) {
        res2 += processBuffers(buffers);
    }
    IO.println(String.format("Problem 2: %d", res2));
}

long processBuffers(List<ArrayList<Character>> buffers) {
    var operator = buffers.getLast().getFirst(); // operator
    var len = buffers.getFirst().size();
    var numbers = LongStream.range(0, len).map((k) -> Long.parseLong(buffers.subList(0, buffers.size()-1).stream().collect(StringBuilder::new, (acc, cl) -> acc.append(cl.get((int)k)), StringBuilder::append).toString().trim())).boxed().toList();
    var sessionRes = numbers.stream().reduce(operator == '*' ? 1L : 0L, (acc, x) -> operator == '*'? acc*x : acc+x, Long::sum);
    for(var b: buffers) {
        b.clear();
    }
    return sessionRes;
}
