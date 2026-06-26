package types

type Authorize struct {
	Key    string
	Header string
	Cookie string
}

/*	for unique auth types (eg: key in body)
	we can add special field and fill that
	when we need use this type of authorization */
