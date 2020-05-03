CACHE = [i for i in range(1024)]

def levenshtein_distance(source: str, target: str):
    if len(source) == 0:
        return len(target)

    if len(target) == 0:
        return len(source)

    for i in range(len(target)+1):
        CACHE[i] = i

    for i, source_char in enumerate(source):
        next_dist = i + 1

        for j, target_char in enumerate(target):
            current_dist = next_dist

            dist_if_substitute = CACHE[j]
            if source_char != target_char:
                dist_if_substitute += 1

            dist_if_insert = current_dist + 1
            dist_if_delete = CACHE[j + 1] + 1

            next_dist = min(dist_if_substitute, dist_if_insert, dist_if_delete)

            CACHE[j] = current_dist

        CACHE[len(target)] = next_dist

    return CACHE[len(target)]


if __name__ == "__main__":
    from datetime import datetime

    with open("../sample.txt", "r") as sample_file:
        lines = sample_file.read().split("\n")

    def benchmark():
        for _ in range(10_000):
            last_value = ""
            for line in lines:
                levenshtein_distance(last_value, line)
                last_value = line

    start = datetime.now()
    benchmark()
    duration = datetime.now() - start

    print(f"{duration.total_seconds()}s")
