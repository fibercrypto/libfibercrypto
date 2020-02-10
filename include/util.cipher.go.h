typedef struct
{
    GoString_ Address;
} util__GenericAddress;

typedef struct
{
    void *data;
    GoInt_ len;
    GoInt_ cap;
} core__Checksum;