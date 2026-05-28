use clap::Parser;

#[derive(Parser)]
#[command(about = "Convert a value from a unit to another")]
struct Args {
    #[arg(long)]
    value: i32,

    #[arg()]
    from: String,

    #[arg()]
    to: String,
}

fn main() {
    println!("Hello, world!");
}
