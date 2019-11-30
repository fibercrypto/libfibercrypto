/**
*
* typemaps for Handles
*
**/

/* Handle reference typemap. */
%typemap(in, numinputs=0) Handle* (Handle temp) {
	$1 = &temp;
}

/* Handle out typemap. */
%typemap(argout) Handle* {
	%append_output( SWIG_From_long_SS_long(*$1) );
}

/* Handle not as pointer is input. */
%typemap(in) Handle {
	SWIG_AsVal_long_SS_long($input, (long*)&$1);
}


%apply Handle { Strings__Handle
	}

%apply Handle* { Strings__Handle*
	}