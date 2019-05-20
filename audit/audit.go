package audit

/*
*  FUNCTION:   CheckErr
*  PARAM:      err=  Error object, module= Name source
*  RETURN:     boolean value.
*  DESCRIPTION:Method that evaluates the error object
 */

func CheckErr(err error, module string) bool {
	var r bool
	if err != nil {
		s := err.Error()
		println("Error: " + s)
		r = false
	} else {
		r = true
	}
	return r
}
