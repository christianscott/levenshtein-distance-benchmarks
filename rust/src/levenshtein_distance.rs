#[derive(Debug, Default)]
pub struct LevenshteinDistance {
    source: Vec<char>,
    target: Vec<char>,
    cache: Vec<usize>,
}

impl LevenshteinDistance {
    pub fn distance(&mut self, source: &str, target: &str) -> usize {
        if source.is_empty() {
            return target.len();
        }
        if target.is_empty() {
            return source.len();
        }

        self.source.clear();
        self.source.extend(source.chars());
        self.target.clear();
        self.target.extend(target.chars());
        self.cache.clear();
        self.cache.extend(0..=self.target.len());

        for (i, source_char) in self.source.iter().enumerate() {
            let mut next_dist = i + 1;

            for (j, target_char) in self.target.iter().enumerate() {
                let current_dist = next_dist;

                let mut dist_if_substitute = self.cache[j];
                if source_char != target_char {
                    dist_if_substitute += 1;
                }

                let dist_if_insert = current_dist + 1;
                let dist_if_delete = self.cache[j + 1] + 1;

                next_dist = min(
                    dist_if_delete,
                    min(dist_if_insert, dist_if_substitute),
                );

                self.cache[j] = current_dist;
            }

            self.cache[target.len()] = next_dist;
        }

        self.cache[target.len()]
    }
}

fn min(a: usize, b: usize) -> usize {
    if a < b { a } else { b }
}
