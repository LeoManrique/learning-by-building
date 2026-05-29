use clap::Parser;
mod convertors;
use convertors::convert;

#[derive(Parser)]
#[command(about = "Convert a value from a unit to another")]
struct Args {
    value: f64,
    #[arg(long, default_value_t = 2)]
    precision: usize,
    #[arg(long)]
    from: String,
    #[arg(long)]
    to: String,
}

fn main() {
    let args = Args::parse();
    let precision = args.precision;
    match convert(&args.from, &args.to, args.value) {
        Ok(x) => println!("{x:.precision$}"),
        Err(err) => {
            eprintln!("{err}");
            std::process::exit(2);
        }
    }
}
