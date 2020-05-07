mod levenshtein_distance;

use levenshtein_distance::levenshtein_distance;

fn main() {

    let mut lines: Vec<&str> = include_str!("../../sample.txt").split('\n').collect();
    lines.insert(0, "");
    let lines: Vec<Vec<char>> = lines.iter().map(|line| line.chars().collect()).collect();

    let benchmark = || {
        for _ in 0..10000 {
            lines.iter().zip(lines[1..].iter())
                .for_each(|(a, b)| {
                    levenshtein_distance(a, b);
                });
        }
    };

    use std::time::Instant;
    let now = Instant::now();

    {
        benchmark();
    }

    let elapsed = now.elapsed();
    let sec = (elapsed.as_secs() as f64) + (elapsed.subsec_nanos() as f64 / 1000_000_000.0);
    print!("{}", sec);

    // check
    let answers: Vec<String> = (1..lines.len()-1)
        .map(|i| levenshtein_distance(&lines[i], &lines[i+1]))
        .map(|dist| dist.to_string())
        .collect();
    eprintln!("{}", answers.join(","));
}
