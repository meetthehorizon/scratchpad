use std::{cmp::Ordering, io};

fn main() {
    println!("Welcome to the guessing game!");
    println!("You have to guess a secret number between 1 and 100.");

    let secret_number: u32 = rand::random_range(1..=100);
    println!("[DEBUG] The secret number is {secret_number}");

    'game: loop {
        println!("Please Enter your guess:");

        let mut guess = String::new();
        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line!");

        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => {
                println!("Not a valid number, please try again");
                continue 'game;
            }
        };
        match guess.cmp(&secret_number) {
            Ordering::Less => println!("oh no! your guess was too low"),
            Ordering::Greater => println!("oh no! your guess was too high"),
            Ordering::Equal => {
                println!("you win!");
                break;
            }
        }
    }
}
