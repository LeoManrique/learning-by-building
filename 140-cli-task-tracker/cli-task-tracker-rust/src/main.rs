use std::fs;
//use std::io::ErrorKind;

fn main() {
    let result = fs::read_to_string("./db/tasks.json");
    match result {
        Ok(contents) => println!("got string: {:?}", contents),
        //Err(e) if e.kind() == ErrorKind::NotFound => println!("no tasks file yet — starting fresh"),
        Err(e) => println!("got error: {:?}", e),
    }
}
