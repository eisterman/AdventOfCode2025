require IEx

defmodule Solution do
  def part1(data) do
    problem1 =
      data
      |> Enum.map(fn x ->
        [a, b] = x |> Enum.map(&String.to_integer(&1))
        Enum.to_list(a..b) |> Enum.map(&Integer.to_string(&1))
      end)
      |> List.flatten()
      |> Enum.flat_map(fn x -> if !is_valid(x), do: [x], else: [] end)
      |> Enum.map(&String.to_integer(&1))
      |> Enum.sum()

    IO.puts("Solution of Part 1: #{problem1}")
  end

  def is_valid(n) do
    if rem(String.length(n), 2) == 0 do
      {a, b} = String.split_at(n, div(String.length(n), 2))
      a != b
    else
      true
    end
  end

  def part2(data) do
    problem2 =
      data
      |> Enum.map(fn x ->
        [a, b] = x |> Enum.map(&String.to_integer(&1))
        Enum.to_list(a..b) |> Enum.map(&Integer.to_string(&1))
      end)
      |> List.flatten()
      |> Enum.flat_map(fn x -> if !is_valid2(x), do: [x], else: [] end)
      |> Enum.map(&String.to_integer(&1))
      |> Enum.sum()

    IO.puts("Solution of Part 2: #{problem2}")
  end

  def is_valid2(x) do
    len = String.length(x)

    if len < 2 do
      true
    else
      good_splits = 1..(len - 1) |> Enum.filter(&(rem(len, &1) == 0))
      Enum.all?(good_splits, &is_valid2_onesplit(x, &1))
    end
  end

  def is_valid2_onesplit(x, split_every) do
    splitted = string_splitter(x, split_every)
    Enum.any?(splitted, fn v -> v != hd(splitted) end)
  end

  def string_splitter(x, split_every) do
    if String.length(x) >= split_every do
      {h, t} = String.split_at(x, split_every)
      [h | string_splitter(t, split_every)]
    else
      []
    end
  end

  def main() do
    # input = File.read!(Path.join(__DIR__, "../test.in"))
    input = File.read!(Path.join(__DIR__, "../input.in"))

    data =
      input
      |> String.trim()
      |> String.split(",")
      |> Enum.map(&String.split(&1, "-"))

    # part1(data)
    part2(data)
  end
end

Solution.main()
