#include <stdio.h>

int constructFlag() {
  char flag[44];
  int enc[44] = {946, 932, 952, 938, 960, 932, 980, 962, 1008, 996, 984, 992, 1000, 562, 964, 566, 992, 974, 976, 560, 984, 964, 968, 556, 964, 974, 930, 988, 1008, 528, 986, 1000, 556, 568, 984, 542, 982, 562, 1006, 572, 572, 900, 624, 872};

  for (int i = 0; i < 44; i++) {
    enc[i] = enc[i] / 2;
  }

  for (int i = 0; i < 44; i++) {
    enc[i] = enc[i] ^ 415;
  }

  for (int i = 0; i < 44; i++) {
    flag[i] = enc[i] - i;
  }

  printf("Processing completed!");

  return 0;
}