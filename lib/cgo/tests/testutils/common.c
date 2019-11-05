#include "common.h"

#include <string.h>
#include <stdio.h>

#include "libfibercrypto.h"
#include "fcerrors.h"

int MEMPOOLIDX = 0;
void *MEMPOOL[1024 * 256];

GoInt_ isU8Eq(GoUint8_ p1[], GoUint8_ p2[], size_t len)
{
    for (GoInt i = 0; i < len; i++) {
        if (p1[i] != p2[i]) {
            return 0;
        }
    }
    return 1;
}

void* registerMemCleanup(void* p)
{
    int i;
    for (i = 0; i < MEMPOOLIDX; i++) {
        if (MEMPOOL[i] == NULL) {
            MEMPOOL[i] = p;
            return p;
        }
    }
    MEMPOOL[MEMPOOLIDX++] = p;
    return p;
}

int copyGoSlice_toGoSlice(GoSlice* pdest, GoSlice_* psource, int elem_size)
{
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
