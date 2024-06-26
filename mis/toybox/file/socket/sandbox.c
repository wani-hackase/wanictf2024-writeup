#include <stdio.h>
#include <unistd.h>
#include <seccomp.h>

void init() {
  setvbuf(stdin, NULL, _IONBF, 0);
  setvbuf(stdout, NULL, _IONBF, 0);
  setvbuf(stderr, NULL, _IONBF, 0);
  alarm(3);
}

int main(int argc, char **argv) {
  init();

  if (argc != 2) {
    printf("usage: %s <path>\n", argv[0]);
    return 1;
  }
  char *executable_path = argv[1];

  scmp_filter_ctx ctx = seccomp_init(SCMP_ACT_KILL);
  if (ctx == NULL) {
    printf("seccomp_init failed\n");
    return 1;
  }

  if (seccomp_rule_add(ctx, SCMP_ACT_ALLOW, SCMP_SYS(read), 0) < 0 ||
      seccomp_rule_add(ctx, SCMP_ACT_ALLOW, SCMP_SYS(write), 0) < 0 ||
      seccomp_rule_add(ctx, SCMP_ACT_ALLOW, SCMP_SYS(stat), 0) < 0 ||
      seccomp_rule_add(ctx, SCMP_ACT_ALLOW, SCMP_SYS(fstat), 0) < 0 ||
      seccomp_rule_add(ctx, SCMP_ACT_ALLOW, SCMP_SYS(lstat), 0) < 0 ||
      seccomp_rule_add(ctx, SCMP_ACT_ALLOW, SCMP_SYS(access), 0) < 0 ||
      seccomp_rule_add(ctx, SCMP_ACT_ALLOW, SCMP_SYS(getpid), 0) < 0 ||
      seccomp_rule_add(ctx, SCMP_ACT_ALLOW, SCMP_SYS(exit), 0) < 0 ||
      seccomp_rule_add(ctx, SCMP_ACT_ALLOW, SCMP_SYS(execve), 1,
                       SCMP_A0_64(SCMP_CMP_EQ, (scmp_datum_t) executable_path)) < 0) {
    printf("seccomp_rule_add failed\n");
    return 1;
  }

  if (seccomp_load(ctx) < 0) {
    printf("seccomp_load failed\n");
    return 1;
  }

  if (execl(executable_path, executable_path, NULL) == -1) {
    printf("exec failed\n");
    return 1;
  }

  return 0;
}
