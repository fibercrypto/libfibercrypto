#include "json.h"
#include "libfibercrypto.h"
#include "skyerrors.h"
#include "skytest.h"
#include "skytypes.h"
#include <stdio.h>
#include <stdlib.h>
#include <sys/stat.h>
#include <time.h>
#include <unistd.h>
#define BUFFER_SIZE 1024
#define STRING_SIZE 128
#define JSON_FILE_SIZE 4096
#define JSON_BIG_FILE_SIZE 102400

// Define function SKY_handle_close to avoid including libskycoin.h
void SKY_handle_close(Handle p0);

extern int MEMPOOLIDX;
extern void *MEMPOOL[1024 * 256];

int JSONPOOLIDX = 0;
json_value *JSON_POOL[128];

int HANDLEPOOLIDX = 0;
Handle HANDLE_POOL[128];

int WALLETPOOLIDX = 0;

int stdout_backup;
int pipefd[2];

void freeRegisteredMemCleanup(void *p) {
  int i;
  for (i = 0; i < MEMPOOLIDX; i++) {
    if (MEMPOOL[i] == p) {
      free(p);
      MEMPOOL[i] = NULL;
      break;
    }
  }
}

int registerJsonFree(void *p) {
  int i;
  for (i = 0; i < JSONPOOLIDX; i++) {
    if (JSON_POOL[i] == NULL) {
      JSON_POOL[i] = p;
      return i;
    }
  }
  JSON_POOL[JSONPOOLIDX++] = p;
  return JSONPOOLIDX - 1;
}

void freeRegisteredJson(void *p) {
  int i;
  for (i = 0; i < JSONPOOLIDX; i++) {
    if (JSON_POOL[i] == p) {
      JSON_POOL[i] = NULL;
      json_value_free((json_value *)p);
      break;
    }
  }
}

int registerHandleClose(Handle handle) {
  int i;
  for (i = 0; i < HANDLEPOOLIDX; i++) {
    if (HANDLE_POOL[i] == 0) {
      HANDLE_POOL[i] = handle;
      return i;
    }
  }
  HANDLE_POOL[HANDLEPOOLIDX++] = handle;
  return HANDLEPOOLIDX - 1;
}

void closeRegisteredHandle(Handle handle) {
  int i;
  for (i = 0; i < HANDLEPOOLIDX; i++) {
    if (HANDLE_POOL[i] == handle) {
      HANDLE_POOL[i] = 0;
      FC_handle_close(handle);
      break;
    }
  }
}

void cleanupMem() {
  int i;

  void **ptr;
  for (i = MEMPOOLIDX, ptr = MEMPOOL; i; --i) {
    if (*ptr)
      memset(ptr, 0, sizeof(void *));
    ptr++;
  }
  for (i = JSONPOOLIDX, ptr = (void *)JSON_POOL; i; --i) {
    if (*ptr)
      json_value_free(*ptr);
    ptr++;
  }
  for (i = 0; i < HANDLEPOOLIDX; i++) {
    if (HANDLE_POOL[i])
      FC_handle_close(HANDLE_POOL[i]);
  }
}

json_value *loadJsonFile(const char *filename) {
  FILE *fp;
  struct stat filestatus;
  int file_size;
  char *file_contents;
  json_char *json;
  json_value *value;

  if (stat(filename, &filestatus) != 0) {
    return NULL;
  }
  file_size = filestatus.st_size;
  file_contents = (char *)malloc(filestatus.st_size);
  if (file_contents == NULL) {
    return NULL;
  }
  fp = fopen(filename, "rt");
  if (fp == NULL) {
    free(file_contents);
    return NULL;
  }
  if (fread(file_contents, file_size, 1, fp) != 1) {
    fclose(fp);
    free(file_contents);
    return NULL;
  }
  fclose(fp);

  json = (json_char *)file_contents;
  value = json_parse(json, file_size);
  free(file_contents);
  return value;
}

void setup(void) { srand(time(NULL)); }

void teardown(void) { cleanupMem(); }

// TODO: Move to libsky_io.c
void fprintbuff(FILE *f, void *buff, size_t n) {
  unsigned char *ptr = (unsigned char *)buff;
  fprintf(f, "[ ");
  for (; n; --n, ptr++) {
    fprintf(f, "%02d ", *ptr);
  }
  fprintf(f, "]");
}

int parseBoolean(const char *str, int length) {
  int result = 0;
  if (length == 1) {
    result = str[0] == '1' || str[0] == 't' || str[0] == 'T';
  } else {
    result = strncmp(str, "true", length) == 0 ||
             strncmp(str, "True", length) == 0 ||
             strncmp(str, "TRUE", length) == 0;
  }
  return result;
}

int copySlice(GoSlice_ *pdest, GoSlice_ *psource, int elem_size) {
  pdest->len = psource->len;
  pdest->cap = psource->len;
  int size = pdest->len * elem_size;
  pdest->data = malloc(size);
  if (pdest->data == NULL)
    return FC_ERROR;
  registerMemCleanup(pdest->data);
  memcpy(pdest->data, psource->data, size);
  return FC_OK;
}

int cutSlice(GoSlice_ *slice, int start, int end, int elem_size,
             GoSlice_ *result) {
  int size = end - start;
  if (size <= 0)
    return FC_ERROR;
  void *data = malloc(size * elem_size);
  if (data == NULL)
    return FC_ERROR;
  registerMemCleanup(data);
  result->data = data;
  result->len = size;
  result->cap = size;
  char *p = slice->data;
  p += (elem_size * start);
  memcpy(data, p, elem_size * size);
  return FC_OK;
}

int concatSlices(GoSlice_ *slice1, GoSlice_ *slice2, int elem_size,
                 GoSlice_ *result) {
  int size1 = slice1->len;
  int size2 = slice2->len;
  int size = size1 + size2;
  if (size <= 0)
    return FC_ERROR;
  void *data = malloc(size * elem_size);
  if (data == NULL)
    return FC_ERROR;
  registerMemCleanup(data);
  result->data = data;
  result->len = size;
  result->cap = size;
  char *p = data;
  if (size1 > 0) {
    memcpy(p, slice1->data, size1 * elem_size);
    p += (elem_size * size1);
  }
  if (size2 > 0) {
    memcpy(p, slice2->data, size2 * elem_size);
  }
  return FC_OK;
}
