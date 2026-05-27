use clap::Parser;
use rand::RngExt;
use std::cmp::Ordering;
use std::io::{self, BufRead};

#[derive(Parser)]
#[command(about = "Guess a secret number in a configurable range")]
struct Args {
    #[arg(long, default_value_t = 1)]
    lower: i32,

    #[arg(long, default_value_t = 100)]
    upper: i32,
}

fn main() {
    let args = Args::parse();

    let lower_limit = args.lower;
    let upper_limit = args.upper;

    if lower_limit > upper_limit {
        eprintln!("lower ({lower_limit}) should not exceed upper ({upper_limit})");
        std::process::exit(2);
    }

    let mut number_of_attempts = 0;

    let mut rng = rand::rng();
    let picked_number = rng.random_range(lower_limit..=upper_limit);
    println!("Welcome to Number Guessing Game!");
    println!("I've picked a number between {lower_limit} and {upper_limit}, can you guess it?");

    //game logic
    loop {
        let mut line = String::new();
        match io::stdin().lock().read_line(&mut line) {
            Ok(0) => {
                eprintln!("No more input, goodbye!");
                std::process::exit(1);
            }
            Ok(_) => {}
            Err(e) => {
                eprintln!("input error: {e}");
                std::process::exit(1);
            }
        }
        let user_picked_number: i32 = match line.trim().parse() {
            Ok(n) => n,
            Err(_) => {
                println!("That's not a valid number, try again:");
                continue;
            }
        };

        if user_picked_number > upper_limit || user_picked_number < lower_limit {
            println!("That's not between {lower_limit} and {upper_limit}");
            continue;
        }

        number_of_attempts += 1;

        match user_picked_number.cmp(&picked_number) {
            Ordering::Less => println!("Higher :D, try again:"),
            Ordering::Greater => println!("Lower :P, try again:"),
            Ordering::Equal => {
                println!(
                    "That's correct! The right number was {picked_number}. In only {number_of_attempts} attempts"
                );
                break;
            }
        }
    }
}
