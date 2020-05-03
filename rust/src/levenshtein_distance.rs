pub fn levenshtein_distance(source: &str, target: &str) -> usize {
    if source.is_empty() {
        return target.len();
    }

    if target.is_empty() {
        return source.len();
    }

    let mut cache: Vec<usize> = (0..=target.len()).collect();

    for (i, source_char) in source.chars().enumerate() {
        let mut next_dist = i + 1;

        for (j, target_char) in target.chars().enumerate() {
            let current_dist = next_dist;

            let mut dist_if_substitute = cache[j];
            if source_char != target_char {
                dist_if_substitute += 1;
            }

            let dist_if_insert = current_dist + 1;
            let dist_if_delete = cache[j + 1] + 1;

            next_dist = std::cmp::min(
                dist_if_substitute,
                std::cmp::min(dist_if_insert, dist_if_delete),
            );

            cache[j] = current_dist;
        }

        cache[target.len()] = next_dist;
    }

    cache[target.len()]
}
