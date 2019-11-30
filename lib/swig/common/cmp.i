%inline {
	int equalSlices(GoSlice* slice1, GoSlice* slice2, int elem_size){
	  if(slice1->len != slice2->len)
		return 0;
	  return memcmp(slice1->data, slice2->data, slice1->len * elem_size) == 0;
	}
}
