%inline {
	int equalSlices(GoSlice* slice1, GoSlice* slice2, int elem_size){
	  if(slice1->len != slice2->len)
		return 0;
	  return memcmp(slice1->data, slice2->data, slice1->len * elem_size) == 0;
	}
	int equalTransactions(coin__Transaction* t1, coin__Transaction* t2){
		if( t1->Length != t2->Length || t1->Type != t2->Type ){
			return 0;
		}
		if( memcmp(&t1->InnerHash, &t2->InnerHash, sizeof(cipher__SHA256)) != 0 )
			return 0;
		if(!equalSlices((GoSlice*)&t1->Sigs, (GoSlice*)&t2->Sigs, sizeof(cipher__Sig)))
			return 0;
		if(!equalSlices((GoSlice*)&t1->In, (GoSlice*)&t2->In, sizeof(cipher__SHA256)))
			return 0;
		if(!equalSlices((GoSlice*)&t1->Out, (GoSlice*)&t2->Out, sizeof(coin__TransactionOutput)))
			return 0;
		return 1;
	}
}
