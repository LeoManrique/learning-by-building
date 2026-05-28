use std::io::Empty;

use clap::Parser;

#[derive(Debug, Clone, Copy)]
enum Unit {
    Celsius,
    Fahrenheit,
    Kelvin,
}

#[derive(Parser)]
#[command(about = "Convert a value from a unit to another")]
struct Args {
    value: f64,
    from: String,
    to: String,
}

fn parse_unit(s: &str) -> Option<Unit> {
    match s {
        "C" => Some(Unit::Celsius),
        "F" => Some(Unit::Fahrenheit),
        "K" => Some(Unit::Kelvin),
        _ => None,
    }
}

fn main() {
    let args = Args::parse();
    let from = args.from;
    let to = args.to;
    println!("DEBUG: from={from}, to={to}");
    println!("Hi");
}
