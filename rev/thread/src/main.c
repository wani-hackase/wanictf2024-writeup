#include <pthread.h>
#include <stdio.h>
#include <string.h>

#define LEN 45
#define MOD 3

// ロック変数
pthread_mutex_t mutex;

int inputs[LEN];
int flags[LEN];
int ans[LEN] = {168, 138, 191, 165, 765, 89,  222, 36,  101, 271, 222, 35,
                349, 66,  44,  222, 9,   101, 222, 81,  239, 319, 36,  83,
                349, 72,  83,  222, 9,   83,  331, 36,  101, 222, 54,  83,
                349, 18,  74,  292, 63,  95,  334, 213, 11};

void *thread1Function(void *arg) {
  int index = *(int *)arg; // 引数をintにキャスト
  int flag = 0;

  // indexが0なら0->1->2
  // indexが1なら1->2->0
  // という順番で処理が行われるようにしたい
  // indexが0でflagが0なら0
  // indexが0でflagが1なら1
  // indexが0でflagが2なら2
  // indexが1でflagが0なら1
  // indexが1でflagが1なら2
  // indexが1でflagが2なら0

  while (flag < MOD) {
    // ロックを取得
    pthread_mutex_lock(&mutex);

    int mod = (index + flags[index]) % MOD;
    if (mod == 0) {
      inputs[index] = inputs[index] * 3;
    }
    if (mod == 1) {
      inputs[index] = inputs[index] + 5;
    }
    if (mod == 2) {
      inputs[index] = inputs[index] ^ 127;
    }
    flags[index] = flags[index] + 1;
    flag = flags[index];

    // ロックを解放
    pthread_mutex_unlock(&mutex);
  }

  return NULL;
}

int main() {
  char inputString[LEN + 1];
  printf("FLAG: ");
  if (scanf("%45s", inputString) != 1) {
    printf("Failed to scan.\n");
    return 1;
  }

  if (strlen(inputString) != LEN) {
    printf("Incorrect.\n");
    return 1;
  }

  for (int i = 0; i < LEN; i++) {
    inputs[i] = inputString[i];
  }

  // ロックの初期化
  pthread_mutex_init(&mutex, NULL);

  pthread_t thread1[LEN];
  int threadNum[LEN]; // スレッド番号を格納する配列

  for (int i = 0; i < LEN; i++) {
    threadNum[i] = i;
    pthread_create(&thread1[i], NULL, thread1Function, &threadNum[i]);
  }
  for (int i = 0; i < LEN; i++) {
    pthread_join(thread1[i], NULL);
  }

  // ロックの破棄
  pthread_mutex_destroy(&mutex);

  for (int i = 0; i < LEN; i++) {
    if (inputs[i] != ans[i]) {
      printf("Incorrect.\n");
      return 1;
    }
  }

  printf("Correct!\n");

  return 0;
}
