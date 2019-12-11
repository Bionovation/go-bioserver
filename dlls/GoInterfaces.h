#include <stdlib.h>

//#define BIOAPI __declspec(dllimport)

/*
int ReadSlideInfo(char* path, char* slideInfo);
int ReadSlideTile(char* path, int level, int x, int y, int tileSize, char* buf, int bufSize);
int ReadSlideNail(char* path, char* jpegBuf, int bufSize);
int CloseSlide(char* path);
*/


int ReadSlideInfo(int* pathRnue, char* slideInfo);
int ReadSlideTile(int* pathRnue, int level, int x, int y, int tileSize, char* buf, int bufSize);
int CloseSlide(int* pathRnue);

