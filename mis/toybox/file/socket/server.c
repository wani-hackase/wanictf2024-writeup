#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/sendfile.h>

void init() {
  setvbuf(stdin, NULL, _IONBF, 0);
  setvbuf(stdout, NULL, _IONBF, 0);
  setvbuf(stderr, NULL, _IONBF, 0);
  alarm(30);
}

void banner() {
  FILE *fp = fopen("banner.txt", "r");
  if (fp == NULL) {
    printf("banner not available\n");
    exit(1);
  }
  sendfile(fileno(stdout), fileno(fp), NULL, 1024);
}

void check_flag() {
  FILE *fp = fopen("flag.txt", "r");
  if (fp == NULL) {
    printf("flag not available\n");
    exit(1);
  }
}

int main() {
  init();
  banner();
  check_flag();

  char id[38];
  printf("ID > ");
  if (fgets(id, sizeof(id), stdin) == NULL) {
    printf("cannot read from stdin\n");
    return 1;
  }

  for (int i = 0; i < sizeof(id); i++) {
    if (id[i] == '/') {
      printf("invalid ID\n");
      return 1;
    } else if (id[i] == '\n') {
      id[i] = '\0';
      break;
    }
  }

  char path[128];
  snprintf(path, sizeof(path), "./data/%s", id);

  if (execl("./sandbox", "./sandbox", path, NULL) == -1) {
    printf("exec failed\n");
    return 1;
  }

  return 0;
}
