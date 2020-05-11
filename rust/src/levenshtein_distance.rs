pub fn levenshtein_distance(source: &str, target: &str) -> usize {
    if source.is_empty() {
        return target.chars().count();
    }

    if target.is_empty() {
        return source.chars().count();
    }

    let target_len = target.chars().count();
    let mut cache: Vec<usize> = (0..=target_len).collect();

    for (i, source_char) in source.chars().enumerate() {
        let mut next_dist = i + 1;

        for (j, target_char) in target.chars().enumerate() {
            let current_dist = next_dist;

            let dist_if_substitute = {
                if source_char == target_char {
                    cache[j]
                } else {
                    cache[j] + 1
                }
            };

            let dist_if_insert = current_dist + 1;
            let dist_if_delete = cache[j + 1] + 1;

            next_dist = min(
                dist_if_delete,
                min(dist_if_insert, dist_if_substitute),
            );

            cache[j] = current_dist;
        }

        cache[target_len] = next_dist;
    }

    cache[target_len]
}

fn min(a: usize, b: usize) -> usize {
    if a < b { a } else { b }
}
