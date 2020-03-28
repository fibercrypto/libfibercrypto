typedef struct {
  const char *p;
  GoInt_ n;
} core__UID;

typedef struct {
  GoString_ Name;
  GoString_ Ticker;
  GoString_ Family;
  BOOL HasBip44;
  GoInt32_ Bip44CoinType;
  GoInt32_ Accuracy;
} core__AltcoinMetadata;