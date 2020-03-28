#include "check.h"
#include "fcassert.h"
#include "fcstring.h"
#include <string.h>

GoInt equalSlices(GoSlice *slice1, GoSlice *slice2, int elem_size)
{
    if (slice1->len != slice2->len)
        return 0;
    return (memcmp(slice1->data, slice2->data, slice1->len * elem_size) == 0);
}

GoInt equalSlices_(GoSlice_ *slice1, GoSlice_ *slice2, int elem_size)
{
    if (slice1->len != slice2->len)
        return 0;
    return (memcmp(slice1->data, slice2->data, slice1->len * elem_size) == 0);
}

GoInt isGoStringEq(GoString string1, GoString string2)
{
    return (string1.n == string2.n) &&
           (strncmp(string1.p, string2.p, string1.n) == 0);
}

GoInt isGoString_Eq(GoString_ string1, GoString_ string2)
{
    return (string1.n == string2.n) &&
           (strcmp(string1.p, string2.p) == 0);
}

GoInt isGoSliceEq(GoSlice *slice1, GoSlice *slice2)
{
    return (slice1->len == slice2->len) &&
           (memcmp(slice1->data, slice2->data, slice1->len) == 0);
}

GoInt isGoSlice_Eq(GoSlice_ *slice1, GoSlice_ *slice2)
{
    return (slice1->len == slice2->len) &&
           (memcmp(slice1->data, slice2->data, slice1->len) == 0);
}

GoInt isutil__GenericAddress_Eq(util__GenericAddress p0, util__GenericAddress p1)
{
    return isGoString_Eq(p0.Address, p1.Address);
}