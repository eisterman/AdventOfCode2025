use std::fs;

fn main() -> std::io::Result<()> {
    // let x = fs::read_to_string("../test.data")?;
    let x = fs::read_to_string("../input.data")?;
    let banks: Vec<&str> = x.split_whitespace().collect();
    let mut result = 0;
    for &bank in banks.iter() {
        // Find greater number in all bank except last character
        // Find in the string after the number the greatest
        let digits: Vec<u32> = bank.chars().map(|c| c.to_digit(10).unwrap()).collect();
        let first_battery = digits[0..bank.len() - 1].iter().max().unwrap();
        let first_battery_idx = digits.iter().position(|r| r == first_battery).unwrap();
        let second_battery = digits.iter().skip(first_battery_idx + 1).max().unwrap();

        result += first_battery * 10 + second_battery;
    }
    println!("Problem 1: {}", result);
    // Problem 2
    let mut result2 = 0;
    for &bank in banks.iter() {
        let digits: Vec<u32> = bank.chars().map(|c| c.to_digit(10).unwrap()).collect();
        result2 += joltage(&digits, 12);
    }
    println!("Problem 2: {}", result2);
    Ok(())
}

fn joltage(bank: &[u32], n: usize) -> u64 {
    if n > 0 {
        let battery = bank[0..bank.len() - (n - 1)].iter().max().unwrap();
        let battery_idx = bank.iter().position(|r| r == battery).unwrap();
        return (*battery as u64 * 10_u64.pow(n as u32 - 1))
            + joltage(&bank[battery_idx + 1..], n - 1);
    } else {
        return 0;
    }
}
