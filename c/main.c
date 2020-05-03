#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <strings.h>

int min(int a, int b)
{
  return a > b ? b : a;
}

int min_3(int a, int b, int c)
{
  return min(a, min(b, c));
}

int levenshtein_distance(const char *source, const char *target)
{
  size_t source_len = strlen(source);
  size_t target_len = strlen(target);

  if (source_len == 0) return target_len;
  if (target_len == 0) return source_len;

  int *cache = malloc(sizeof(int)*(target_len+1));

  for (int i = 0; i < source_len; i++) {
    int next_dist = i + 1;

    for (int j = 0; j < target_len; j++) {
      int current_dist = next_dist;

      int dist_if_subst = cache[j];
      if (source[i] != target[j]) {
        dist_if_subst++;
      }

      int dist_if_insert = current_dist + 1;
      int dist_if_delete = cache[j + 1] + 1;

      next_dist = min_3(dist_if_subst, dist_if_insert, dist_if_delete);

      cache[j] = current_dist;
    }

    cache[target_len] = next_dist;
  }

  return cache[target_len];
}

int main()
{
  static char *lines[1024];

  FILE* stream = fopen("sample.txt", "r");
  if (stream == NULL) {
    exit(EXIT_FAILURE);
  }

  int n_lines;
	size_t len = 0;
  for (int i = 0; i < 1024; i++) {
    char* line = malloc(sizeof(char)*512);
    if (getline(&line, &len, stream) == -1) {
      n_lines = i-1;
      free(line);
      break;
    }

    lines[i] = line;
	}

  for (int i = 0; i < 10000; i++) {
    char* last_line = NULL;
    for (int line = 0; line < n_lines; line++) {
      levenshtein_distance(last_line, lines[line]);
      last_line = lines[line];
    }
  }

  return EXIT_SUCCESS;
}
