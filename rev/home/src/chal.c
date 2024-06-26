#define _GNU_SOURCE
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <sys/ptrace.h>
#include <sys/types.h>
#include <sys/wait.h>

void constructFlag();

int main() {
    char current_directory[1024];
    if (getcwd(current_directory, sizeof(current_directory)) != NULL) {
        if (strstr(current_directory, "Service") != NULL) {
            printf("Check passed!\n");
            if (ptrace(PTRACE_TRACEME, 0, NULL, NULL) == -1) {
                printf("Debugger detected!\n");
                return EXIT_FAILURE;
            }
            constructFlag();
        }
        else {
            printf(";)\n");
        }
    } else {
        perror("Error");
        return EXIT_FAILURE;
    }

    return EXIT_SUCCESS;
}