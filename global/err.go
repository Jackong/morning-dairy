/**
 * User: jackong
 * Date: 11/12/13
 * Time: 11:14 AM
 */
package global

type AccessError struct {
	Code int
	Msg string
}

func (this AccessError) Error() string{
    return this.Msg
}

const (
	CODE_OK = 0
)
