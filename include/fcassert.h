#ifndef LIBFC_ASSERT_H
#define LIBFC_ASSERT_H

#include "fcerrors.h"
#include "libfibercrypto.h"

extern GoInt isGoStringEq(GoString string1, GoString string2);

extern GoInt isGoString_Eq(GoString_ string1, GoString_ string2);

extern GoInt isutil__GenericAddress_Eq(util__GenericAddress p0, util__GenericAddress p1);

extern GoInt isGoSliceEq(GoSlice *slice1, GoSlice *slice2);

extern GoInt isGoSlice_Eq(GoSlice_ *slice1, GoSlice_ *slice2);

#endif // LIBFC_ASSERT_H
