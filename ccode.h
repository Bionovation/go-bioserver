#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>

void test(){
    printf("hello there!\n");
}

int readfile(char* path, char* buf, int32_t len){
    FILE* pf = fopen(path, "rb");
    if(pf == NULL) return 0;
    fseek(pf, 0, SEEK_END);
    long lSize = ftell(pf);
    if(lSize > len) return 0;
    rewind(pf);
    fread(path, sizeof(char), lSize, pf);
    fclose(pf);
    return lSize;
}