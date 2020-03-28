#ifndef FCCALLBACK_H
#define FCCALLBACK_H

// typedef GoUint32_ (*PasswordReaderFunc)(GoString_ pString,
// KeyValueStore__Handle pKeyValue, GoString_ *pStringOut, void *context);

static inline GoUint32_ callPasswordReader(PasswordReader *passwordReader,
                                           GoString_ pString,
                                           KeyValueStore__Handle pKeyValue,
                                           GoString_ *pStringOut) {
  return passwordReader->callback(pString, pKeyValue, pStringOut,
                                  passwordReader->context);
}

#endif