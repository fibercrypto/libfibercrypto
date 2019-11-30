/*GoInt64* as function return typemap*/
%typemap(argout) GoInt64* {
	%append_output( SWIG_From_long_SS_long( *$1 ) );
}

/*GoUint64* as function return typemap*/
%typemap(argout) GoUint64* {
	%append_output( SWIG_From_unsigned_SS_long_SS_long( *$1 ) );
}

/*GoUint32* as function return typemap*/
%typemap(argout) GoUint32* {
	%append_output( SWIG_From_unsigned_SS_int( *$1 ) );
}

/*GoUint8* as function return typemap*/
%typemap(argout) GoUint8* {
	%append_output( SWIG_From_unsigned_SS_char( *$1 ) );
}

%apply int* OUTPUT {GoInt*}
%apply int* OUTPUT {GoUint*}
%apply int* OUTPUT {GoUint8*}
%apply int* OUTPUT {GoInt8*}
%apply int* OUTPUT {GoUint16*}
%apply int* OUTPUT {GoInt16*}
%apply int* OUTPUT {GoInt32*}
%apply long* OUTPUT {GoInt64*}
%apply long* OUTPUT {GoUint64*}
