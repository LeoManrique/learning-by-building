use rand::RngExt;

fn main() {
    // parse flags
    /*let lowerLimit, upperLimit int
	flag.IntVar(&lowerLimit, "lower", 1, "lower bound (inclusive)")
	flag.IntVar(&upperLimit, "upper", 100, "upper bound (inclusive)")
	flag.Parse()*/
    let lower_limit = 1;
    let upper_limit = 100;

    //check flags
	/*if lowerLimit > upperLimit {
		fmt.Fprintf(os.Stderr, "lower (%d) should not exceed upper (%d)\n", lowerLimit, upperLimit)
		os.Exit(2)
	}*/

    //initalize vars
	/*numberOfAttempts := 0
	notYetGuessed := true
	pickedNumber := lowerLimit + rand.IntN(upperLimit-lowerLimit+1)

	scanner := bufio.NewScanner(os.Stdin)*/

    let mut number_of_attempts = 0;

    let mut rng = rand::rng();
    let picked_number = rng.random_range(lower_limit..=upper_limit);

    //welcome user
	/*fmt.Println("Welcome to Number Guessing Game!")
	fmt.Printf("I've picked a number between %d and %d, can you guess it?\n",
		lowerLimit,
		upperLimit)*/
    println!("Welcome to Number Guessing Game!");
    println!("I've picked a number between {lower_limit} and {upper_limit}, can you guess it?");

    //game logic
    loop {
        let user_picked_number = 10;
        number_of_attempts += 1;
        if user_picked_number == picked_number || true {
            println!("That's the right number! In only {number_of_attempts} attempts");
            break;
        }
    }
	/*for notYetGuessed {
		if ok := scanner.Scan(); !ok {
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "input error:", err)
				os.Exit(1)
			}
			fmt.Println("No more input — goodbye!")
			os.Exit(1)
		}
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("That's not a valid number, try again:")
			continue
		}

		if num > upperLimit || num < lowerLimit {
			fmt.Printf("That's not between %d and %d, try again:\n", lowerLimit, upperLimit)
			continue
		}

		numberOfAttempts++

		if num == pickedNumber {
			fmt.Printf("You guessed it after %d attempts! It was %d. See you! :)\n", numberOfAttempts, pickedNumber)
			notYetGuessed = false
		} else if num > pickedNumber {
			fmt.Println("Lower :P, try again:")
		} else {
			fmt.Println("Higher :D, try again:")
		}
	}*/
}
